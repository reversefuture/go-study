package main

import (
	"fmt"
	"time"
)

func mainChanBufferBlocked() {
	// ch := make(chan int, 2) // 有缓冲buffer channel，缓冲区大小为2个元素，2s后 <-ch 接收数据的时候不会阻塞
	ch := make(chan int) // 无缓冲buffer channel， <-ch 接收数据，会阻塞直到有人发送

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Sending %d...\n", i)
			ch <- i // 发送数据
			fmt.Printf("%d sent\n", i)
		}
	}()

	time.Sleep(time.Second * 2) // 模拟延迟
	for i := 1; i <= 3; i++ {
		fmt.Printf("Receiving %d...\n", i)
		fmt.Println(<-ch) // 阻塞式接收数据
		fmt.Printf("%d received\n", i)
	}
}

// Sending 1...
// 1 sent
// Sending 2...
// 2 sent
// Sending 3...
// Receiving 1...
// 1
// 1 received
// Receiving 2...
// 2
// 2 received
// Receiving 3...
// 3 sent
// 3
// 3 received
