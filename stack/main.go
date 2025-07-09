package main

import "fmt"

// 2390. Removing Stars From a String
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

type Stack struct {
	ar []int
}

func NewStack() *Stack {
	return &Stack{ar: make([]int, 0)}
}

func (s *Stack) pushBack(x int) {
	s.ar = append(s.ar, x)
}

func (s *Stack) popBack() {
	s.ar = s.ar[:len(s.ar)-1]
}

func (s *Stack) lastElem() int {
	return s.ar[len(s.ar)-1]
}

// 735. Asteroid Collision
func asteroidCollision(asteroids []int) []int {
	stack := NewStack()

	for _, asteroid := range asteroids {
		// Если летит вправо, то добавляем
		if asteroid > 0 {
			stack.pushBack(asteroid)
		} else {
			// Пока последний астероид летит вправо и он меньше астероида, который летит влево
			for len(stack.ar) != 0 && stack.lastElem() > 0 && stack.lastElem() < -asteroid {
				stack.popBack() // Уничтожаем астероиды
			}

			if len(stack.ar) != 0 && stack.lastElem() == -asteroid {
				// Если астероиды равны, то уничтожаем последний добавленный
				stack.popBack()
			} else if len(stack.ar) == 0 || stack.lastElem() < 0 {
				// Если стек пустой или последний добавленный элемент летит влево, то добавляем
				stack.pushBack(asteroid)
			}
		}
	}

	return stack.ar
}

func main() {
	res := removeStars("leet*code")
	fmt.Println(res)
}
