package main

import (
	"fmt"
	"math/bits"
)

func countBits(n int) []int {
	// 338. Counting Bits
	var res = make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = bits.OnesCount(uint(i))
	}

	return res
}

func BrianKernighanAlgorithm(n int) int {
	// Алгоритм нашел в интернете
	count := 0
	for n != 0 {
		count++
		fmt.Printf("%b\n", n)
		n = n & (n - 1)
	}

	return count
}

func singleNumber(nums []int) int {
	// 136. Single Number
	// По сути наугад сделал. Посмотрел подсказку про XOR
	var res int
	for _, v := range nums {
		res = res ^ v
	}

	return res
}

func main() {
	// fmt.Println(countBits(5))
	// BrianKernighanAlgorithm(10)
	singleNumber([]int{1, 2, 3, 4, 3, 2, 1})
}
