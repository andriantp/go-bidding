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

// ======================== Validation ========================

func ValidateBid(bid Bid) error {
	val, err := rdb.Get(ctx, "highest_bid").Result()

	if err != nil && err != redis.Nil {
		return err
	}

	currentHighest := 0

	if val != "" {
		currentHighest, _ = strconv.Atoi(val)
	}

	// reject jika <= highest bid
	if bid.Price <= currentHighest {
		return fmt.Errorf(
			"bid rejected: current highest is %d",
			currentHighest,
		)
	}

	// update state
	rdb.Set(ctx, "highest_bid", bid.Price, 0)
	rdb.Set(ctx, "highest_bidder", bid.User, 0)
	fmt.Printf("bid accepted: %s bids %d\n", bid.User, bid.Price)

	return nil
}

// ======================== Helpers ========================

func ParseBid(message []byte) (Bid, error) {
	var bid Bid

	err := json.Unmarshal(message, &bid)

	return bid, err
}
