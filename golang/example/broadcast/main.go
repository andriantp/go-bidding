package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// simpan semua client
var clients = make(map[*websocket.Conn]bool)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}

	clients[conn] = true
	fmt.Println("client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			delete(clients, conn)
			conn.Close()
			return
		}

		fmt.Println("broadcast:", string(msg))

		// broadcast ke semua client
		for c := range clients {
			err := c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				c.Close()
				delete(clients, c)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("server jalan di :8080")
	http.ListenAndServe(":8080", nil)
}