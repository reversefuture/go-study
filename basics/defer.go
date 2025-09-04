package main

import (
	"fmt"
	"os"
)

// 尽管 Go 有一个垃圾回收器，一些资源仍然需要我们显式地释放他们。

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

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a")) // 被推迟函数的实参在推迟执行时就会求值
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func mainDefer2() {
	// b()
	// entering: b
	// in b
	// entering: a
	// in a
	// leaving: a
	// leaving: b

	// 	for i := 0; i < 3; i++ {
	// 		defer fmt.Println(i) // 输出：2， 1， 0（不是 0 1 2），参数实时求出但是函数LIFO执行
	// 	}

	// x := 10
	// defer fmt.Println("deferred:", x) //会立即对参数 x 求值
	// x = 20
	// fmt.Println("immediate:", x)

	f()  // 1
	f2() // 2
}

// immediate: 20
// deferred: 10
func f() {
	i := 1
	defer fmt.Println(i) // 输出：1，不是 2
	i++
	return
}

func f2() {
	i := 1
	defer func() { //通过闭包引用变量, 闭包真实执行时i变了
		fmt.Println(i) // 输出：2
	}()
	i++
	return
}
