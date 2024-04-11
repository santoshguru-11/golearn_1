package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 42
	var s string = strconv.Itoa(i) // convert int to string
	fmt.Printf("%q", s)
}
