package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type ListNode struct {
	head *Node
}

func (ln *ListNode) insertAfter(existingNode *Node, val int) {
	if existingNode == nil {
		fmt.Println("Given node doesn't exist.")
		return
	}

	newNode := &Node{Value: val}
	newNode.Next = existingNode.Next
	existingNode.Next = newNode
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

	ListNode.head = &Node{Value: 10, Next: nil}
	ListNode.head.Next = &Node{Value: 21, Next: nil}
	ListNode.head.Next.Next = &Node{Value: 33, Next: nil}

	fmt.Println("Original Linked List:")
	ListNode.display()

	// Inserting a node after the second node
	nodeToInsertAfter := ListNode.head.Next
	ListNode.insertAfter(nodeToInsertAfter, 25)

	fmt.Println("Linked List after inserting a node:")
	ListNode.display()
}
