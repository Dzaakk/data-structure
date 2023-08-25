package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func Reverse(current, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}

	next := current.Next
	current.Next = prev
	return Reverse(next, current)
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

	reversedHead := Reverse(head, nil)

	fmt.Println("Reversed linked list:")
	displayList(reversedHead)
}
