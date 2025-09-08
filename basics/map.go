package main

import (
	"encoding/json"
	"fmt"
)

/*
Map 集合是无序的 key-value 数据结构。

Map 集合中的 key / value 可以是任意类型，但所有的 key 必须属于同一数据类型，所有的 value 必须属于同一数据类型，key 和 value 的数据类型可以不相同。

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.
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

// 生成 JSON
func mainMapJson() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{ //嵌套map
		"username": "Tom",
		"age":      "30",
		"hobby":    []string{"读书", "爬山"},
	}
	fmt.Println("map data :", res)

	//序列化
	jsons, errs := json.Marshal(res) //return byte[]
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)
}

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

	b := map[string]int{}
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	val1, ok1 := a["brand"] // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
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
