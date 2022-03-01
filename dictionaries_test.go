package dictionaries

import (
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func equalInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func equalInt2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !equalInt(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestContainsWord(t *testing.T) {
	testCases := []struct {
		words    []string
		target   string
		expected bool
	}{
		{[]string{"psh", "gg", "w"}, "psh", true},
		{[]string{"psh", "gg", "w"}, "gg", true},
		{[]string{"psh", "gg", "w"}, "w", true},
		{[]string{"psh", "gg", "w"}, "s", false},
		{[]string{"psh", "gg", ""}, "", true},
		{[]string{"psh", "gg", "w"}, "", false},
	}

	for _, testCase := range testCases {
		answer := ContainsWord(testCase.words, testCase.target)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v '%s' expected %t, got %t", testCase.words, testCase.target, testCase.expected, answer)
		}
	}
}

func TestFilterByLen(t *testing.T) {
	testCases := []struct {
		words    []string
		len      int
		expected []string
	}{
		{[]string{"psh", "gg", "w"}, 0, []string{}},
		{[]string{"psh", "gg", "w"}, 1, []string{"w"}},
		{[]string{"psh", "gg", "w"}, 2, []string{"gg"}},
		{[]string{"psh", "gg", "w"}, 3, []string{"psh"}},
		{[]string{"psh", "gg", "w"}, 4, []string{}},
	}

	for _, testCase := range testCases {
		answer := FilterByLen(testCase.words, testCase.len)
		if !equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %v %d expected %v, got %v", testCase.words, testCase.len, testCase.expected, answer)
		}
	}
}

func TestSortUnique(t *testing.T) {
	testCases := []struct {
		w        []string
		expected []string
	}{
		{[]string{}, []string{}},
	}

	for _, testCase := range testCases {
		answer := SortUnique(testCase.w)
		if !equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.w, testCase.expected, answer)
		}
	}
}

func TestLetterFrequency(t *testing.T) {
	testCases := []struct {
		m         []string
		expected1 []int
		expected2 [][]int
	}{
		{[]string{}, []int{}, [][]int{}},
	}

	for _, testCase := range testCases {
		answer1, answer2 := LetterFrequency(testCase.m)
		if !equalInt(answer1, testCase.expected1) {
			t.Errorf("ERROR: For %v expected %v %v, got %v %v", testCase.m, testCase.expected1, testCase.expected2, answer1, answer2)
		}
		if !equalInt2(answer2, testCase.expected2) {
			t.Errorf("ERROR: For %v expected %v %v, got %v %v", testCase.m, testCase.expected1, testCase.expected2, answer1, answer2)
		}
	}
}
