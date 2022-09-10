package main

import "testing"

func TestProfitable(t *testing.T) {
	t.Run("should find a possibly to buy and sell", func(t *testing.T) {
		prices := "7,1,5,3,6,4"
		result := findBestProfitableOperation(prices)
		expected := 5
		if result != expected {
			t.Errorf("Result: +%v, Expected: %+v", result, expected)
		}
	})

	t.Run("should not find a possibly to buy and sell", func(t *testing.T) {
		prices := "7,6,4,3,1"
		result := findBestProfitableOperation(prices)
		expected := 0
		if result != expected {
			t.Errorf("Result: +%v, Expected: %+v", result, expected)
		}
	})
}
