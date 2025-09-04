package main

import (
	"fmt"
	"log"
)

func mainPrint() {
	var x1 uint64 = 1<<64 - 1
	// %x用于字符串、字节数组以及整数，并生成一个很长的十六进制字符串
	fmt.Printf("%d %x; %d %x\n", x1, x1, int64(x1), int64(x1)) // 18446744073709551615 ffffffffffffffff; -1 -1
	// 整数
	fmt.Printf("十进制: %d\n", 255)
	fmt.Printf("十六进制: %x\n", 255)   // ff%
	fmt.Printf("大写十六进制: %X\n", 255) // FF
	fmt.Printf("二进制: %b\n", 255)    // 11111111
	fmt.Printf("八进制: %o\n", 255)    // 377

	// 浮点数
	fmt.Printf("浮点数: %f\n", 3.1415926)
	fmt.Printf("科学计数: %e\n", 123456.0)
	fmt.Printf("自动格式: %g\n", 0.000123)

	// 字符串
	fmt.Printf("字符串: %s\n", "Golang")
	fmt.Printf("带引号: %q\n", "Hello\nWorld")

	// 指针
	x := 10
	fmt.Printf("地址: %p\n", &x)

	// 类型和默认值
	fmt.Printf("类型: %T\n", x)
	fmt.Printf("值: %v\n", x)
	fmt.Printf("Go 语法: %#v\n", x)

	// 宽度和精度
	fmt.Printf("右对齐: |%10d|\n", 42)   // |        42|
	fmt.Printf("左对齐: |%-10d|\n", 42)  // |42        |
	fmt.Printf("补零: %08d\n", 42)      // 00000042
	fmt.Printf("小数: %.2f\n", 3.14159) // 3.14

	t := &T{7, -2.35, "abc\tdef"}
	t2 := T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)   // &{7 -2.35 abc   def}
	fmt.Printf("%+v\n", t)  //&{a:7 b:-2.35 c:abc     def}
	fmt.Printf("%#v\n", t)  // &main.T{a:7, b:-2.35, c:"abc\tdef"}
	fmt.Printf("%#v\n", t2) // main.T{a:7, b:-2.35, c:"abc\tdef"}

	fmt.Printf("%d%%\n", 3) //  %%	Prints the % sign

	// 自定义多个是多参数print
	Println(1, "is", true) //2025/08/29 12:43:16 1 is true
}

type T struct {
	a int
	b float64
	c string
}

// 若你想控制自定义类型的默认格式，只需为该类型定义一个具有 String() string 签名的方法
// func (t *T) String() string {
// 	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c) //7/-2.35/"abc\tdef"
// }

// 请勿通过调用 Sprintf 来构造 String 方法，因为它会无限递归你的 String 方法
// 要解决这个问题也很简单：将该实参转换为基本的字符串类型，它没有这个方法。

type MyString string

func (m MyString) String() string {
	// return fmt.Sprintf("MyString=%s", m) // 错误：会无限递归
	return fmt.Sprintf("MyString=%s", string(m)) // 可以：注意转换
}

// 将打印例程的实参直接传入另一个这样的例程。Printf 的签名为其最后的实参使用了 ...interface{} 类型，这样格式的后面就能出现任意数量，任意类型的形参了。
// 在 Printf 函数中，v 看起来更像是 []interface{} 类型的变量，但如果将它传递到另一个变参函数中，它就像是常规实参列表了。
// func Printf(format string, v ...interface{}) (n int, err error) {}

// Println 通过 fmt.Println 的方式将日志打印到标准记录器
func Println(v ...interface{}) { //...interface{}: 可变参数，可以传入任意数量、任意类型的参数。
	log.Output(2, fmt.Sprintln(v...)) // Output takes parameters (int, string)
}
