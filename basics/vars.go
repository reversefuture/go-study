package main

import (
	"fmt"
)

// var
// Can be used inside and outside of functions
// Variable declaration and value assignment can be done separately
// :=
// Can only be used inside functions
// Variable declaration and value assignment cannot be done separately (must be done in the same line)

func mainVar() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	// var a, b, c, d int = 1, 3, 5, 7 // declar multiple
	// var a2, b2 = 6, "Hello"         //type keyword is not specified, you can declare different types of variables on the same line

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

// In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

func mainDefault() {
	var a string // default: ""
	var b int    // default: 0
	var c bool   // default: false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// Go Variable Declaration in a Block
func mainVarBlock() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
