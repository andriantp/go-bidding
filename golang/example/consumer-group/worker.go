package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func StartWorker(room string, consumer string) {
	stream := GetRoomStream(room)
	group := GetRoomConsumerGroup(room)
	fmt.Println("worker started:", consumer)
	for {
		streams, err := rdb.XReadGroup(
			ctx,
			&redis.XReadGroupArgs{
				Group:    group,
				Consumer: consumer,
				Streams: []string{
					stream,
					">",
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
					continue
				}
				payload := fmt.Sprintf("%v", payloadVal)

				if err = ProcessBidEvent(room, []byte(payload));err != nil {
					fmt.Println("process error:", err)
					continue
				}

				if err = AckEvent(room,msg.ID); err != nil {
					fmt.Println("ack error:",err)
					continue
				}

				fmt.Println("event acknowledged:",msg.ID)
			}
		}
	}
}
