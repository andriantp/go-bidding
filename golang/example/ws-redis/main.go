package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Butuh argumen port")
		return
	}

	port := os.Args[1]
	fmt.Println("Port:", port)

	// start redis subscriber (background)
	go StartRedisSubscriber()

	http.HandleFunc("/ws", WsHandler)

	fmt.Println("server jalan di :", port)
	http.ListenAndServe(":"+port, nil)
}
