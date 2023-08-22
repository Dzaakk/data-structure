package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func CountNodes(head *ListNode) int {
	count := 0
	current := head

	for current != nil {
		count++
		current = current.Next
	}

	return count
}
func main() {
	var head *ListNode

	head = &ListNode{Value: 10, Next: nil}
	head.Next = &ListNode{Value: 14, Next: nil}
	head.Next.Next = &ListNode{Value: 20, Next: nil}

	nodeCount := CountNodes(head)
	fmt.Printf("Number of nodes in the linked list: %d\n", nodeCount)
}
