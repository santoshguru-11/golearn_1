/*
5. Stack Using a Linked List
In this example, a stack is implemented using a linked list:
*/
package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

// Push an element onto the stack
func (s *Stack) Push(value int) {
	newNode := &Node{value: value}
	newNode.next = s.top
	s.top = newNode
}

// Pop an element from the stack
func (s *Stack) Pop() int {
	if s.top == nil {
		return -1 // Return -1 if stack is empty
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Popped:", stack.Pop()) // Output: Popped: 30
	fmt.Println("Popped:", stack.Pop()) // Output: Popped: 20
}
