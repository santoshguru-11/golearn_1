package main

import (
	"fmt"
)

func main() {
	// Boolean
	var b bool
	fmt.Print("Enter a boolean value (true/false): ")
	fmt.Scanf("%t", &b)
	fmt.Println("Boolean:", b)

	// Numeric Types
	var i int
	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d", &i)
	fmt.Println("Integer:", i)

	var f float64
	fmt.Print("Enter a float: ")
	fmt.Scanf("%f", &f)
	fmt.Println("Float:", f)

	var real, imag float64
	fmt.Print("Enter a complex number (real part): ")
	fmt.Scanf("%f", &real)
	fmt.Print("Enter a complex number (imaginary part): ")
	fmt.Scanf("%f", &imag)
	c := complex(real, imag)
	fmt.Println("Complex:", c)

	// String
	var s string
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &s)
	fmt.Println("String:", s)

	// Array
	var arr [3]int
	fmt.Print("Enter three integers for array separated by space: ")
	fmt.Scanf("%d %d %d", &arr[0], &arr[1], &arr[2])
	fmt.Println("Array:", arr)

	// Slice
	var slice []int
	fmt.Print("Enter the size of slice: ")
	var size int
	fmt.Scanf("%d", &size)
	slice = make([]int, size)
	fmt.Printf("Enter %d integers for slice separated by space: ", size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &slice[i])
	}
	fmt.Println("Slice:", slice)

	// Map
	var m map[string]int
	m = make(map[string]int)
	var key string
	var value int
	fmt.Print("Enter the size of map: ")
	fmt.Scanf("%d", &size)
	fmt.Printf("Enter %d key-value pairs for map (key value separated by space):\n", size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%s %d", &key, &value)
		m[key] = value
	}
	fmt.Println("Map:", m)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	var p Person
	fmt.Print("Enter name and age separated by space: ")
	fmt.Scanf("%s %d", &p.Name, &p.Age)
	fmt.Println("Struct:", p)

	// Pointer
	var ptr *int
	i = 0
	ptr = &i
	fmt.Print("Enter an integer value to change pointer value: ")
	fmt.Scanf("%d", &i)
	fmt.Println("Pointer:", ptr)
}

/*

Enter a boolean value (true/false): true
Boolean: true

Enter an integer: 42
Integer: 42

Enter a float: 3.14
Float: 3.14

Enter a complex number (real part): 5
Enter a complex number (imaginary part): 7
Complex: (5+7i)

Enter a string: Hello, Go!
String: Hello, Go!

Enter three integers for array separated by space: 1 2 3
Array: [1 2 3]

Enter the size of slice: 3
Enter 3 integers for slice separated by space: 4 5 6
Slice: [4 5 6]

Enter the size of map: 2
Enter 2 key-value pairs for map (key value separated by space):
a 1
b 2
Map: map[a:1 b:2]

Enter name and age separated by space: Alice 30
Struct: {Alice 30}

Enter an integer value to change pointer value: 100
Pointer: 0x40e020
*/
