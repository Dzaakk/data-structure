package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func (head *ListNode) Search(key int) bool {
	if head == nil {
		return false
	}

	if head.Value == key {
		return true
	}

	return head.Next.Search(key)
}

func main() {
	var head *ListNode

	head = &ListNode{Value: 10, Next: nil}
	head.Next = &ListNode{Value: 14, Next: nil}
	head.Next.Next = &ListNode{Value: 20, Next: nil}

	ValueToSearch := 20
	found := head.Search(ValueToSearch)

	if found {
		fmt.Printf("Value %d found in the linked list.\n", ValueToSearch)
	} else {
		fmt.Printf("Value %d not found in the linked list.\n", ValueToSearch)
	}
}
