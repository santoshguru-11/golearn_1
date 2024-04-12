package main

import "fmt"

func main() {
	// Declaring variables
	var a int
	a = 10

	// Assignment operators
	fmt.Println("Assignment Operators:")
	fmt.Println("Initial value of a:", a)

	a += 5 // Equivalent to: a = a + 5
	fmt.Println("a += 5:", a)

	a -= 3 // Equivalent to: a = a - 3
	fmt.Println("a -= 3:", a)

	a *= 2 // Equivalent to: a = a * 2
	fmt.Println("a *= 2:", a)

	a /= 4 // Equivalent to: a = a / 4
	fmt.Println("a /= 4:", a)

	a %= 2 // Equivalent to: a = a % 2
	fmt.Println("a %= 2:", a)

	a <<= 1 // Equivalent to: a = a << 1 (left shift)
	fmt.Println("a <<= 1:", a)

	a >>= 2 // Equivalent to: a = a >> 2 (right shift)
	fmt.Println("a >>= 2:", a)

	a &= 3 // Equivalent to: a = a & 3 (bitwise AND)
	fmt.Println("a &= 3:", a)

	a |= 5 // Equivalent to: a = a | 5 (bitwise OR)
	fmt.Println("a |= 5:", a)

	a ^= 6 // Equivalent to: a = a ^ 6 (bitwise XOR)
	fmt.Println("a ^= 6:", a)
}
