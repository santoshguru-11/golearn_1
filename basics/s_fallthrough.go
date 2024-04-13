package main

import "fmt"

func main() {
	// Example with fall through
	fruit := "banana"

	switch fruit {
	case "apple":
		fmt.Println("Selected fruit is apple.")
		fallthrough
	case "banana":
		fmt.Println("Selected fruit is banana.")
		fallthrough
	case "orange":
		fmt.Println("Selected fruit is orange.")
	default:
		fmt.Println("Unknown fruit.")
	}
}
