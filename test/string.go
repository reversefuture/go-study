package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func mainString() {
	s := "你好, world"
	fmt.Println(len(s)) // 13, bype length

	bs := []byte(s)
	fmt.Println("byte len: ", len(bs)) // 13

	rs := []rune(s)
	fmt.Println("rune len: ", len(rs)) // 9 unicode char length

	fmt.Println("utf8.RuneCountInString: ", utf8.RuneCountInString(s)) // 9

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ' ') // 输出每个byte: 228 32189 32160 32229 32165 32189 3244 3232 32119 32111 32114 32108 32100 32
	}
	fmt.Println("--------")
	for k, v := range s {
		fmt.Print(k, ':', v) //输出每个rune代码点： 0 58 203203 58 229096 58 447 58 328 58 1199 58 11110 58 11411 58 10812 58 100
	}
	fmt.Println("--------")
	// ✅ 正确：range string 自动解码为 rune
	for i, r := range s {
		fmt.Printf("位置 %d: %c (rune=%d)\n", i, r, r) //位置 0: 你 (rune=20320).....
	}

	// rune字面量
	var r rune = '中'
	fmt.Println(r)        // 输出 Unicode 码点：20013
	fmt.Printf("%U\n", r) // 输出：U+4E2D

	// 使用 Unicode 转义，自动
	fmt.Println('\u4E2D')     // '中'
	fmt.Println('\U0001F600') // '😀'

	//  判断 rune 类型（unicode 包）
	unicode.IsLetter(r)  // true
	unicode.IsDigit('3') // true
	unicode.IsSpace(' ') // true

	fmt.Println(utf8.DecodeRuneInString(s)) // 20320 3
	fmt.Println(utf8.ValidRune(r))          // true
}
