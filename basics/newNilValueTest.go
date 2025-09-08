package main

import (
	"fmt"
	"sync"
)

// Counter 是一个线程安全的计数器
type Counter struct {
	mu sync.Mutex // 零值即为解锁状态，无需初始化
	n  int        // 计数器，零值为 0
}

// Inc 增加计数器的值。它使用互斥锁来保证线程安全。
func (c *Counter) Inc() {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 函数返回时自动释放锁
	c.n++               // 安全地增加计数
}

// Value 返回当前计数器的值。它也使用互斥锁来保证读取时的线程安全。
func (c *Counter) Value() int {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 函数返回时自动释放锁
	return c.n          // 返回当前值
}

const (
	Enone  = iota
	Eio    = iota
	Einval = iota
)

func main() {
	slice1 := make([]int, 10, 100)
	fmt.Println(slice1) //[0 0 0 0 0 0 0 0 0 0]

	slice2 := new([]int)
	fmt.Println(slice2) // &[]  指向一个空（nil）slice 的指针”，而且这个 slice 是零值（nil slice），不能直接使用 append 等操作，除非先赋值
	// 当 * 出现在一个指针变量前时，表示“取该指针指向的值”。
	// slice2 = append(slice2, 1) //编译错误：cannot use slice2 (type *[]int) as []int
	*slice2 = append(*slice2, 1)

	slice3 := [...]int{}
	fmt.Println(slice3) // []

	// 方法1: 使用 new 创建 (返回指针)
	c1 := new(Counter)
	c1.Inc()
	fmt.Println("c1 value:", c1.Value()) // 输出: c1 value: 1

	// 方法2: 直接声明变量 (零值可用)
	var c2 Counter
	c2.Inc()
	c2.Inc()
	fmt.Println("c2 value:", c2.Value()) // 输出: c2 value: 2

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"} // 数组，...编译时确定固定长度,
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}    // slice, 运行时动态确定长度，长度 = 最大索引 + 1
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	a2 := [...]string{
		0: "zero",
		2: "two",
	}
	// 等价于 [3]string{"zero", "", "two"}
	fmt.Println(a[0]) //no error
	fmt.Println(s)    // [no error Eio invalid argument]
	fmt.Println(m)    //map[0:no error 1:Eio 2:invalid argument]
	fmt.Println(a2)   // [zero  two]
}
