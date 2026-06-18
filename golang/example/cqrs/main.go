package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 6 {
		fmt.Println("usage: go run . <port> <room> <group> <current-shard> <total-shard>")
		return
	}

	port := os.Args[1]
	room := os.Args[2]
	group := os.Args[3]

	currentShard, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatalf("currentShard:%v", err)
	}

	totalShard, err := strconv.Atoi(os.Args[5])
	if err != nil {
		log.Fatalf("totalShard:%v", err)
	}

	fmt.Println("server on    :", port)
	fmt.Println("room         :", room)
	fmt.Println("group        :", group)
	fmt.Println("currentShard :", currentShard)
	fmt.Println("totalShard   :", totalShard)

	primary := GetShard(room, totalShard)
	backup := GetBackupShard(primary, totalShard)

	fmt.Println("primary shard:", primary)
	fmt.Println("backup shard :", backup)

	isPrimary := IsOwner(room, currentShard, totalShard)
	isBackup := IsBackupOwner(room, currentShard, totalShard)

	if !isPrimary && !isBackup {
		fmt.Println("skip room:", room)
		return
	}

	/*if isPrimary {
		fmt.Println("mode : PRIMARY")
	}
	if isBackup {
		fmt.Println("mode : BACKUP")
	}*/

	mode := "PRIMARY"
	if isBackup {
		mode = "BACKUP"
	}

	fmt.Println("mode :", mode)

	// if err := RecoverFromSnapshot(room); err != nil {
	// 	log.Fatalf("RecoverFromSnapshot:%v", err)
	// }
	
	if err := ReplayBidHistory(room); err != nil {
		log.Fatalf("ReplayBidHistory:%v", err)
	}

	if err := CreateConsumerGroup(room); err != nil {
		log.Fatalf("CreateConsumerGroup:%v", err)
	}

	go StartRedisSubscriber(room)
	go StartWorker(room, group)

	http.HandleFunc("/ws", WsHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("ListenAndServe:%v", err)
	}
}
