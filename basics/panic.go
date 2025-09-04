package main

import (
	"fmt"
)

func mainPanic1() {
	fmt.Println("Start")
	panic("Something went wrong!")
	fmt.Println("This will not be printed")
}

// Start
// panic: Something went wrong!

// goroutine 1 [running]:
// main.main()
//         /path/to/main.go:7 +0x4d

func mainPanic2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start")
	panic("Oops! A problem occurred")
	fmt.Println("This won't run")
}

// Start
// Recovered from panic: Oops! A problem occurred

func mainPanic3() {
	safeDivide(10, 0)
}
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Caught panic:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	fmt.Println("Result:", a/b)
}

// aught panic: division by zero
