package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func deleteLastNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	current := head
	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil

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

	fmt.Println("Original linked list:")
	displayList(head)

	head = deleteLastNode(head)

	fmt.Println("Linked list after deleting the last node:")
	displayList(head)
}
