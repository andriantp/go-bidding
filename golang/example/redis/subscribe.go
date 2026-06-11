package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func subscribe(rdb *redis.Client, channel string) {
	// subscribe ke channel
	sub := rdb.Subscribe(ctx, channel)

	// ambil channel Go
	ch := sub.Channel()

	fmt.Println("listening to channel:", channel)

	for msg := range ch {
		fmt.Println("received:", msg.Payload)
	}
}
