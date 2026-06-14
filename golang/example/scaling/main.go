package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . <port> <room>")
		return
	}

	port := os.Args[1]
	room := os.Args[2]

	fmt.Println("starting server on :",port)
	fmt.Println("room :",room)

	// recovery state
	if err := RecoverStateFromStream(room); err != nil {
		panic(err)
	}

	// realtime subscriber
	go StartRedisSubscriber(room)

	// websocket route
	http.HandleFunc("/ws", WsHandler)

	// start server
	if err := http.ListenAndServe(":"+port,nil,); err != nil {
		panic(err)
	}
}
