package dictionaries

// go fmt ./... && go vet ./... && go test

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var ()

// AllDicts returns a slice of filenames of all of the dictionaries.
func AllDicts(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	dicts := []string{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".dict") {
			continue
		}
		dicts = append(dicts, file.Name())
	}

	return dicts
}

// ContainsWord returns true if the word is in the slice
func ContainsWord(words []string, target string) bool {
	for _, word := range words {
		if word == target {
			return true
		}
	}

	return false
}

// LoadFile returns the contents of a file split on newlines
func LoadFile(file string) []string {
	raw, _ := ioutil.ReadFile(file)
	words := strings.Split(string(raw), "\n")

	// Strip trailing blank lines
	for words[len(words)-1] == "" {
		words = words[:len(words)-1]
	}

	return words
}

// FilterByLen returns all words of a given len from a given list of words
func FilterByLen(words []string, l int) []string {
	matches := []string{}

	for _, word := range words {
		if len(word) == l {
			matches = append(matches, word)
		}
	}

	return matches
}

// SortUnique sorts a list and removes any duplicates
func SortUnique(s []string) []string {
	if len(s) <= 0 {
		return []string{}
	}

	// Make a copy so we do not corrupt the backing array of s
	s2 := make([]string, len(s))
	copy(s2, s)

	sort.Strings(s2)

	last := s2[0]
	for i := 1; i < len(s2); {
		if s2[i] == last {
			// Delete this duplicate
			s2 = append(s2[:i], s2[i+1:]...)
			continue
		}
		last = s2[i]
		i++
	}

	return s2
}

// LetterFrequency returns maps of the frequency of letters in the given words
// overall and by letter position
// TODO: Extend this to handles slices with varying word lengths
func LetterFrequency(words []string) ([]int, [][]int) {
	if len(words) == 0 {
		return nil, nil
	}

	letterLen := 256 // This is too many, but it is fast
	positions := len(words[0])
	lFreq := make([]int, letterLen)
	lbpFreq := make([][]int, positions)

	for i := range lbpFreq {
		lbpFreq[i] = make([]int, letterLen)
	}

	for _, match := range words {
		for i := 0; i < positions; i++ {
			lbpFreq[i][match[i]]++
			lFreq[match[i]]++
		}
	}

	return lFreq, lbpFreq
}

// PrettyPrintFreq returns a formatted string representation of the given frequencies
func PrettyPrintFreq(f []int) string {
	out := []string{}

	for key, val := range f {
		if val == 0 {
			continue
		}
		str := fmt.Sprintf("%c:%3d", key, val)
		out = append(out, str)
	}

	return fmt.Sprintf("  %s\n", strings.Join(SortUnique(out), " "))
}

// Anagrams returns a map of anagram clusters and the count of words in the cluster.
func Anagrams(words []string) map[string]int {
	anagrams := map[string]int{}

	for _, word := range words {
		r := []rune(word)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		anagrams[string(r)]++
	}

	return anagrams
}

// WordLegths returns a sorted slice of all word lengths.
func WordLengths(words []string) []int {
	sizes := map[int]bool{}
	for _, word := range words {
		sizes[len(word)] = true
	}

	lengths := []int{}
	for length, _ := range sizes {
		lengths = append(lengths, length)
	}

	sort.Ints(lengths)

	return lengths
}
