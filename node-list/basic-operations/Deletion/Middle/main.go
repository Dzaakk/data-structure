package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func deleteNodeFromMiddle(head *ListNode, target int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Value == target {
		return head.Next
	}

	prev := head
	current := head.Next
	for current != nil && current.Value != target {
		prev = current
		current = current.Next
	}

	if current != nil {
		prev.Next = current.Next
		current.Next = nil
	}

	return head
}

func createLinkedList(values []int) *ListNode {
	var head *ListNode
	var current *ListNode

	for _, value := range values {
		newNode := &ListNode{Value: value}
		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = current.Next
		}
	}
	return head
}

func displayList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	values := []int{10, 20, 30, 40, 50}
	head := createLinkedList(values)

	fmt.Println("Original Linked List:")
	displayList(head)

	targetValue := 30
	head = deleteNodeFromMiddle(head, targetValue)

	fmt.Println("After deleting the first node:")
	displayList(head)
}
