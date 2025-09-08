package main

import (
	"fmt"
	"regexp"
)

func mainReg() {
	r := regexp.MustCompile(`\d+`)

	fmt.Println(r.MatchString("abc123"))      // true
	fmt.Println(r.FindString("abc123def456")) // "123"

	// 返回所有匹配的子串。n 表示最多返回几个，-1 表示全部
	fmt.Println(r.FindAllString("abc123def456", -1)) // ["123" "456"]

	// 返回第一个匹配的子串及其子组（括号内的捕获组）
	r = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	matches := r.FindStringSubmatch("今天是2024-04-05")
	fmt.Println(matches) // ["2024-04-05" "2024" "04" "05"]

	// 替换所有匹配项
	r = regexp.MustCompile(`\d+`)
	result := r.ReplaceAllString("abc123def456", "XXX")
	fmt.Println(result) // "abcXXXdefXXX"

	// 按正则分割字符串
	r = regexp.MustCompile(`\s+`)
	parts := r.Split("a b   c\t\nd", -1)
	fmt.Println(parts) // ["a" "b" "c" "d"]

}
