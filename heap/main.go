package main

import (
	"fmt"
)

// left = 2 * i + 1
// right = 2 * i + 2
// parent = (i - 1) / 2

// insert - add to the end, bubble up until it's ok
// delete - swap first and last elem, delete last elem, bubble down first elem until it's ok
// heapify - create from unsorted array

// for max heap
func heapify(nums []int, indx int) {
	largest := indx
	left := 2*indx + 1
	right := 2*indx + 2

	size := len(nums)

	// for max heap
	if left < size && nums[left] > nums[largest] {
		largest = left
	}
	if right < size && nums[right] > nums[largest] {
		largest = right
	}

	if largest != indx {
		nums[indx], nums[largest] = nums[largest], nums[indx]
		heapify(nums, largest)
	}
}

// actually delete root
func deleteRoot(nums []int) []int {
	nums[0] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	heapify(nums, 0)

	return nums
}

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i)
	}

	for i := 0; i < k-1; i++ {
		nums = deleteRoot(nums)
	}

	return nums[0]
}

func main() {
	vals := []int{-1, 2, 0}
	fmt.Println(findKthLargest(vals, 3))
}
