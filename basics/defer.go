package main

import (
	"fmt"
	"os"
	"time"
)

// defer释放资源
func mainDefer() {
	file, err := os.Open("go.sum")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 无论什么情况，在函数返回之后（本例中为 main() ），defer 将被执行
	defer file.Close()
	// 读取文件
}

// defer 调用的函数，参数的值在 defer 定义时就确定了。defer 函数内部所使用的变量的值需要在这个函数运行时才确定
func enter(s string) string {
	fmt.Println("entering:", s)
	return s
}

func leave(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer leave(enter("a")) // 被推迟函数的实参在推迟执行时就会求值,所以enter会先执行
	fmt.Println("in a")
}

func b() {
	defer leave(enter("b"))
	fmt.Println("in b")
	a()
}

// entering: b
// in b
// entering: a
// in a
// leaving: a
// leaving: b

func mainDefer0() {
	b()

	// 	for i := 0; i < 3; i++ {
	// 		defer fmt.Println(i) // 输出：2， 1， 0（不是 0 1 2），参数实时求出但是函数LIFO执行
	// 	}

	deferTest1() // 3
	deferTest2() // 4
	deferTest3() // 3
	// deferTest4() // 无输出
	deferTest5() // 无输出
	GoA()        //
}

// 结论：闭包获取变量相当于引用传递，而非值传递。
func deferTest1() {

	var a = 1
	var b = 2

	defer fmt.Println(a + b)
	a = 2
}

// main
// 3
func deferTest2() {
	var a = 1
	var b = 2

	defer func() {
		fmt.Println(a + b)
	}()
	a = 2
}

// main
// 4

// 结论：传参是值复制。
func deferTest3() {
	var a = 1
	var b = 2

	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)

	a = 2
}

// 3

// 结论：当os.Exit()方法退出程序时，defer不会被执行。
func deferTest4() {
	defer fmt.Println("1")
	os.Exit(0)
	fmt.Println("Test4")
}

// 输出：main

// 结论：defer 只对当前协程有效。
func deferTest5() {
	GoA()
	time.Sleep(1 * time.Second)
	fmt.Println("test 5")
}

func GoA() {
	defer func() { //先捕获error A，不会中断执行
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	go GoB() // 不sleep, 捕获 error A, error B没被捕获，B中断执行
	// time.Sleep(1 * time.Second) // 如果sleep，error A没被捕获，error B没被捕获，B中断执行
	panic("error A")
}

func GoB() {
	panic("error B")
}
