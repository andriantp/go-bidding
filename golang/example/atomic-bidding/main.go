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

	go StartRedisSubscriber()

	http.HandleFunc("/ws", WsHandler)
	fmt.Println("server jalan di :", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
