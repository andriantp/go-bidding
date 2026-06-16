package main

import "fmt"

func ProcessBidEvent(room string, message []byte) error {
	bid, err := ParseBid(message)
	if err != nil {
		return fmt.Errorf("ParseBid:%w", err)
	}

	// atomic validation
	if err = ValidateBidAtomicWithRetry(room, bid); err != nil {
		return fmt.Errorf("ValidateBidAtomicWithRetry:%w", err)
	}

	// publish realtime
	Publish(room, message)

	fmt.Println("processed bid: from room", room, "by user", bid.User, "with price", bid.Price)
	return nil
}
