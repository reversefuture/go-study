package main

import (
	"fmt"
	"time"
)

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}
func comsumer(ch chan string) {
	fmt.Println("consumer start")
	for { // 阻塞式循环等待读取
		val := <-ch
		fmt.Println(val)
	}
	fmt.Println("consumer end")
}

func main() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	go comsumer(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

// main start
// consumer start
// producer start
// producer end
// a
// b
// c
// d
// main end
