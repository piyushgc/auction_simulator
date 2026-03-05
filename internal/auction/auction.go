package auction

import (
	"auction/internal/bidder"
	"auction/internal/model"
	"context"
	"fmt"
	"sync"
	"time"
)

const TotalBidders = 100

func RunAuction(a model.Auction, timeout time.Duration) {
	fmt.Println("Auction started:", a.ID)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	bidChan := make(chan model.Bid, 100)

	var wg sync.WaitGroup
	wg.Add(TotalBidders)

	for i := 1; i <= TotalBidders; i++ {
		go func(id int) {
			defer wg.Done()
			bidder.SimulateBidder(ctx, id, bidChan)
		}(i)
	}

	go func() {
		wg.Wait()
		close(bidChan)
	}()

	var winner model.Bid

	for {
		select {

		case bid := <-bidChan:

			if bid.Amount > winner.Amount {
				winner = bid
			}

		case <-ctx.Done():

			fmt.Printf(
				"Auction %d Winner -> Bidder %d Amount %d\n",
				a.ID,
				winner.BidderID,
				winner.Amount,
			)

			return
		}
	}
}
