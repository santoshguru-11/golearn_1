package main

import (
	"fmt"
)

type Stack struct {
	elements []int
}

// Push an element onto the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop an element from the stack
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1 // Return -1 if stack is empty
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Popped:", stack.Pop()) // Output: Popped: 30
	fmt.Println("Popped:", stack.Pop()) // Output: Popped: 20
	fmt.Println("Popped:", stack.Pop()) // Output: Popped: 10
}
