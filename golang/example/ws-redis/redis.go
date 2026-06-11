package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6380",
})

var channel = "auction_room"

// publish ke redis
func Publish(message []byte) {
	err := rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println("publish error:", err)
	}
}

// subscriber
func StartRedisSubscriber() {
	sub := rdb.Subscribe(ctx, channel)
	ch := sub.Channel()

	fmt.Println("subscribed to:", channel)

	for msg := range ch {
		Broadcast([]byte(msg.Payload))
	}
}
