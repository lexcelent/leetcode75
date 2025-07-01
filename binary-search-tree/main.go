package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// макс глубина
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	current := root
	for current != nil {
		if val == current.Val {
			return current
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}

func isLeaftNode(root *TreeNode) bool {
	// Нет потомков
	return root.Left == nil && root.Right == nil
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// Рекурсивное удаление, обрати внимание на 3 кейса
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		if isLeaftNode(root) {
			// case 1: Если нода - лист
			return nil
		} else if root.Left == nil {
			// case 2: один узел
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// case 3: два дочерних узла
			node := getMostLeftOfRight(root.Right)
			root.Val = node.Val
			root.Right = deleteNode(root.Right, node.Val)
			return root
		}
	}
}

func getMostLeftOfRight(node *TreeNode) *TreeNode {
	// Самый левый узел правого поддерева
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

func main() {

}
