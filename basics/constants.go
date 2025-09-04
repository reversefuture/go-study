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

func mainConstant() {
	fmt.Println(PI)
	fmt.Println(A)

	var testSize ByteSize
	testSize = 10000
	var testSize2 ByteSize = 1000
	testSize3 := ByteSize(100) // 短类型声明+强制类型转换
	fmt.Println(testSize)
	fmt.Println(testSize2)
	fmt.Println(testSize3)
}

// 类型别名，底层 数值类型，可以和float64强制转换：size = ByteSize(f)
type ByteSize float64

// 枚举常量
// iota常量计数器，从 0 开始，只在const中使用
const (
	_           = iota             // 通过赋予空白标识符来忽略第一个值,后面的值从1开始
	KB ByteSize = 1 << (10 * iota) // iota=1， 1<<10 = 2^10 = 1024
	// KB2 =1e3  // 1000 (SI)国际单位制
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// 注意：
// KB, MB 等是 ByteSize 类型，不是 int
// var x int = 1024 和 var y ByteSize = KB 类型不同
// 不能直接比较或运算，需转换

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
