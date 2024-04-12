package main

import "fmt"

func main() {
	// Declaring variables
	var a, b uint8
	a = 5
	b = 3

	// Bitwise operators
	fmt.Println("Bitwise Operators:")
	fmt.Println("a & b:", a&b)   // Bitwise AND
	fmt.Println("a | b:", a|b)   // Bitwise OR
	fmt.Println("a ^ b:", a^b)   // Bitwise XOR
	fmt.Println("a &^ b:", a&^b) // Bit clear (AND NOT)
	fmt.Println("a << 1:", a<<1) // Left shift
	fmt.Println("b >> 1:", b>>1) // Right shift
}
