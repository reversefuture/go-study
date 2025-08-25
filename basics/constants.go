package main

import (
	"fmt"
)

const PI = 3.14 // unchangeable and read-only.
const (         // Multiple Constants Declaration
	A = 1 // untyped
	B = 3.14
	C = "Hi!"
)

func mainConstants() {
	fmt.Println(PI)
	fmt.Println(A)
}
