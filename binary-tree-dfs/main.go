package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsModified(root *TreeNode) []int {
	// Возвращает список листьев дерева, используя DFS

	// База
	stack := []*TreeNode{root}

	var result []int

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
		}
	}

	return result
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 872. Leaf-Similar Trees

	s1 := dfsModified(root1)
	s2 := dfsModified(root2)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func main() {

}
