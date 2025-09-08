package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data    = make(map[string]int)
	rwMutex sync.RWMutex
	wg      sync.WaitGroup
)

// 读操作
func reader(id int) {
	defer wg.Done() // 减少wg的任务计数器，Add(-1)
	for i := 0; i < 3; i++ {
		rwMutex.RLock() // 获取读锁
		value := data["counter"]
		rwMutex.RUnlock() // 释放读锁
		fmt.Printf("Reader %d 读取值: %d\n", id, value)
		time.Sleep(1000 * time.Millisecond) // 等待时间更长，让出cpu时间给writer，阻塞wg.Done
	}
}

// 写操作
func writer(id int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		rwMutex.Lock() // 获取写锁（独占）
		data["counter"]++
		fmt.Printf("Writer %d 将值增加为: %d\n", id, data["counter"])
		rwMutex.Unlock() // 释放写锁
		time.Sleep(150 * time.Millisecond)
	}
}

func mainRW() {
	data["counter"] = 0

	// 启动 3 个读 goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i)
	}

	// 启动 2 个写 goroutine
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go writer(i)
	}

	wg.Wait()
	fmt.Println("最终值为:", data["counter"])
}

// Writer 2 将值增加为: 1
// Reader 2 读取值: 1
// Reader 3 读取值: 1
// Reader 1 读取值: 1
// Writer 1 将值增加为: 2
// Writer 1 将值增加为: 3
// Writer 2 将值增加为: 4
// Writer 2 将值增加为: 5
// Writer 1 将值增加为: 6
// Reader 1 读取值: 6
// Reader 2 读取值: 6
// Reader 3 读取值: 6
// Reader 2 读取值: 6
// Reader 1 读取值: 6
// Reader 3 读取值: 6
// 最终值为: 6
