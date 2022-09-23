package main

import (
	"math"
	"strconv"
	"strings"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findBestProfitableOperation(pricesStr string) int {
	var profit int = 0
	prices := strings.Split(pricesStr, ",")

	min := math.MaxInt

	for _, price := range prices {
		p, err := strconv.Atoi(price)
		if err != nil {
			panic(err)
		}

		if p > min {
			profit = max(profit, p-min)
		} else {
			min = p
		}
	}

	return profit
}
