package main

import (
	"context"
	"fmt"
	"time"
)

func worker2(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Context canceled: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(500 * time.Millisecond) // 当worker执行time.Sleep时，它会进入阻塞状态，此时CPU会被分配给其他可运行的goroutine。
		}
	}
}

func mainWithTimeout() {
	// 创建一个带有超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 确保在函数结束时取消context

	// 启动多个worker goroutine
	for i := 1; i <= 3; i++ {
		go worker2(ctx, i)
	}

	// 等待一段时间，以便观察worker的执行
	time.Sleep(3 * time.Second)
}

// Worker 3: Working...
// Worker 1: Working...
// Worker 2: Working...
// Worker 2: Working...
// Worker 3: Working...
// Worker 1: Working...
// Worker 1: Working...
// Worker 2: Working...
// Worker 3: Working...
// Worker 2: Working...
// Worker 3: Working...
// Worker 1: Working...
// Worker 2: Context canceled: context deadline exceeded
// Worker 3: Context canceled: context deadline exceeded
// Worker 1: Context canceled: context deadline exceeded
