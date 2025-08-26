package main

import (
	"fmt"
)

func mainArr() {
	var arr1 = [3]int{1, 2, 3}    // here length is defined
	var arr12 = [...]int{1, 2, 3} // here length is inferred
	arr2 := [5]int{4, 5, 6, 7, 8}
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(arr12)
	fmt.Println(arr2)

	cars[2] = "MI"
	fmt.Println(cars[0])
	fmt.Println(cars[2])
}

func mainArr2() {
	arr1 := [5]int{}              //not initialized, all will be 0
	arr2 := [5]int{1, 2}          //partially initialized, remaining will be 0
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arr4 := [5]int{1: 10, 2: 40}  // initializes only the second and third elements

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Println(len(arr2))

	for index, value := range arr1 { // range 遍历
		fmt.Printf("%d: %d", index, value)
	}
}
