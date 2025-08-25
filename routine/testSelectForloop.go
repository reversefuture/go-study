package main

import (
	"fmt"
	"time"
)

func mainSelectForloop() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "from ch2"
	}()
	// 确保多次执行 select 语句，从而处理所有channel的数据
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

// from ch1
// from ch2
