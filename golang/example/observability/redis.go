package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(
	&redis.Options{
		Addr: "localhost:6380",
	},
)

// ======================== Model ========================

type Bid struct {
	User  string `json:"user"`
	Price int    `json:"price"`
}

// ======================== Publish ========================

func Publish(room string, message []byte) {
	channel := GetRoomChannel(room)
	err := rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println("publish error:", err)
	}
}

// ======================== Subscribe ========================

func StartRedisSubscriber(room string) {
	channel := GetRoomChannel(room)
	sub := rdb.Subscribe(ctx, channel)
	ch := sub.Channel()
	fmt.Println("subscribed to:", channel)

	for msg := range ch {
		Broadcast(room, []byte(msg.Payload))
	}
}

// ======================== Parse ========================

func ParseBid(message []byte) (Bid, error) {
	var bid Bid
	err := json.Unmarshal(message, &bid)
	return bid, err
}

// ======================== Save Event ========================
func SaveBidEvent(room string, bid Bid) error {
	stream := GetRoomStream(room)
	_, err := rdb.XAdd(
		ctx,
		&redis.XAddArgs{
			Stream: stream,
			Values: map[string]interface{}{
				"user":  bid.User,
				"price": bid.Price,
			},
		},
	).Result()

	return err
}

// ======================== Atomic Validation ========================

func ValidateBidAtomicWithRetry(room string, bid Bid) error {
	maxRetry := 3
	highestBidKey := GetHighestBidKey(room)
	highestBidderKey := GetHighestBidderKey(room)

	for i := 0; i < maxRetry; i++ {
		err := rdb.Watch(ctx, func(tx *redis.Tx) error {
			val, err := tx.Get(ctx, highestBidKey).Result()
			if err != nil && err != redis.Nil {
				return err
			}

			currentHighest := 0
			if val != "" {
				currentHighest, _ = strconv.Atoi(val)
			}

			// validation
			if bid.Price <= currentHighest {
				return fmt.Errorf(
					"bid rejected: current highest is %d",
					currentHighest,
				)
			}

			// atomic transaction
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, highestBidKey, bid.Price, 0)
				pipe.Set(ctx, highestBidderKey, bid.User, 0)
				return nil
			},
			)
			return err
		},
			highestBidKey,
		)

		// success
		if err == nil {
			return nil
		}

		// retry on conflict
		if err == redis.TxFailedErr {
			fmt.Println("retry transaction...")
			continue
		}

		return err
	}

	return fmt.Errorf("failed after max retry")
}

// ======================== Recovery ========================

func RecoverStateFromStream(room string) error {
	stream := GetRoomStream(room)
	streams, err := rdb.XRange(ctx, stream, "-", "+").Result()
	if err != nil {
		return err
	}

	highestBid := 0
	highestBidder := ""

	for _, msg := range streams {
		priceStr := fmt.Sprintf("%v", msg.Values["price"])
		price, _ := strconv.Atoi(priceStr)
		user := fmt.Sprintf("%v", msg.Values["user"])

		fmt.Println("replay:", user, price)
		if price > highestBid {
			highestBid = price
			highestBidder = user
		}
	}

	fmt.Println("recovered highest bid:", highestBid)
	fmt.Println("recovered highest bidder:", highestBidder)
	return nil
}

// ======================== Helper ========================
func GetRoomChannel(room string) string      { return "auction_room:" + room }
func GetRoomStream(room string) string       { return "stream:" + room }
func GetHighestBidKey(room string) string    { return "highest_bid:" + room }
func GetHighestBidderKey(room string) string { return "highest_bidder:" + room }
func PushBidToStream(room string, message []byte) error {
	stream := GetRoomStream(room)
	_, err := rdb.XAdd(
		ctx,
		&redis.XAddArgs{
			Stream: stream,
			Values: map[string]interface{}{
				"payload": string(message),
			},
		},
	).Result()

	return err
}

// ======================== Consumer Group ========================
func GetRoomConsumerGroup(room string) string {
	return "group:" + room
}

func CreateConsumerGroup(room string) error {
	stream := GetRoomStream(room)
	group := GetRoomConsumerGroup(room)
	err := rdb.XGroupCreateMkStream(
		ctx,
		stream,
		group,
		"$",
	).Err()

	if err != nil {
		// group already exists
		if err.Error() == "BUSYGROUP Consumer Group name already exists" {
			return nil
		}
		return err
	}
	fmt.Println("consumer group created:", group)
	return nil
}

// ======================== Acknowledgement ========================
func AckEvent(room string, messageID string) error {
	stream := GetRoomStream(room)
	group := GetRoomConsumerGroup(room)
	return rdb.XAck(
		ctx,
		stream,
		group,
		messageID,
	).Err()
}

// ======================== Snapshot ========================
type Snapshot struct {
	HighestBid    int
	HighestBidder string
	LastEventID   string
}

func GetRoomSnapshotKey(room string) string {
	return "snapshot:" + room
}

func SaveSnapshot(room string, s Snapshot) error {
	key := GetRoomSnapshotKey(room)
	return rdb.HSet(
		ctx,
		key,
		map[string]interface{}{
			"highest_bid":    s.HighestBid,
			"highest_bidder": s.HighestBidder,
			"last_event_id":  s.LastEventID,
		},
	).Err()
}

func LoadSnapshot(room string) (Snapshot, error) {
	key := GetRoomSnapshotKey(room)
	result, err := rdb.HGetAll(
		ctx,
		key,
	).Result()

	if err != nil {
		return Snapshot{}, err
	}

	if len(result) == 0 {
		return Snapshot{}, nil
	}

	highestBid, _ := strconv.Atoi(
		result["highest_bid"],
	)

	return Snapshot{
		HighestBid:    highestBid,
		HighestBidder: result["highest_bidder"],
		LastEventID:   result["last_event_id"],
	}, nil
}

func RecoverFromSnapshot(room string) error {
	snapshot, err := LoadSnapshot(
		room,
	)
	if err != nil {
		return err
	}

	fmt.Println("snapshot highest bid:", snapshot.HighestBid)
	fmt.Println("snapshot highest bidder:", snapshot.HighestBidder)
	startID := GetNextStreamID(snapshot.LastEventID)
	stream := GetRoomStream(room)
	streams, err := rdb.XRange(
		ctx,
		stream,
		startID,
		"+",
	).Result()

	if err != nil {
		return err
	}

	for _, msg := range streams {
		fmt.Println("replay recent event:", msg.ID)
	}

	return nil
}

func GetNextStreamID(id string) string {
	if id == "" {
		return "-"
	}
	return "(" + id
}

func AppendBidEvent(room string, message []byte) error {
	stream := GetRoomStream(room)
	_, err := rdb.XAdd(
		ctx,
		&redis.XAddArgs{
			Stream: stream,
			Values: map[string]interface{}{
				"payload": string(message),
			},
		},
	).Result()

	return err
}

func ReplayBidHistory(room string) error {
	stream := GetRoomStream(room)
	streams, err := rdb.XRange(
		ctx,
		stream,
		"-",
		"+",
	).Result()

	if err != nil {
		return err
	}

	highestBid := 0
	highestBidder := ""
	for _, msg := range streams {
		payload := fmt.Sprintf("%v", msg.Values["payload"])
		var bid Bid
		if err := json.Unmarshal([]byte(payload), &bid); err != nil {
			log.Printf("[replay] [room=%s] invalid replay payload message_id=%s", room, msg.ID)
			continue
		}

		//metric
		atomic.AddUint64(&replayEvents, 1)

		// log.Printf("[trace=%s] [replay] [room=%s] replay user=%s price=%d", msg.ID, room, bid.User, bid.Price)
		log.Printf("[replay] [room=%s] replay user=%s price=%d message_id=%s", room, bid.User, bid.Price, msg.ID)
		if bid.Price > highestBid {
			highestBid = bid.Price
			highestBidder = bid.User
		}
	}

	log.Printf("[replay] [room=%s] recovered highest bid=%d bidder=%s", room, highestBid, highestBidder)

	return nil
}
