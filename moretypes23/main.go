// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful.

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordsMap := map[string]int{}
	fields := strings.Fields(s)
	for _, f := range fields {
		wordsMap[f] += 1
	}
	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
