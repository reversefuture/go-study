package main

//用 go run .执行才能自动包含study.go
import (
	"fmt"
)

func mainInterImplement() {
	name := "Tom"
	s, err := NewStudy(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}

// Tom 听 english
// Tom 说 english
// Tom 读 english
// Tom 写 english
