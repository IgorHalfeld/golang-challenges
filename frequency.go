package main

import (
	"strings"

	"github.com/gosimple/slug"
)

// Não usei generics pra manter a simplicidade
func contains(str []string, match string) bool {
	for _, word := range str {
		normalizedWord := slug.Make(word)
		if normalizedWord == match {
			return true
		}
	}

	return false
}

func getMoreFrequentWordFilteringByForbiddenWords(paragraph string, forbiddenWords []string) string {
	var mostFrequentAllowed string
	hashmapCount := map[string]int{}
	words := strings.Split(paragraph, " ")

	for _, word := range words {
		// a estratégia de usar slugs pra evitar maisculuas e acentuação
		normalizedWord := slug.Make(word)

		if contains(forbiddenWords, normalizedWord) {
			continue
		}

		_, exists := hashmapCount[normalizedWord]
		if exists {
			hashmapCount[normalizedWord] += 1
		} else {
			hashmapCount[normalizedWord] = 1
		}
	}

	lastHighFrequency := 0
	for word := range hashmapCount {
		frequency := hashmapCount[word]

		if frequency > lastHighFrequency {
			lastHighFrequency = frequency
			mostFrequentAllowed = word
		}

	}

	return mostFrequentAllowed
}
