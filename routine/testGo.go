package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func mainGo() {
	// go + func 启动一个goroutine
	go printNumbers()
	time.Sleep(time.Second * 3) // 等待goroutine执行完毕
	fmt.Println("Main function")
}
