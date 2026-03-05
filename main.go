package main

import (
	"auction/internal/auction"
	"auction/internal/metrics"
	"auction/internal/resource"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	timer := metrics.StartTimer()

	cpuLimiter := resource.NewCPULimiter(8)

	fmt.Println("Starting Auction Simulation")

	auction.StartAuctions(cpuLimiter)

	total := timer.Stop()

	fmt.Println("All auctions finished")
	fmt.Println("Total execution time:", total)
}
