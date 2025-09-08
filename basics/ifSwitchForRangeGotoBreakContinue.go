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

// 默认每个 case 带有 break
// case 中可以有多个选项
// fallthrough 不跳出，并执行下一个 case
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

// continue跳过本次循环，只能用于 for。
func mainFor1() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		fmt.Println(i)
	}
}

// break跳出当前循环，可⽤于 for、switch、select。
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
	// 多个参数初始化
	for i, j, k := 1, 2, 3; i < 2; i++ {
		fmt.Println(i, j, k) // 1 2 3
	}
}

// uses range to iterate over an array and print both the indexes and the values at each (
// Range can be used on array,slice,string, string(v is the rune code)，Channel(till chan close),
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

	s := "你好, world"
	for i, r := range s {
		fmt.Printf("字节索引=%d, 字符='%c', unicode码点=%d\n", i, r, r)
	}

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for v := range ch {
		fmt.Println(v) //
	}
}

// 改变函数内代码执行顺序，不能跨函数使用。
func mainGoto() {
	fmt.Println("begin")

	for i := 1; i <= 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i =", i)
	}

END:
	fmt.Println("end")
}

// begin
// i = 1
// i = 2
// i = 3
// i = 4
// i = 5
// end
