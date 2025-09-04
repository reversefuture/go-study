package main

import (
	"fmt"
	"unicode/utf8"
)

/**
bool keyword and can only take the values true or false

Signed Integers：
Type	Size	Range
int	Depends on platform:
32 bits in 32 bit systems and, Default
64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
-9223372036854775808 to 9223372036854775807 in 64 bit systems
int8	8 bits/1 byte	-128 to 127
int16	16 bits/2 byte	-32768 to 32767
int32	32 bits/4 byte	-2147483648 to 2147483647
int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807

Unsigned Integers：
Type	Size	Range
uint	Depends on platform:
32 bits in 32 bit systems and
64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and
0 to 18446744073709551615 in 64 bit systems
uint8	8 bits/1 byte	0 to 255
uint16	16 bits/2 byte	0 to 65535
uint32	32 bits/4 byte	0 to 4294967295
uint64	64 bits/8 byte	0 to 18446744073709551615

Float:
Type	Size	Range
float32	32 bits	-3.4e+38 to 3.4e+38. Default
float64	64 bits	-1.7e+308 to +1.7e+308.

String values must be surrounded by double quotes
**/

func mainTypes() {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value:  \n%v", x, x)
	fmt.Printf("Type: %T, value:  \n%v", y, y)

	var x2 uint = 500
	var y2 uint = 4500
	fmt.Printf("Type: %T, value: %v \n", x2, x2)
	fmt.Printf("Type: %T, value: %v \n", y2, y2)

	var x3 float32 = 123.78
	var y3 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x3, x3)
	fmt.Printf("Type: %T, value: %v", y3, y3)

	// var x float32= 3.4e+39 // overflow
}

// `string` 是**值类型**，但它的底层数据是不可变的，所以即使共享，也不会出问题。有引用类型的性能特征（共享存储）。
func mainString() {
	s1 := "hello"
	s2 := s1
	s2 = "World"
	fmt.Println((s1))
	fmt.Println((s2))
}

func mainString2() {
	stra := "the spice must flow" // 字符串是不可变的
	byts := []byte(stra)          // 实际上创建了数据的副本
	strb := string(byts)          //实际上创建了数据的副本
	fmt.Println(strb)

	fmt.Println(int64(33))

	s := "椒"
	fmt.Println(len(s))                    //3
	fmt.Println(len([]byte(s)))            //3
	fmt.Println(len([]rune(s)))            // 输出: 1 ✅ 正确的“字符个数”
	fmt.Println(utf8.RuneCountInString(s)) // 输出: 1 ✅ 推荐（更高效，不分配内存）
}
