/*
2. Queue of Custom Structs
In this example, a queue is implemented using a slice of custom structs:
*/
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Queue struct {
	elements []Person
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value Person) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the first element of the queue
func (q *Queue) Dequeue() Person {
	if len(q.elements) == 0 {
		return Person{} // Return zero value struct if queue is empty
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func main() {
	queue := &Queue{}

	queue.Enqueue(Person{Name: "John", Age: 25})
	queue.Enqueue(Person{Name: "Jane", Age: 30})

	fmt.Println("Dequeued:", queue.Dequeue()) // Output: Dequeued: {John 25}
	fmt.Println("Dequeued:", queue.Dequeue()) // Output: Dequeued: {Jane 30}
}
