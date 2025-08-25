package main

import (
	"fmt"
	"time"
)

func mainDoneClose() {
	done := make(chan struct{}) // struct{} 是一个空结构体，它不占用任何内存空间，通常用于传递信号而不是实际的数据。

	go func() {
		for {
			select {
			case <-done: // 如果 done 通道被关闭，退出 goroutine
				fmt.Println("Goroutine exiting...")
				return
			default:
				fmt.Println("Goroutine working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(done) // 关闭通道，通知 goroutine 退出
	time.Sleep(1 * time.Second)
}

// Goroutine working...
// Goroutine working...
// Goroutine working...
// Goroutine working...
// Goroutine exiting...

func mainDoneBlock() {
	done := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second) // 模拟任务执行
		close(done)                 // 任务完成，关闭通道
	}()

	<-done // 阻塞，直到 done 通道被关闭
	fmt.Println("Task completed.")
}

// Task completed.
