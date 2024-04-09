package main

import (
	"fmt"
)

func main() {
	var name string = "Joe"
	var i int = 78
	fmt.Printf("Hey, %v! You have scored %d/100 in Physics", name, i)
}

/*
Verb Description
%v default format
%T type of the value
%d integers
%c character
%q quoted characters/string
%s plain string
%t true or false
%f floating numbers
%.2f floating numbers upto 2 decimal places
*/
