package main

import (
	"fmt"
	"slices"
)

func uniqueOccurrences(arr []int) bool {
	if len(arr) == 1 {
		return true
	}
	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	fmt.Println(m)
	occurrences := make(map[int]bool)
	for _, v := range m {
		if _, have := occurrences[v]; have {
			return false
		}
		occurrences[v] = true
	}

	return true
}

func compareSet(word1, word2 string) bool {
	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)

	// build
	for _, v := range word1 {
		m1[v] = true
	}

	for _, v := range word2 {
		m2[v] = true
	}

	// compare
	for _, v := range word1 {
		if _, ok := m2[v]; !ok {
			return false
		}
	}

	return true
}

func compareFreq(word1, word2 string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	// build
	for _, v := range word1 {
		m1[v]++
	}

	for _, v := range word2 {
		m2[v]++
	}

	// get keys as slice
	var s1, s2 []int
	for _, v := range m1 {
		s1 = append(s1, v)
	}

	for _, v := range m2 {
		s2 = append(s2, v)
	}

	slices.Sort(s1)
	slices.Sort(s2)

	// compare
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func closeStrings(word1 string, word2 string) bool {
	// 1657. Determine if Two Strings Are Close

	// Should be same len
	if len(word1) != len(word2) {
		return false
	}

	// Should be equal alphabet
	if !compareSet(word1, word2) {
		return false
	}

	// Should be equal frequences
	if !compareFreq(word1, word2) {
		return false
	}

	return true
}

func rowCount(row []int, grid [][]int) int {
	var count int
	var column = make([]int, len(grid))

	compareSlices := func(s1, s2 []int) bool {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}

		return true
	}

	for i := range grid {

		// build column
		for j := range grid {
			column[j] = grid[j][i]
		}

		if compareSlices(row, column) {
			count++
		}
	}

	return count
}

func equalPairs(grid [][]int) int {
	// 2352. Equal Row and Column Pairs
	// Понимаю, что можно решить двумя способами. Выбрал без мапы,
	// т.к. не хочу выделять память
	var equals int

	for _, row := range grid {
		equals += rowCount(row, grid)
	}

	return equals
}

func main() {
	vals := []int{2, 3, 3, 3, 4, 4, 5, 5, 5, 5}
	// vals1 := []int{1, 2}
	fmt.Println(uniqueOccurrences(vals))
}
