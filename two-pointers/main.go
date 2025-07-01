package main

import (
	"fmt"
)

// Переместить нули в конец массива без сохранения порядка массива
func moveZeroes(nums []int) {
	i := 0
	j := len(nums) - 1
	for {
		if i == len(nums) {
			break
		}

		if i == j {
			break
		}

		if nums[j] == 0 {
			j--
			continue
		}

		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			i++
		}
	}
}

// leetcode - мое решение одно из худших???
// надо решить заново позже
func moveZeroesOrder(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}
}

func isSubsequence(s string, t string) bool {
	// Является ли строка s подстрокой строки t
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	i := 0 // s pointer
	j := 0 // t pointer

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			if i == len(s)-1 {
				return true
			}
			i++
		}

		j++
	}

	return false
}

func maxOperations(nums []int, k int) int {
	// 1679. Max Number of K-Sum Pairs
	// Надо решить через два поинтера
	m := make(map[int]int)
	var counter int
	for _, v := range nums {
		m[v]++
	}

	for _, x := range nums {
		if value, ok := m[k-x]; ok && value > 0 && m[x] > 0 {
			if value == 1 && k-x == x {
				continue
			}
			m[k-x]--
			m[x]--
			counter++
		}
	}

	return counter
}

// Сам не решил, не очень понял как решать через два поинтера
func maxArea(height []int) int {
	// 11. Container With Most Water

	left := 0                // start
	right := len(height) - 1 // end

	var maximum int
	var current int

	for left != right {
		current = (right - left) * (min(height[left], height[right]))
		maximum = max(current, maximum)

		if height[left] <= height[right] {
			left++
			continue
		}

		if height[left] > height[right] {
			right--
			continue
		}
	}

	return maximum
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
