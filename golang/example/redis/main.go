package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Butuh argumen")
		return
	}

	fmt.Println("Argumen:", os.Args[1])

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	})

	channel := "auction_room"
	switch os.Args[1] {
	case "connect":
		connect(rdb)
	case "publish":
		publish(rdb, channel)
	case "subscribe":
		subscribe(rdb, channel)
	default:
		fmt.Println("Argumen tidak dikenali")
	}

}
