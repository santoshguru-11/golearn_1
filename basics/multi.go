package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	fmt.Print("Enter a string and a number: ", "\n")

	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count : ", count, "\n")
	fmt.Println("error: ", err, "\n")
	fmt.Println("a: ", a, "\n")
	fmt.Println("b: ", b, "\n")

}
