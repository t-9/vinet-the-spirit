package main

import (
	"fmt"
	"log"

	"vinet/market"
)

func main() {
	markets, err := market.GetMarkets()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ProductCode, MarketType, Alias")
	for _, m := range markets {
		fmt.Println(m)
	}
}
