package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 1431. Kids With the Greatest Number of Candies
	result := make([]bool, len(candies))
	greatest := slices.Max(candies)

	for i, v := range candies {
		if v+extraCandies >= greatest {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

func isPalindrome(s string) bool {
	// 125. Valid Palindrome

	// 1. Избавляемся от лишних символов
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}

	// Проверяем палиндром
	str := strings.ToLower(sb.String())
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(" "))
}
