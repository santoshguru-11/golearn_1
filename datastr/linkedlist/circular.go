package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

func (list *CircularLinkedList) add(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		newNode.next = newNode // Point to itself to form a circle
	} else {
		current := list.head
		for current.next != list.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = list.head
	}
}

// Other methods for deleting, searching, etc.

func main() {
	list := CircularLinkedList{}
	list.add(1)
	list.add(2)
	list.add(3)

	// Print the circular linked list
	current := list.head
	fmt.Println("Circular Linked List:")
	for {
		fmt.Println(current.data)
		current = current.next
		if current == list.head {
			break
		}
	}
}
