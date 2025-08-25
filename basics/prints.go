package main

import (
	"fmt"
)

func mainPrint() {
	var i, j string = "Hello", "World"

	fmt.Print(i)         // same line, no newline
	fmt.Print(j, "\n")   // newline is added at the end
	fmt.Print(i, " ", j) // print multiple in same line
	fmt.Println(i, j)    // newline is added at the end

	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign
	// more: https://www.w3schools.com/go/go_formatting_verbs.php
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
