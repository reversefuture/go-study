package main

import (
	"fmt"
)

/*
Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.
*/

/*
Allowed Key Types
The map key can be of any data type for which the equality operator (==) is defined. These include:
Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)

Invalid key types are:
Slices
Maps
Functions
These types are invalid because the equality operator (==) is not defined for them

The map values can be any type.
*/

func mainMap1() {
	// var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	// b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	// fmt.Printf("a\t%v\n", a)
	// fmt.Printf("b\t%v\n", b)

	var a = make(map[string]string) // The map is empty now
	fmt.Println(len(a))
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	fmt.Println(len(a))

	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["model"] // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	delete(a, "year") // remove element
}

// iterate over the elements in a map
func mainMap2() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}

// Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
func mainMap3() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order in slice b
		fmt.Printf("%v : %v, ", element, a[element])
	}
}

/*
当你需要将映射作为结构体字段的时候
*/
type Saiyan3 struct {
	Name    string
	Friends map[string]*Saiyan3
}

func mainMap() {
	goku := &Saiyan3{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan3, 10),
	}
	goku.Friends["krillin"] = &Saiyan3{
		Name:    "krillin",
		Friends: make(map[string]*Saiyan3),
	}

	// 像 make，这种特定用于映射和数组。我们可以定义为复合方式：
	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
	fmt.Println(lookup)
	// 使用 for 组合 range 关键字迭代映射：
	for key, value := range lookup {
		fmt.Println(key, value)
	}
}
