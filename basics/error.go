package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
)

// mainError1 示例：基础错误处理
func mainError1() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program <number>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("not a valid number:", err)
	} else {
		fmt.Println("Parsed number:", n)
	}
}

func mainErrorAssert() {
	filename := "testError.txt"
	for try := 0; try < 2; try++ {
		_, err := os.Create(filename)
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			// deleteTempFiles() // 恢复一些空间。
			fmt.Println("Delete temp files")
			continue
		}
		return
	}
}

// 建你自己的错误类型。唯一的要求是你必须实现内建 error 接口的契约
type MyError interface {
	Error() string
}

type InvalidCountError struct {
	count string
}

func (e *InvalidCountError) Error() string {
	return "Invalid count: " + e.count + " (must be a positive integer)"
}

// process2 处理输入，返回整数或错误
func process2() (int64, MyError) {
	if len(os.Args) != 2 {
		return 0, errors.New("usage: program <number>")
	}

	count := os.Args[1]
	intCount, err := strconv.ParseInt(count, 10, 64) // 正确的 bitSize：64
	if err != nil {
		return 0, fmt.Errorf("cannot parse '%s' as integer: %w", count, err)
	}

	if intCount < 1 {
		return 0, &InvalidCountError{count: count}
	}

	return intCount, nil
}

func mainCustomError() {
	// mainError1()

	num, err := process2()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Valid count:", num)
}

// Go 标准库中有一个使用 error 变量的通用模式。例如， io 包中有一个 EOF 变量它是这样定义的：
// var EOF = errors.New("EOF")
// 这是一个包级别的变量（被定义在函数之外），可以被其他包访问（首字母大写）,如io.EOF
func mainEOF() {
	var input int
	_, err := fmt.Scan(&input)

	if err == io.EOF {
		fmt.Println("no more input!")
	}
}
