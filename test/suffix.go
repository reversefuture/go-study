package main

import (
	"fmt"
	"index/suffixarray"
)

func mainsuffixarray() {
	text := "banana"
	pattern := "ana"
	data := []byte{2, 1, 2, 3, 2, 1, 2, 4, 4, 2, 1}
	sep := []byte{2, 1}

	fmt.Println(suffixarray.New(data).Lookup(sep, -1)) // [9 0 4]

	// 创建后缀数组索引
	index := suffixarray.New([]byte(text))

	// 查找所有匹配
	offsets := index.Lookup([]byte(pattern), -1) // -1 表示返回所有匹配

	for _, offset := range offsets {
		fmt.Printf("匹配位置: %d, 子串: %q\n", offset, text[offset:offset+len(pattern)])
	}
}

// 匹配位置: 3, 子串: "ana"
// 匹配位置: 1, 子串: "ana"
