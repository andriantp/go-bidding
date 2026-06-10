package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all (dev only)
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}

	fmt.Println("client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		fmt.Println("message:", string(msg))

		// echo balik ke client
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("server jalan di :8080")
	http.ListenAndServe(":8080", nil)
}
