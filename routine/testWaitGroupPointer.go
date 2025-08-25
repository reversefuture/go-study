package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) { // *sync.WaitGroup 类型指针
	defer wg.Done() // 确保在函数结束时调用Done减少wg的计数器
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func mainWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // 每启动一个goroutine，计数器加1
		go worker(i, &wg) //&wg取地址操作
	}

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("All workers done")
}

// Worker 3 starting
// Worker 2 starting
// Worker 1 starting
// Worker 1 done
// Worker 2 done
// Worker 3 done
// All workers done
