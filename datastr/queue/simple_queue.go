/*
1. Simple Queue Using a Slice
In this example, a queue is implemented using a slice of integers:
*/

package main

import (
	"fmt"
)

type Queue struct {
	elements []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the first element of the queue
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		return -1 // Return -1 if queue is empty
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func main() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Dequeued:", queue.Dequeue()) // Output: Dequeued: 10
	fmt.Println("Dequeued:", queue.Dequeue()) // Output: Dequeued: 20
	fmt.Println("Dequeued:", queue.Dequeue()) // Output: Dequeued: 30
}
