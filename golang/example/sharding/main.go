package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("usage: go run . <port> <room> <group> <current-shard> <total-shard>")
		return
	}

	port := os.Args[1]
	room := os.Args[2]
	group := os.Args[3]

	currentShard, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatalf("currentShard:%v", err)
	}

	totalShard, err := strconv.Atoi(os.Args[5])
	if err != nil {
		log.Fatalf("totalShard:%v", err)
	}

	fmt.Println("server on    :", port)
	fmt.Println("room         :", room)
	fmt.Println("group        :", group)
	fmt.Println("currentShard :", currentShard)
	fmt.Println("totalShard   :", totalShard)

	if !IsOwner(room, currentShard, totalShard) {
		fmt.Println("skip room, not owner:", room)
		return
	}

	// recovery from snapshot
	if err := RecoverFromSnapshot(room); err != nil {
		log.Fatalf("RecoverFromSnapshot:%v", err)
	}

	if err := CreateConsumerGroup(room); err != nil {
		log.Fatalf("CreateConsumerGroup:%v", err)
	}

	// realtime subscriber
	go StartRedisSubscriber(room)

	// start worker
	go StartWorker(room, group)

	// websocket route
	http.HandleFunc("/ws", WsHandler)

	// start server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("ListenAndServe:%v", err)
	}
}
