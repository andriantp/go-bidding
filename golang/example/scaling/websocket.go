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

var rooms = make(
	map[string]map[*websocket.Conn]bool,
)

var mu sync.Mutex

// ======================== WebSocket ========================

func WsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// get room
	room := r.URL.Query().Get("room")

	if room == "" {
		http.Error(
			w,
			"room required",
			http.StatusBadRequest,
		)
		return
	}

	// upgrade connection
	conn, err := upgrader.Upgrade(
		w,
		r,
		nil,
	)

	if err != nil {
		fmt.Println("upgrade error:",err)
		return
	}

	// register room
	mu.Lock()

	if rooms[room] == nil {

		rooms[room] = make(
			map[*websocket.Conn]bool,
		)
	}

	rooms[room][conn] = true

	mu.Unlock()

	fmt.Println("client connected to room:",room)

	// read loop
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)

			mu.Lock()
			delete(rooms[room],conn)
			mu.Unlock()

			conn.Close()
			return
		}

		// parse payload
		bid, err := ParseBid(msg)
		if err != nil {
			fmt.Println("invalid payload:", err)

			continue
		}

		// atomic validation
		err = ValidateBidAtomicWithRetry(
			room,
			bid,
		)
		if err != nil {
			fmt.Println("validation error:", err)
			continue
		}

		// save event
		err = SaveBidEvent(
			room,
			bid,
		)

		if err != nil {

			fmt.Println(
				"save stream error:",
				err,
			)

			continue
		}

		// publish realtime
		Publish(
			room,
			msg,
		)
	}
}

// ======================== Broadcast ========================

func Broadcast(
	room string,
	message []byte,
) {

	mu.Lock()
	defer mu.Unlock()

	for c := range rooms[room] {

		err := c.WriteMessage(
			websocket.TextMessage,
			message,
		)

		if err != nil {

			c.Close()

			delete(
				rooms[room],
				c,
			)
		}
	}
}