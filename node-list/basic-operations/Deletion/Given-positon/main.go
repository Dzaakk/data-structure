package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

func deleteNodeAtPosition(head *ListNode, position int) *ListNode {
	if head == nil {
		return nil
	}

	if position == 0 {
		return head.next
	}

	current := head
	for i := 0; i < position-2 && current.next != nil; i++ {
		current = current.next
	}

	if current.next == nil {
		return head
	}

	current.next = current.next.next

	return head
}
func createLinkedList(values []int) *ListNode {
	var head *ListNode
	var current *ListNode

	for _, value := range values {
		newNode := &ListNode{value: value}
		if head == nil {
			head = newNode
			current = head
		} else {
			current.next = newNode
			current = current.next
		}
	}

	return head
}

func displayList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	values := []int{10, 20, 30, 40, 50}
	head := createLinkedList(values)

	fmt.Println("Original linked list:")
	displayList(head)

	positionToDelete := 3
	head = deleteNodeAtPosition(head, positionToDelete)

	fmt.Printf("Linked list after deleting node at position %d:\n", positionToDelete)
	displayList(head)
}
