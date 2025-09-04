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

func mainNew() {
	// 方法1: 使用 new 创建 (返回指针)
	c1 := new(Counter)
	c1.Inc()
	fmt.Println("c1 value:", c1.Value()) // 输出: c1 value: 1

	// 方法2: 直接声明变量 (零值可用)
	var c2 Counter
	c2.Inc()
	c2.Inc()
	fmt.Println("c2 value:", c2.Value()) // 输出: c2 value: 2

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Println(a) //[no error Eio invalid argument]
	fmt.Println(s) // [no error Eio invalid argument]
	fmt.Println(m) //map[0:no error 1:Eio 2:invalid argument]
}
