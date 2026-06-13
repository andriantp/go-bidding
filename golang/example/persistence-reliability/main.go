package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <port>")
		return
	}

	port := os.Args[1]

	fmt.Println("starting server on :", port)


	// recovery state
	if err := RecoverStateFromStream(); err != nil {
		panic(err)
	}

	// realtime subscriber
	go StartRedisSubscriber()

	// websocket route
	http.HandleFunc("/ws", WsHandler)

	// start server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}