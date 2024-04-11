package main

import "fmt"

func main() {

	const name = "Harry Potter"
	const is_muggle = false
	const age = 12
	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T", age, age)

}
