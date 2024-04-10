package main

import "fmt"

func main() {

	var i int
	fmt.Printf("%d", i, "\n")
	var fl float64
	fmt.Printf("%.2f", fl, "\n")
	var name string
	fmt.Print("Enter your name: ", "\n")
	fmt.Scanf("%s", &name)

	fmt.Println("Hey there, ", name, "\n")

	var name_1 string
	var is_muggle bool
	fmt.Print("Enter your name & are you a muggle: ", "\n")
	fmt.Scanf("%s %t", &name_1, &is_muggle)
	fmt.Println(name_1, is_muggle, "\n")

}
