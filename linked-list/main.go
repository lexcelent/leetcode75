package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverse list task, very effective solution, but could be better
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	var next *ListNode

	for {
		if curr == nil {
			break
		}

		next = curr.Next
		curr.Next = prev

		prev, curr = curr, next
	}

	return prev
}

// get linked list lenght
func getLenght(head *ListNode) int {
	count := 0
	for current := head; current != nil; current = current.Next {
		count++
	}
	return count
}

// удалить середину. Можно лучше, через медленный/быстрый указатель
func deleteMiddle(head *ListNode) *ListNode {
	total := getLenght(head)
	if total == 1 {
		return nil
	}

	middle := total / 2

	idx := 0
	for current := head; current != nil; current = current.Next {
		if idx+1 == middle {
			current.Next = current.Next.Next
			break
		}

		idx++
	}

	return head
}

func showList(head *ListNode) {
	idx := 0
	for current := head; current != nil; current = current.Next {
		fmt.Printf("idx: %d, %#v\n", idx, current)
		idx++
	}
}

func showEvens(head *ListNode) {
	idx := 0
	for current := head; current != nil; current = current.Next {
		if idx%2 == 0 {
			fmt.Printf("idx: %d, %#v\n", idx, current)
		}
		idx++
	}
}

func showOdds(head *ListNode) {
	idx := 0
	for current := head; current != nil; current = current.Next {
		if idx%2 != 0 {
			fmt.Printf("idx: %d, %#v\n", idx, current)
		}
		idx++
	}
}

// Нечетный четный порядок.
// Снача нечетный.некст ссылается на четный. Потом четный.некст ссылается на нечетный.
// Затем ссылаться хвостом нечетных на голову четных
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddTail := head
	evenHead := head.Next
	evenTail := head.Next

	for evenTail != nil && evenTail.Next != nil {
		oddTail.Next = evenTail.Next
		oddTail = oddTail.Next

		evenTail.Next = oddTail.Next
		evenTail = evenTail.Next
	}

	oddTail.Next = evenHead

	return head
}

func newNode(x int) *ListNode {
	return &ListNode{x, nil}
}

func newList(nums ...int) *ListNode {
	var head *ListNode
	current := head
	for _, num := range nums {
		if head == nil {
			head = newNode(num)
			current = head
			continue
		}

		current.Next = newNode(num)
		current = current.Next
	}

	return head
}

// На основе готовых решений, но можно и без них, если уметь...
// fast and slow pointers
func pairSum(head *ListNode) int {
	total := getLenght(head)
	if total == 2 {
		return head.Val + head.Next.Val
	}

	middleIdx := total / 2
	idx := 0
	maxValue := 0

	// get middle node
	middle := head
	for {
		if idx == middleIdx {
			break
		}
		middle = middle.Next
		idx++
	}
	middle = reverseList(middle)

	current := head
	for current != nil && middle != nil {
		if current.Val+middle.Val > maxValue {
			maxValue = current.Val + middle.Val
		}
		current = current.Next
		middle = middle.Next
	}

	return maxValue
}

// найти середину
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	lst := newList(vals...)
	res := oddEvenList(lst)

	showList(res)
	// showEvens(lst)
	// showOdds(lst)
}
