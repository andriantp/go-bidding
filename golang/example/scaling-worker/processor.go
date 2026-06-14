package main

import "fmt"

func ProcessBidEvent(room string, message []byte) {
	bid, err := ParseBid(message)
	if err != nil {
		fmt.Println("invalid payload")
		return
	}

	// atomic validation
	err = ValidateBidAtomicWithRetry(room, bid)
	if err != nil {
		fmt.Println(err)
		return
	}

	// publish realtime
	Publish(room, message)

	fmt.Println("processed bid: from room", room, "by user", bid.User, "with price", bid.Price)
}
