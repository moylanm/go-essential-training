package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text)
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	for word, count := range wordMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}
