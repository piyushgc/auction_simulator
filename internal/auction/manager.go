package auction

import (
	"auction/internal/model"
	"auction/internal/resource"
	"math/rand"
	"sync"
	"time"
)

const TotalAuctions = 40

func StartAuctions(cpuLimiter *resource.CPULimiter) {

	var wg sync.WaitGroup
	wg.Add(TotalAuctions)

	for i := 1; i <= TotalAuctions; i++ {

		auction := model.Auction{
			ID: i,
		}

		for j := 0; j < 20; j++ {
			auction.Attributes[j] = rand.Intn(100)
		}

		go func(a model.Auction) {

			defer wg.Done()

			cpuLimiter.Acquire()
			defer cpuLimiter.Release()

			RunAuction(a, 3*time.Second)

		}(auction)
	}

	wg.Wait()
}
