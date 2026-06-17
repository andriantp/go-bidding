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
func WsHandler(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	if room == "" {
		http.Error(w, "room required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}

	RegisterClient(room, conn)

	fmt.Println("client connected to room:", room)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			RemoveClient(room, conn)
			conn.Close()
			return
		}

		// push event to stream
		err = PushBidToStream(room, msg)
		if err != nil {
			fmt.Println("push stream error:", err)
		}
	}
}

// ======================== Helper Client Management ========================
func RegisterClient(room string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	if rooms[room] == nil {
		rooms[room] = make(
			map[*websocket.Conn]bool,
		)
	}
	rooms[room][conn] = true
}

func RemoveClient(room string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(rooms[room], conn)
}

// ======================== Broadcast ========================
func Broadcast(room string, message []byte) {
	mu.Lock()
	defer mu.Unlock()
	for conn := range rooms[room] {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			conn.Close()
			delete(rooms[room], conn)
		}
	}
}
