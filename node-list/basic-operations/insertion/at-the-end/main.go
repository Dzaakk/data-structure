package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type ListNode struct {
	head *Node
}

func (ln *ListNode) append(val int) {
	newNode := &Node{Value: val}

	if ln.head == nil {
		ln.head = newNode
		return
	}

	current := ln.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
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

	ListNode.append(7)
	ListNode.append(1)
	ListNode.append(3)

	fmt.Println("Linked List after appending nodes :")
	ListNode.display()
}
