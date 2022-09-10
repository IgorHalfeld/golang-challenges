package main

import "testing"

func TestFrequency(t *testing.T) {
	t.Run("should return most frequency word with an array of forbidden words", func(t *testing.T) {
		paragraph := "Bob hit a ball, the hit BALL flew long after it was hit."
		forbiddenWords := []string{"hit"}

		result := getMoreFrequentWordFilteringByForbiddenWords(paragraph, forbiddenWords)
		expected := "ball"
		if result != expected {
			t.Errorf("Result: +%v, Expected: %+v", result, expected)
		}
	})

	t.Run("should return most frequency word with empty array of forbidden words", func(t *testing.T) {
		paragraph := "a."
		forbiddenWords := []string{""}

		result := getMoreFrequentWordFilteringByForbiddenWords(paragraph, forbiddenWords)
		expected := "a"
		if result != expected {
			t.Errorf("Result: +%v, Expected: %+v", result, expected)
		}
	})
}
