package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findBestProfitableOperation(pricesStr string) int {
	fmt.Println("Prices:", pricesStr)
	var profit int = 0
	prices := strings.Split(pricesStr, ",")

	profits := []int{}
	for idx, price := range prices {

		for idxTrade, priceToTrade := range prices {
			p, err := strconv.Atoi(price)
			if err != nil {
				panic(err)
			}
			pt, err := strconv.Atoi(priceToTrade)
			if err != nil {
				panic(err)
			}

			// não chegar preços passados
			if idxTrade >= idx {
				profits = append(profits, pt-p)
			}
		}
	}

	for _, p := range profits {
		if p > profit {
			profit = p
		}
	}

	return profit
}
