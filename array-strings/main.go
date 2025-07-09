package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// 1768. Merge Strings Alternately
// Мой второй вариант
func mergeAlternately(word1 string, word2 string) string {
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

// 1431. Kids With the Greatest Number of Candies
func kidsWithCandies(candies []int, extraCandies int) []bool {
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

// 125. Valid Palindrome
func isPalindrome(s string) bool {

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

// 605. Can Place Flowers
// На будущее: переписать код под switch-case
func canPlaceFlowers(flowerbed []int, n int) bool {
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

// 238. Product of Array Except Self
// Нужно будет повторить
func productExceptSelf(nums []int) []int {
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

// Наибольший общий делитель. Алгоритм Евклида
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 1071. Greatest Common Divisor of Strings
// Повторить бы. Или запомнить
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	g := gcd(len(str1), len(str2))

	return str1[:g]
}

// 1979. Find Greatest Common Divisor of Array
func findGCD(nums []int) int {
	var smallest, biggest = math.MaxInt, math.MinInt

	for _, v := range nums {
		smallest = min(smallest, v)
		biggest = max(biggest, v)
	}

	return gcd(smallest, biggest)
}

// 334. Increasing Triplet Subsequence
func increasingTriplet(nums []int) bool {
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

// 300. Longest Increasing Subsequence
// Вроде алгоритм запомнил. Нужно будет попробовать еще
func lengthOfLIS(nums []int) int {
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

// 443. String Compression
func compress(chars []byte) int {
	var count = 1
	var leftPtr = 1

	// Перевод числа в строку и запись в strings.Builder
	writeNumber := func(i int) {
		number := strconv.Itoa(i)
		for i := 0; i < len(number); i++ {
			chars[leftPtr] = number[i]
			leftPtr++
		}
	}

	// Найти группу
	for i := 1; i < len(chars); i++ {

		// Проверяем в какой мы группе
		if chars[i] == chars[i-1] {
			count++
		} else {
			// Если группа изменилась, то
			// 1. Записываем количество в прошлой группе и обнуляем
			if count != 1 {
				writeNumber(count)
			}
			count = 1

			// 2. Записываем букву новой группы
			chars[leftPtr] = chars[i]
			leftPtr++
		}
	}

	if count > 1 {
		writeNumber(count)
	}

	return len(chars[:leftPtr])
}

func main() {
	// fmt.Println(isPalindrome(" "))
	// fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
	// fmt.Println(gcdOfStrings("ABAB", "ABABAB"))
	// fmt.Println(lengthOfLIS([]int{1, 0, 1, 2, 3, 0, 4, 1, 5}))
	// compress([]byte{'a', 'a', 'a', 'a', 'a'})
	fmt.Println(compress([]byte{'a'}))
}
