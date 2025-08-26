package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
Slice：是一个引用类型，它包含三个部分：
指向底层数组的指针（pointer）
长度（len）
容量（cap）
所以切片本身只是一个“描述符”，指向底层数组的一段。

len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

调整slice大小：
1. slice again [start:end]
2. append 是相当特别的。如果底层数组满了，它将创建一个更大的数组并且复制所有原切片中的值。Go 使用 2x 算法来增加数组长度
scores = append(scores, 5) // 重新赋值以防创建底层新数组

这有四种方式初始化一个切片：
names := []string{"leto", "jessica", "paul"}
checks := make([]bool, 10) // 需要用到切片具体索引时
var names []string // 指向空的切片，用于当元素数量未知时与 append 连接
scores := make([]int, 0, 20) // 知道最大长度

[X:] 是 从 X 到结尾 的简写，然而 [:X] 是 从开始到 X 的简写。
不像其他的语言，Go 不支持负数索引。如果我们想要切片中除了最后一个元素的所有值，可以这样写：
scores := []int{1, 2, 3, 4, 5}
scores = scores[:len(scores)-1]

交换元素：
  source[index], source[lastIndex] = source[lastIndex], source[index]

删除元素：removeAtIndex

拷贝：func copy(dst, src []Type) int
copy(worst[2:4], scores[:5]) //[0 0 4 15 0]， 只复制到指定位置
*/

func mainSlice1() {
	myslice1 := []int{}        // empty slice of 0 length and 0 capacity
	fmt.Println(len(myslice1)) //0
	fmt.Println(cap(myslice1)) //0
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2)) //4
	fmt.Println(cap(myslice2)) //4
	fmt.Println(myslice2)

}

func mainSlice2() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4] //create a slice from an array

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))   //2
	fmt.Printf("capacity = %d\n", cap(myslice)) //4
}

func mainSlice3() {
	myslice1 := make([]int, 5, 10) // create slice with make([]type, length, capacity)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))   //5
	fmt.Printf("capacity = %d\n", cap(myslice1)) //10

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))   //5
	fmt.Printf("capacity = %d\n", cap(myslice2)) //5
}

func mainSlice4() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)

	myslice1 = append(myslice1, 20, 21) // append 2 elements to slice, capacity will grow 2 * appended length
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))   //8
	fmt.Printf("capacity = %d\n", cap(myslice1)) //12

	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...) // append one slice to another, capacity is mySlice1
	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))   //11
	fmt.Printf("capacity=%d\n", cap(myslice3)) //12
}

func mainSlice5() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))   //4
	fmt.Printf("capacity = %d\n", cap(myslice1)) //5

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))   //2
	fmt.Printf("capacity = %d\n", cap(myslice1)) //5

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))   //6
	fmt.Printf("capacity = %d\n", cap(myslice1)) //10
}

//  When using slices, Go loads all the underlying elements into the memory.
// If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
// The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
// copy(dest, src)

func mainSlice6() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))   //15
	fmt.Printf("capacity = %d\n", cap(numbers)) //15

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))   //5
	fmt.Printf("capacity = %d\n", cap(numbersCopy)) //5
}

func mainSlice7() {
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4] // 切片不改变原来的scores
	slice[0] = 999
	fmt.Println(slice)  // [999 4]
	fmt.Println(scores) //[1 2 999 4 5]

	scores = removeAtIndex(scores, 2)
	fmt.Println(scores) // [1 2 5 4]
}

func mainCopy() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	// copy(worst, scores[:5]) //[4 5 17 31 35]
	// copy(worst[2:4], scores[:5]) //[0 0 4 15 0]， 只复制到指定位置
	// copy(worst[2:4], scores[:7]) //[0 0 10 15 0]
	copy(worst, scores[:7]) //[13 24 37 52 65]
	fmt.Println(worst)
}

// 不会保持顺序
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// 交换最后一个值和想去移除的值
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
