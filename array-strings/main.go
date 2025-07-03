package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
)

func mergeAlternately(word1 string, word2 string) string {
	// 1768. Merge Strings Alternately
	// Мой второй вариант
	var sb strings.Builder

	for i, j := 0, 0; i+j < len(word1)+len(word2); {
		if i < len(word1) {
			sb.WriteByte(word1[i])
			i++
		}

		if j < len(word2) {
			sb.WriteByte(word2[j])
			j++
		}
	}

	return sb.String()
}

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

func gcd(a, b int) int {
	// Наибольший общий делитель. Алгоритм Евклида
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	// 1071. Greatest Common Divisor of Strings
	// Повторить бы. Или запомнить
	if str1+str2 != str2+str1 {
		return ""
	}

	g := gcd(len(str1), len(str2))

	return str1[:g]
}

func findGCD(nums []int) int {
	// 1979. Find Greatest Common Divisor of Array
	var smallest, biggest = math.MaxInt, math.MinInt

	for _, v := range nums {
		smallest = min(smallest, v)
		biggest = max(biggest, v)
	}

	return gcd(smallest, biggest)
}

func increasingTriplet(nums []int) bool {
	// 334. Increasing Triplet Subsequence
	lst := []int{nums[0]}
	currentLen := len(lst)

	conditionLength := 3

	for _, num := range nums[1:] {
		if num > lst[len(lst)-1] {
			lst = append(lst, num)
			currentLen++
		} else {
			// Меняем последовательность
			indx := 0
			for indx < len(lst) && lst[indx] < num {
				indx++
			}
			lst[indx] = num
		}

		if currentLen >= conditionLength {
			return true
		}
	}

	return false
}

func lengthOfLIS(nums []int) int {
	// 300. Longest Increasing Subsequence
	// Вроде алгоритм запомнил. Нужно будет попробовать еще
	lst := []int{nums[0]}
	maxLen := len(lst)

	for _, num := range nums[1:] {
		if num > lst[len(lst)-1] {
			lst = append(lst, num)
			maxLen++
		} else {
			// Меняем последовательность
			indx := 0
			for indx < len(lst) && lst[indx] < num {
				indx++
			}
			lst[indx] = num
		}
	}

	return maxLen
}

func main() {
	fmt.Println(isPalindrome(" "))
	// fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
	// fmt.Println(gcdOfStrings("ABAB", "ABABAB"))
	fmt.Println(lengthOfLIS([]int{1, 0, 1, 2, 3, 0, 4, 1, 5}))
}
