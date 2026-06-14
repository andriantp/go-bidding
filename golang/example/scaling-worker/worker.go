package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func StartWorker(room string) {
	stream := GetRoomStream(room)
	lastID := "0"
	fmt.Println("worker started for room:", room)
	for {
		streams, err := rdb.XRead(
			ctx,
			&redis.XReadArgs{
				Streams: []string{
					stream,
					lastID,
				},
				Block: 0,
				Count: 1,
			},
		).Result()

		if err != nil {
			fmt.Println("worker read error:", err)
			continue
		}

		for _, s := range streams {
			for _, msg := range s.Messages {
				payloadVal, ok := msg.Values["payload"]
				if !ok {
					fmt.Println("skip old event format without payload message id:", msg.ID)
					lastID = msg.ID
					continue
				}

				payload := fmt.Sprintf("%v",payloadVal)
				ProcessBidEvent(room, []byte(payload))
				lastID = msg.ID
			}
		}
	}
}
