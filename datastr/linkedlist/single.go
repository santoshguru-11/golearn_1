package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) add(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Other methods for deleting, searching, etc.

func main() {
	list := LinkedList{}
	list.add(1)
	list.add(2)
	list.add(3)

	// Print the linked list
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
