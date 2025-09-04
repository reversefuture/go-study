package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	lock    sync.Mutex // 单个锁
)

func mainMutex1() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10) // 没有主程序会立马执行完退出
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

var (
	lock2 sync.Mutex
)

func mainMutex2() {
	go func() { lock2.Lock() }() // 未释放锁
	time.Sleep(time.Millisecond * 10)
	lock2.Lock() // 死锁
}

// fatal error: all goroutines are asleep - deadlock!
