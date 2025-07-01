package main

import (
	"fmt"
	"math"
	"slices"
)

func getSum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func findMaxAverage(nums []int, k int) float64 {
	maximum := math.Inf(-1)

	for i := 0; i <= len(nums)-k; i++ {
		sum := getSum(nums[i : i+k])
		avg := float64(sum) / float64(k)
		if maximum < avg {
			maximum = avg
		}
	}

	return maximum
}

func longestOnes(nums []int, k int) int {
	// 1004. Max Consecutive Ones III
	// Мое решение было очень большим. Это взято с сайта...
	l, r := -1, -1

	for r < len(nums)-1 {
		r++

		if nums[r] == 0 {
			k--
		}

		if k < 0 {
			l++
			if nums[l] == 0 {
				k++
			}
		}

	}

	return r - l
}

func findMaxConsecutiveOnes(nums []int) int {
	// 485. Max Consecutive Ones
	// Можно проще: размер окна = 1 и просто считать последовательно единицы. Обнулять, если 0
	// Отдельно фиксиовать максимальное значение
	l, r := -1, -1
	windowSize := 0 // размер зависит от единиц в окне

	for r < len(nums)-1 {
		r++

		if nums[r] == 0 {
			l = r
		}

		windowSize = max(r-l, windowSize)
	}

	return windowSize
}

func maxPower(s string) int {
	// 1446. Consecutive Characters
	maxChar := 0
	var prevChar rune
	charCount := 0

	for _, v := range s {
		if v != prevChar {
			prevChar = v
			charCount = 1
		} else {
			charCount++
		}
		maxChar = max(maxChar, charCount)
	}

	return maxChar
}

func maxVowels(s string, k int) int {
	// 1456. Maximum Number of Vowels in a Substring of Given Length
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	maximum := 0
	current := 0

	l, r := 0, 0
	for r < len(s) {
		if r-l > k-1 {
			// Контролируем размер окна
			if slices.Contains(vowels, s[l]) {
				current--
			}
			l++
		}

		if slices.Contains(vowels, s[r]) {
			current++
		}

		maximum = max(current, maximum)

		r++
	}

	return maximum
}

func longestSubarray(nums []int) int {
	// 1493. Longest Subarray of 1's After Deleting One Element
	l, r := 0, 0
	k := 1
	deleted := 0

	maxSubarray := 0

	for r < len(nums) {
		if nums[r] == 0 {
			deleted++
		}

		if deleted > k {
			if nums[l] == 0 {
				deleted--
			}
			l++
		}

		maxSubarray = max(r-l-deleted, maxSubarray)

		r++
	}

	return maxSubarray
}

func main() {
	vars := []int{1, 0, 1, 1, 1, 0, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(vars))
}
