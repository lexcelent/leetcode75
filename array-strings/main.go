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

func canPlaceFlowers(flowerbed []int, n int) bool {
	// 605. Can Place Flowers
	// На будущее: переписать код под switch-case

	if n == 0 {
		// We can place zero flowers
		return true
	}

	placeFlower := func(i int) {
		flowerbed[i] = 1
		n--
	}

	if len(flowerbed) == 1 {
		// Process only one element
		if flowerbed[0] == 0 {
			placeFlower(0)
		}
		return n == 0
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if i == 0 {
			// Правило для левой границы
			if flowerbed[i+1] == 0 {
				placeFlower(i)
			}
			continue
		}

		if i == len(flowerbed)-1 {
			// Правило для правой границы
			if flowerbed[i-1] == 0 {
				placeFlower(i)
			}
			continue
		}

		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			placeFlower(i)
		}

		if n <= 0 {
			return true
		}
	}

	return n == 0
}

func productExceptSelf(nums []int) []int {
	// 238. Product of Array Except Self
	// Нужно будет повторить

	res := make([]int, len(nums))

	for i := range res {
		res[i] = 1
	}

	var prefix, postfix = 1, 1
	for i := 0; i < len(nums); i++ {
		res[i] *= prefix
		prefix *= nums[i]
		fmt.Printf("i: %d\nres: %v\nprefix: %d\n\n", i, res, prefix)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}

func main() {
	fmt.Println(isPalindrome(" "))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}
