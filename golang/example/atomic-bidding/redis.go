package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6380",
})

var channel = "auction_room"

type Bid struct {
	User  string `json:"user"`
	Price int    `json:"price"`
}

// ======================== Publish ========================

func Publish(message []byte) {
	err := rdb.Publish(ctx, channel, message).Err()

	if err != nil {
		fmt.Println("publish error:", err)
	}
}

// ======================== Subscribe ========================

func StartRedisSubscriber() {
	sub := rdb.Subscribe(ctx, channel)
	ch := sub.Channel()

	fmt.Println("subscribed to:", channel)

	for msg := range ch {
		Broadcast([]byte(msg.Payload))
	}
}

// ======================== Parse ========================

func ParseBid(message []byte) (Bid, error) {
	var bid Bid

	err := json.Unmarshal(message, &bid)

	return bid, err
}

// ======================== Atomic Validation ========================

func ValidateBidAtomic(bid Bid) error {
	return rdb.Watch(ctx, func(tx *redis.Tx) error {

		// ambil highest bid saat ini
		val, err := tx.Get(ctx, "highest_bid").Result()
		if err != nil && err != redis.Nil {
			return err
		}

		currentHighest := 0

		if val != "" {
			currentHighest, _ = strconv.Atoi(val)
		}

		// validasi
		if bid.Price <= currentHighest {
			return fmt.Errorf(
				"bid rejected: current highest is %d",
				currentHighest,
			)
		}

		// transaction update
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {

			pipe.Set(ctx, "highest_bid", bid.Price, 0)
			pipe.Set(ctx, "highest_bidder", bid.User, 0)

			return nil
		})

		return err

	}, "highest_bid")
}

func ValidateBidAtomicWithRetry(
	bid Bid,
) error {

	maxRetry := 3

	for i := 0; i < maxRetry; i++ {

		err := rdb.Watch(ctx, func(tx *redis.Tx) error {

			// ambil highest bid saat ini
			val, err := tx.Get(
				ctx,
				"highest_bid",
			).Result()

			if err != nil && err != redis.Nil {
				return err
			}

			currentHighest := 0

			if val != "" {
				currentHighest, _ = strconv.Atoi(val)
			}

			// validasi
			if bid.Price <= currentHighest {
				return fmt.Errorf(
					"bid rejected: current highest is %d",
					currentHighest,
				)
			}

			// transaction update
			_, err = tx.TxPipelined(
				ctx,
				func(pipe redis.Pipeliner) error {

					pipe.Set(
						ctx,
						"highest_bid",
						bid.Price,
						0,
					)

					pipe.Set(
						ctx,
						"highest_bidder",
						bid.User,
						0,
					)

					return nil
				},
			)

			return err

		}, "highest_bid")

		// success
		if err == nil {
			return nil
		}

		// retry jika conflict
		if err == redis.TxFailedErr {
			fmt.Println("retry transaction...")
			continue
		}

		return err
	}

	return fmt.Errorf(
		"failed after max retry",
	)
}
