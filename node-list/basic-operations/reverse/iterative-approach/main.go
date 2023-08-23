package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func Reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	return prev
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
	values := []int{10, 20, 30}
	head := createLinkedList(values)

	fmt.Println("Original linked list:")
	displayList(head)

	reversedHead := Reverse(head)

	fmt.Println("Reversed linked list:")
	displayList(reversedHead)
}
