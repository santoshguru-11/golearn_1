package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SortedLinkedList struct {
	head *Node
}

func (list *SortedLinkedList) add(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil || list.head.data >= data {
		newNode.next = list.head
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil && current.next.data < data {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
}

// Other methods for deleting, searching, etc.

func main() {
	list := SortedLinkedList{}
	list.add(3)
	list.add(1)
	list.add(2)

	// Print the sorted linked list
	current := list.head
	fmt.Println("Sorted Linked List:")
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
