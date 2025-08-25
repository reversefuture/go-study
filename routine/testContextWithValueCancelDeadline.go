package main

import (
	"context"
	"fmt"
	"time"
)

/**
这个例子中，3个 context 是彼此独立的，它们分别用于不同的 goroutine 中，各自的功能和生命周期互不影响。如果需要让多个 context 之间有关联，可以通过派生（例如 context.WithCancel(ctxParent)）来实现
*/

func worker3(ctx context.Context, workerId int) {
	// 从context中获取值
	if value := ctx.Value("key"); value != nil {
		fmt.Printf("Worker %d: Received value: %v\n", workerId, value)
	}

	for { //为了让 goroutine 能持续检查 ctx.Done()，同时又能执行周期性任务
		select {
		case <-ctx.Done():
			// 当context被取消或超时时，打印错误信息并退出
			fmt.Printf("Worker %d: Context canceled: %v\n", workerId, ctx.Err())
			return
		default:
			// 模拟工作，只要context.Done不取消或超时就一直执行!!!
			fmt.Printf("Worker %d: Working...\n", workerId)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker3_2(ctx context.Context, workerId int) {
	if value := ctx.Value("key"); value != nil {
		fmt.Printf("Worker %d: Received value: %v\n", workerId, value)
	}

	ticker := time.NewTicker(500 * time.Millisecond) // return  a channel that will send the current time on the channel after each tick.
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Context canceled: %v\n", workerId, ctx.Err())
			return
		case now := <-ticker.C: //用 定时器ticker 控制节奏
			fmt.Printf("Worker %d: Working..., at %s \n", workerId, now.Format("2006-01-02 15:04:05"))
			// 模拟工作，但不要 Sleep，否则会阻塞 select
		}
	}
}

func mainContexts() {
	// 创建一个带有取消功能的context
	ctxCancel, cancel := context.WithCancel(context.Background())

	// 启动一个goroutine，3秒后取消context。主 goroutine 不会被 time.Sleep 阻塞，程序可以继续执行其他任务
	go func() {
		time.Sleep(3 * time.Second)
		cancel() // go中cancel() 的调用是异步的，会在 3 秒后自动执行。
	}()

	// 启动worker1 goroutine
	go worker3_2(ctxCancel, 1)

	// 等待一段时间，以便观察worker的执行
	time.Sleep(1 * time.Second)

	// 创建一个带有截止时间的context
	deadline := time.Now().Add(2 * time.Second)
	ctxDeadline, _ := context.WithDeadline(context.Background(), deadline)

	// 启动另一个worker2 goroutine， worker1还在继续
	go worker3_2(ctxDeadline, 2)

	// 等待一段时间，以便观察worker的执行
	time.Sleep(3 * time.Second)

	// 创建一个带有键值对的context
	ctxValue := context.WithValue(context.Background(), "key", "example value")

	// 启动第三个worker goroutine
	go worker3_2(ctxValue, 3)

	// 等待一段时间，以便观察worker的执行
	time.Sleep(1 * time.Second)
}

// Worker 1: Working..., at 2025-08-25 18:58:50
// Worker 1: Working..., at 2025-08-25 18:58:51
// Worker 1: Working..., at 2025-08-25 18:58:51
// Worker 2: Working..., at 2025-08-25 18:58:51
// Worker 2: Working..., at 2025-08-25 18:58:52
// Worker 1: Working..., at 2025-08-25 18:58:52
// Worker 1: Working..., at 2025-08-25 18:58:52
// Worker 2: Working..., at 2025-08-25 18:58:52
// Worker 2: Context canceled: context deadline exceeded
// Worker 1: Context canceled: context canceled
// Worker 3: Received value: example value
// Worker 3: Working..., at 2025-08-25 18:58:54
