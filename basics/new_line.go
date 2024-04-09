package main

import (
	"fmt"
)

func main() {
	var (
		name string
		user string
	)
	fmt.Scanf("%v %v", &user, &name)
	fmt.Print(name, "\n")
	fmt.Print(user)
}
