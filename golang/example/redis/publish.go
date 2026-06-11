package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func publish(rdb *redis.Client, channel string) {
	// publish event
	err := rdb.Publish(ctx, channel, "bid:100").Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("event published to channel:", channel)
}
