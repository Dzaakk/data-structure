package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type ListNode struct {
	head *Node
}

func (ln *ListNode) prepend(val int) {
	newNode := &Node{Value: val, Next: ln.head}
	ln.head = newNode
}

func (ln *ListNode) display() {
	current := ln.head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}

	fmt.Println("nil")
}

func main() {
	ListNode := &ListNode{}

	ListNode.prepend(7)
	ListNode.prepend(1)
	ListNode.prepend(3)

	fmt.Println("Linked List after inserting nodes from the beginning:")
	ListNode.display()
}
