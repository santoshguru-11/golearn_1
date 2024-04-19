/*
3. Stack with Capacity
In this example, the stack is implemented with an initial capacity to prevent frequent slice resizing:
*/
package main

import (
	"fmt"
)

type Stack struct {
	elements []int
	capacity int
}

// Push an element onto the stack
func (s *Stack) Push(value int) {
	if len(s.elements) == s.capacity {
		newCapacity := s.capacity * 2
		s.elements = append(s.elements, make([]int, newCapacity)...)
		s.capacity = newCapacity
	}
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
	stack := &Stack{
		capacity: 4,
	}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Println("Popped:", stack.Pop()) // Output: Popped: 50
}
