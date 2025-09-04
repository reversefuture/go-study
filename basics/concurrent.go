package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 唯一的共享状态是通道
func mainConcurrent() {
	c := make(chan int, 100) // 带100缓冲通道
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		// go worker.process(c) // 启动一个新的阻塞协程
		// go worker.selectProcess(c) // 启动一个新的非阻塞协程
		go worker.selectProcessTimeout(c)
	}

	// 主线程：无限循环生产数据，受通道容量限制，可能阻塞。
	for {
		c <- rand.Int()
		fmt.Println("C length: ", len(c))
		time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c // 循环一直阻塞读取c中数据
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

// worker 1 got 2669790136846112673
// worker 0 got 258348563189919694
// worker 2 got 6647039193439426123

// 每个 worker 每隔 50ms 尝试向 c 发送一个随机数，如果 c 已满（无法立即发送），就走 default 分支并打印 "dropped"
func (w *Worker) selectProcess(c chan int) {
	for {
		select {
		case c <- rand.Int(): // 如果通道已满或即使未满但调度不及时，case 分支不能立即执行，就会走 default。
		default: //不等待，打印 "dropped"。
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50) //休息：同步、阻塞当前 goroutine 的操作。 在这期间，该 goroutine 进入等待状态，不会被调度执行，相当于yield。
	}
}

// ......
// C length: 90
// C length: 91
// C length: 97
// dropped
// dropped
//......

func (w *Worker) selectProcessTimeout(c chan int) {
	// 尝试发送数据到 c，等待最多 100ms，如果在这期间第一个 case 还不能执行，就走超时分支。
	for {
		select {
		case c <- rand.Int():
			// case <-time.After(time.Millisecond * 100): //返回一个 chan time.Time，100ms 后会向这个通道写入当前时间。
			// 	fmt.Println("timed out")
		case t := <-time.After(time.Millisecond * 100): // 接收After返回的值
			fmt.Println("timed out at", t)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

// ...
// C length: 90
// C length: 96
// timed out at 2025-08-27 23:24:19.7552313 +0800 CST m=+0.910109201
// timed out at 2025-08-27 23:24:19.8060678 +0800 CST m=+0.960945701
// timed out at 2025-08-27 23:24:19.8060678 +0800 CST m=+0.960945701
// ...
