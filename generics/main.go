package main

import "fmt"

// https://go.dev/doc/tutorial/generics

// Declare the Number interface type to use as a type constraint.
type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints), // Optional: Specify type arguments – the type names in square brackets
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

/*
func with type parameters:
make the function generic, enabling it to work with arguments of different types. You’ll call the function with type arguments and ordinary function arguments.
Each type parameter has a type constraint that acts as a kind of meta-type for the type parameter.
At compile time the type parameter stands for a single type – the type provided as a type argument by the calling code. If the type argument’s type isn’t allowed by the type parameter’s constraint, the code won’t compile.
Keep in mind that a type parameter must support all the operations the generic code is performing on it.
The comparable constraint allows any type whose values may be used as an operand of the comparison operators == and !=.
Using | specifies a union of the two types: int64 and float64.
*/
// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
