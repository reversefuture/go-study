package main

import (
	"fmt"
)

// Function Return Example
func myFunction(x int, y int) int {
	return x + y
}

// Named Return Values
func myFunction2(x int, y int) (res int) {
	res = x + y
	return res //the return statement specifies the variable name
}

// you can use below to type same parameter types
func myFunction2_2(x, y int) (res int) {
	res = x + y
	return res //the return statement specifies the variable name
}

// Multiple Return Values
func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return // result, txt1 is omitted
}

// recursion
func testcount(x int) int {
	if x == 11 { // exit condition
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1) // call self
}

func mainFunc() {
	fmt.Println(myFunction(1, 2))
	fmt.Println(myFunction2(1, 2))
	a, b := myFunction3(5, "Hello")
	fmt.Println(a, b)
	_, b2 := myFunction3(3, "World") // _ 丢弃一个值
	fmt.Println(b2)

	testcount(4)
}
