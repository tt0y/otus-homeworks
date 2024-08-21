package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func countWords(s string) map[string]int {
	// Слайс слов из строки
	words := strings.Fields(s)
	// Мапа слово - частотность
	wordCounted := make(map[string]int)

	for _, word := range words {
		wordCounted[word]++
	}

	return wordCounted
}

func sortWordCounted(wordCounted map[string]int) []string {
	var wordsSorted []string

	for word := range wordCounted {
		wordsSorted = append(wordsSorted, word)
	}

	sort.Slice(wordsSorted, func(i, j int) bool {
		if wordCounted[wordsSorted[i]] == wordCounted[wordsSorted[j]] {
			return wordsSorted[i] < wordsSorted[j]
		}
		return wordCounted[wordsSorted[i]] > wordCounted[wordsSorted[j]]
	})

	if len(wordsSorted) > 10 {
		return wordsSorted[:10]
	}

	return wordsSorted
}

func Top10(s string) []string {
	return sortWordCounted(countWords(s))
}
