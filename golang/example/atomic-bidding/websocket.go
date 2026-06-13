package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

var mu sync.Mutex

// ======================== WebSocket ========================

func WsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	conn, err := upgrader.Upgrade(
		w,
		r,
		nil,
	)

	if err != nil {
		fmt.Println(
			"upgrade error:",
			err,
		)
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	fmt.Println("client connected")

	for {

		_, msg, err := conn.ReadMessage()

		if err != nil {

			fmt.Println(
				"read error:",
				err,
			)

			mu.Lock()
			delete(clients, conn)
			mu.Unlock()

			conn.Close()

			return
		}

		// parse payload
		bid, err := ParseBid(msg)

		if err != nil {
			fmt.Println("invalid payload")
			continue
		}

		// atomic validation + retry
		err = ValidateBidAtomicWithRetry(
			bid,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		// publish hanya jika sukses
		Publish(msg)
	}
}

// ======================== Broadcast ========================

func Broadcast(message []byte) {

	mu.Lock()
	defer mu.Unlock()

	for c := range clients {

		err := c.WriteMessage(
			websocket.TextMessage,
			message,
		)

		if err != nil {

			c.Close()

			delete(clients, c)
		}
	}
}
