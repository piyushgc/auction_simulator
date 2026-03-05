package bidder

import (
	"auction/internal/model"
	"context"
	"math/rand"
)

func SimulateBidder(ctx context.Context, bidderID int, bidChan chan<- model.Bid) {

	// not every bidder responds
	if rand.Intn(100) < 40 {
		return
	}

	bid := model.Bid{
		BidderID: bidderID,
		Amount:   rand.Intn(1000) + 1,
	}

	select {
	case <-ctx.Done():
		return
	case bidChan <- bid:
	}
}
