package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T", err, err)
}
