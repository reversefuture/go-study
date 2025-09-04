package main

import (
	"fmt"
)

func mainIf() {
	x := 4
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is little than 5")
	}
}

func mainSitwch1() {
	day := 8

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}
}

// Multi-case switch
func mainSwitch2() {
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

func mainFor1() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		fmt.Println(i)
	}
}

func mainFor2() {
	for i := 0; i <= 100; i += 10 {
		if i == 30 {
			break
		}
		fmt.Println(i)
	}
	// 带初始化的 if
	if _, err := process2(); err != nil {
		fmt.Print(err)
	}
}

// uses range to iterate over an array and print both the indexes and the values at each (
func mainRange() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruits { // omit index
		fmt.Printf("%v\n", val)
	}

	for idx, _ := range fruits { // omit value
		fmt.Printf("%v\n", idx)
	}

	for i := 0; i <= len(fruits); i += 1 {
		fmt.Println(fruits[i])
	}
}
