package main

import "fmt"

func mainSelect1() {
	ch := make(chan int, 1) // 有缓冲channel，缓冲区大小为1元素

	// 非阻塞发送
	select {
	case ch <- 42:
		fmt.Println("Data sent1")
	default:
		fmt.Println("Send failed1")
	}

	// 再次尝试非阻塞发送
	select {
	case ch <- 43:
		fmt.Println("Data sent2")
	default:
		fmt.Println("Send failed2")
	}
}

// Data sent1
// Send failed2

// ● 第一次发送操作ch <- 42成功，因为缓冲区为空。
// ● 第二次发送操作ch <- 43失败，因为缓冲区已满，执行default分支。

func mainSelect2() {
	ch := make(chan int, 1) // 有缓冲channel，缓冲区大小为1

	// 非阻塞接收，ch无数据
	select {
	case val := <-ch: // short variable declaration
		fmt.Println("Received:", val)
	default:
		fmt.Println("Receive failed")
	}

	// 发送数据
	ch <- 42

	// 再次尝试非阻塞接收,ch有数据
	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	default:
		fmt.Println("Receive failed")
	}
}

// Receive failed
// Received: 42
