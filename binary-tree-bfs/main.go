package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	// 1161. Maximum Level Sum of a Binary Tree
	if root == nil {
		return 0
	}

	// База алгоритма
	queue := []*TreeNode{root}

	// Храним суммы
	sums := []int{}

	// Сумма элементов
	countSum := func(lst []*TreeNode) int {
		total := 0
		for _, v := range lst {
			total += v.Val
		}
		return total
	}

	for len(queue) != 0 {
		// Обрабатываем текущий уровень
		levelSize := len(queue)

		sums = append(sums, countSum(queue))

		for i := 0; i < levelSize; i++ {
			// Проходимся по каждому элементу одного уровня
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}

	maxValueIndex := func(vals []int) int {
		maxVal := math.MinInt
		var idx int
		for i, v := range vals {
			if maxVal < v {
				maxVal, idx = v, i
			}
		}
		return idx
	}

	return maxValueIndex(sums) + 1
}

func rightSideView(root *TreeNode) []int {
	// 199. Binary Tree Right Side View
	var result []int
	if root == nil {
		return result
	}

	// База
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		levelSize := len(queue)
		result = append(result, queue[len(queue)-1].Val)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

func main() {

}
