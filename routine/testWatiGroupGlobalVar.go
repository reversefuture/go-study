package main

import (
	"fmt"
	"sync"
	"time"
)

func mainWaitGroup2() {
	// 创建一个通道（channel），用于在不同 goroutine 之间传递数据
	ch := make(chan int) // 无缓冲，需要立即收发

	// 创建一个 全局WaitGroup组（不用&wg取地址操作和wg *sync.WaitGroup指针了)，用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 启动 3 个 goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // 为每个 goroutine 增加 WaitGroup 的计数器
		go func(id int) { // go: 启动一个新的 goroutine，并发执行一个函数
			defer wg.Done() //defer: 延迟执行函数，通常用于资源清理或计数器减少, 在 goroutine 完成时减少 WaitGroup 的计数器

			// 模拟一些工作
			time.Sleep(time.Second * time.Duration(id))
			fmt.Printf("Goroutine %d is done.\n", id)

			// 向通道发送数据
			ch <- id
		}(i)
	}

	// 启动一个额外的 goroutine 来关闭通道
	go func() {
		wg.Wait() // 等待wg组所有 goroutine 完成
		close(ch) //  关闭通道，表示不再有数据发送。 接收方（即从通道读取数据的goroutine）可以感知到通道已关闭，从而避免无谓的等待。
	}()

	// 使用 select 从通道中读取数据，直到通道关闭。非阻塞式接收必须配合time.Sleep，否则会失败
	for {
		select {
		case result, ok := <-ch: // ok 是一个布尔值，表示通道是否已关闭。如果通道关闭，ok 为 false。
			if !ok {
				// 通道已关闭，退出循环
				fmt.Println("All goroutines are finished.")
				return
			}
			fmt.Printf("Received result from Goroutine %d.\n", result)
		default:
			// 如果没有数据可接收，可以执行其他操作
			// 这里我们简单地等待一段时间
			time.Sleep(100 * time.Millisecond)
		}
	}

	// for result := range ch { //从通道中读取数据，直到通道关闭，阻塞式接收，成功！！！
	// 	fmt.Printf("Received result from Goroutine %d.\n", result)
	// }

	fmt.Println("All goroutines are finished.")
}

// Goroutine 1 is done.
// Received result from Goroutine 1.
// Goroutine 2 is done.
// Received result from Goroutine 2.
// Goroutine 3 is done.
// Received result from Goroutine 3.
// All goroutines are finished.
