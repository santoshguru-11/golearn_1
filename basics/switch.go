package main

import "fmt"

func main() {
	// Example without fall through
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("Selected fruit is apple.")
	case "banana":
		fmt.Println("Selected fruit is banana.")
	case "orange":
		fmt.Println("Selected fruit is orange.")
	default:
		fmt.Println("Unknown fruit.")
	}
}
