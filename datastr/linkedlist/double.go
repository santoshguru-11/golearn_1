package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) add(data int) {
	newNode := &Node{data: data, next: nil, previous: nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.previous = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

// Other methods for deleting, searching, etc.

func main() {
	list := DoublyLinkedList{}
	list.add(1)
	list.add(2)
	list.add(3)

	// Print the doubly linked list forward
	current := list.head
	fmt.Println("Forward:")
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

	// Print the doubly linked list backward
	fmt.Println("\nBackward:")
	current = list.tail
	for current != nil {
		fmt.Println(current.data)
		current = current.previous
	}
}
