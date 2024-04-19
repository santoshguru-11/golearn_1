package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Stack struct {
	elements []Person
}

// Push an element onto the stack
func (s *Stack) Push(value Person) {
	s.elements = append(s.elements, value)
}

// Pop an element from the stack
func (s *Stack) Pop() Person {
	if len(s.elements) == 0 {
		return Person{} // Return zero value struct if stack is empty
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value
}

func main() {
	stack := &Stack{}

	stack.Push(Person{Name: "John", Age: 25})
	stack.Push(Person{Name: "Jane", Age: 30})

	fmt.Println("Popped:", stack.Pop()) // Output: Popped: {Jane 30}
	fmt.Println("Popped:", stack.Pop()) // Output: Popped: {John 25}
}
