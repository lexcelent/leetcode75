package main

import (
	"fmt"
	"slices"
)

/*
Бинарные поиск - это про "сделать две границы - сужать их на 50%"
*/

var pick = 1

func guess(x int) int {
	if x < pick {
		return 1
	}

	if x > pick {
		return -1
	}

	return 0
}

func guessNumber(n int) int {
	if n == 1 {
		return 1
	}

	var num int
	left := 1
	right := n
	var mid int

	for left < right {
		mid = left + ((right - left) >> 1)

		num = guess(mid)
		if num <= 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func findPeakElement(nums []int) int {
	// 162. Find Peak Element
	// Не догадался...
	l, r := 0, len(nums)-1

	for l < r {
		middle := l + (r-l)/2

		if nums[middle] > nums[middle+1] {
			r = middle

		} else {
			l = middle + 1
		}
	}

	return l
}

func findPeaks(mountain []int) []int {
	// 2951. Find the Peaks
	peaks := []int{}

	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			peaks = append(peaks, i)
		}
	}

	return peaks
}

func peakIndexInMountainArray(arr []int) int {
	// 852. Peak Index in a Mountain Array
	l, r := 0, len(arr)-1
	var mid int

	for l < r {
		mid = l + (r-l)/2

		if arr[mid] > arr[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	// 2300. Successful Pairs of Spells and Potions
	pairs := make([]int, len(spells))
	slices.Sort(potions)

	for i, spell := range spells {
		// Алгоритм бинарного поиска
		// Ищем индекс в пределах массива с зельями
		left, right := 0, len(potions)-1
		var mid int
		for left < right {
			// Ищем наименьшее достаточное зелье
			mid = left + (right-left)/2

			if spell*potions[mid] < int(success) {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// С этого момента left указывает на нужное зелье.
		// Если заклинание*достаточное зелье = недостаточно, то ставим 0
		// Иначе все остальные зелья будут успешными
		if spell*potions[left] < int(success) {
			pairs[i] = 0
		} else {
			pairs[i] = len(potions) - left
		}
	}

	return pairs
}

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}
