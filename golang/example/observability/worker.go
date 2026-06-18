package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
)

func StartWorker(room string, consumer string) {
	stream := GetRoomStream(room)
	group := GetRoomConsumerGroup(room)
	log.Printf("[worker=%s] [room=%s] worker started", consumer, room)

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
			log.Printf("[worker=%s] [room=%s] worker read error err=%v", consumer, room, err)
			continue
		}

		for _, s := range streams {
			for _, msg := range s.Messages {
				payloadVal, ok := msg.Values["payload"]
				if !ok {
					log.Printf("[worker=%s] [room=%s] skip invalid payload message_id=%s", group, room, msg.ID)
					continue
				}

				payload := fmt.Sprintf("%v", payloadVal)
				log.Printf("[trace=%s] [worker=%s] [room=%s] processing started", msg.ID, group, room)
				// log.Printf("[worker=%s] [room=%s] processing message_id=%s", group, room, msg.ID)
				if err = ProcessBidEvent(room, msg.ID, []byte(payload)); err != nil {
					log.Printf("[worker=%s] [room=%s] process error message_id=%s err=%v", group, room, msg.ID, err)
					continue
				}

				//metric
				atomic.AddUint64(&processedEvents, 1)

				if err = AckEvent(room, msg.ID); err != nil {
					log.Printf("[worker=%s] [room=%s] ack error message_id=%s err=%v", group, room, msg.ID, err)
					continue
				}

				// log.Printf("[worker=%s] [room=%s] event acknowledged message_id=%s", group, room, msg.ID)
				log.Printf("[trace=%s] [worker=%s] [room=%s] acknowledged", msg.ID, group, room)
			}
		}
	}
}
