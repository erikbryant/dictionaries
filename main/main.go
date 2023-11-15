package main

import (
	"fmt"
	"github.com/erikbryant/dictionaries"
)

func stats(totalWords int, words []string, a map[string]int) {
	dupes := 0
	for _, val := range a {
		if val > 1 {
			dupes += val
		}
	}
	portion := (float64(len(words)) / float64(totalWords)) * 100
	density := (float64(dupes) / float64(len(words))) * 100
	fmt.Printf("Word len: %2d  #Words: %5d  Portion: %2.0f%%  #Anagram clusters: %5d  Density: %2.0f%%\n", len(words[0]), len(words), portion, len(a), density)
}

func main() {
	dicts := dictionaries.AllDicts(".")

	for _, dict := range dicts {
		allWords := dictionaries.LoadFile(dict)
		totalWords := len(allWords)
		lengths := dictionaries.WordLengths(allWords)

		fmt.Println("\nStatistics for", dict)
		fmt.Println("#Words:", totalWords)

		for _, length := range lengths {
			words := dictionaries.FilterByLen(allWords, length)
			a := dictionaries.Anagrams(words)
			stats(totalWords, words, a)
		}
	}
}
