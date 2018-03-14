package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	pksTileVersion := os.Getenv("PKS_TILE_VERSION")
	aggregatorEndpoint := os.Getenv("AGGREGATOR_ENDPOINT")

	fmt.Println("PKS Tile Version: ", pksTileVersion)
	fmt.Println("Aggregator Endpoint: ", aggregatorEndpoint)

	for {
		time.Sleep(10 * time.Second)
	}

}
