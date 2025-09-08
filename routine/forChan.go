package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	count := 0
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		count++
		c <- count // 向c发送数据并等待下次cpu时间
		time.Sleep(time.Nanosecond * 1)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for { // 循环一值读取c中数据
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		if data == 10 {
			close(c) // 会触发 panic: send on closed channel终止
		}
	}
}

// worker 1 got 1
// worker 0 got 2
// worker 2 got 3
// worker 3 got 4
// worker 4 got 5
// worker 1 got 6
// worker 0 got 7
// worker 2 got 8
// worker 3 got 9
// worker 4 got 10
// worker 4 got 0
// worker 4 got 0
// ........
// panic: send on closed channel
