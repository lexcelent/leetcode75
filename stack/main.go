package main

import "fmt"

func removeStars(s string) string {
	result := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			result = result[:len(result)-1]
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}

func main() {
	res := removeStars("leet*code")
	fmt.Println(res)
}
