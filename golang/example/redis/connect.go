package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func connect(rdb *redis.Client) {
	// test koneksi
	err := rdb.Set(ctx, "test", "hello redis", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("value:", val)
}
