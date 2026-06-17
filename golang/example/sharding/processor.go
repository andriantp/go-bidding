package main

import "fmt"

func ProcessBidEvent(room string, messageID string, message []byte) error {
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

	// ======================== Snapshot ========================
	snapshot := Snapshot{
		HighestBid:    bid.Price,
		HighestBidder: bid.User,
		LastEventID:   messageID,
	}

	if err = SaveSnapshot(room, snapshot); err != nil {
		fmt.Println("snapshot save error:", err)
	}

	fmt.Println("processed bid: from room", room, "by user", bid.User, "with price", bid.Price)
	return nil
}
