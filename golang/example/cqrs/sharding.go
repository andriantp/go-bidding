package main

import (
	"fmt"
	"hash/fnv"
)

func GetShard(room string, totalShard int) int {
	h := fnv.New32a()
	h.Write([]byte(room))
	return int(h.Sum32()) % totalShard
}

func IsOwner(room string, currentShard int, totalShard int) bool {
	shard := GetShard(room, totalShard)

	fmt.Printf(
		"room [%s] shard:[%d] current:[%d]\n",
		room,
		shard,
		currentShard,
	)

	return shard == currentShard
}

func GetBackupShard(primaryShard int, totalShard int) int {
	return (primaryShard + 1) % totalShard
}

func IsBackupOwner(room string,currentShard int,totalShard int) bool {
	primary := GetShard(room,totalShard)
	backup := GetBackupShard(primary,totalShard)
	return backup == currentShard
}