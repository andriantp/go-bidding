package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . <port> <room> <group>")
		return
	}

	port := os.Args[1]
	room := os.Args[2]
	group := os.Args[3]

	fmt.Println("starting server on :", port)
	fmt.Println("room  :", room)
	fmt.Println("group :", group)

	// recovery state
	if err := RecoverStateFromStream(room); err != nil {
		panic(err)
	}

	if err := CreateConsumerGroup(room); err != nil {
		panic(err)
	}

	// realtime subscriber
	go StartRedisSubscriber(room)

	// start worker
	go StartWorker(room, group)

	// websocket route
	http.HandleFunc("/ws", WsHandler)

	// start server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
