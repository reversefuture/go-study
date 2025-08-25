package main

import (
	"fmt"
	"log"
	"os"
)

func processFile(filename string) (err error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer func() {
		// 确保文件关闭
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing file: %v", cerr)
			if err == nil {
				err = cerr // 如果主错误为空，将关闭错误作为主错误返回
			}
		}
		log.Println("File closed") // 日志记录文件关闭
	}()

	// 读取文件内容
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// 处理文件内容
	log.Printf("Read %d bytes from file", n) // 日志记录读取的字节数
	return nil
}

func mainTestDefer() {
	filename := "example.txt"
	if err := processFile(filename); err != nil {
		log.Printf("Error: %v", err)
	}
}

// 2025/08/25 18:13:49 Read 4 bytes from file
// 2025/08/25 18:13:49 File closed
