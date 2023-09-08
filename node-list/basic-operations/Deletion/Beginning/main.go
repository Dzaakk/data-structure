package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func deleteFirstNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head.Next
	head.Next = nil

	return newHead
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

	head = deleteFirstNode(head)

	fmt.Println("After deleting the first node:")
	displayList(head)
}
