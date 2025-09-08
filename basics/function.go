package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

/*
传递参数时，将参数复制一份传递到函数中，对参数进行调整后，不影响参数值。

Go 语言默认是值传递。
1. 引用传递
传递参数时，将参数的地址传递到函数中，对参数进行调整后，影响参数值。
func toJson (res *Result){}
toJson(&res)

*/

// Function Return Example
func myFunction(x int, y int) int {
	return x + y
}

// Named Return Values
func myFunction2(x int, y int) (res int) {
	res = x + y
	return res //the return statement specifies the variable name
}

// you can use below to type same parameter types
func myFunction2_2(x, y int) (res int) {
	res = x + y
	return res //the return statement specifies the variable name
}

// Multiple Return Values
func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return // result, txt1 is omitted
}

// recursion 递归
func testrecursion(x int) int {
	if x == 6 { // exit condition
		return 0
	}
	fmt.Println(x)
	return testrecursion(x + 1) // call self
}

func mainFunc0() {
	fmt.Println(myFunction(1, 2))
	fmt.Println(myFunction2(1, 2))
	a, b := myFunction3(5, "Hello")
	fmt.Println(a, b)
	_, b2 := myFunction3(3, "World") // _ 丢弃一个值
	fmt.Println(b2)

	testrecursion(4)

	fmt.Println(MD5("abc")) // 900150983cd24fb0d6963f7d28e17f72
}

// 函数类型，可以用在任何地方 -- 作为字段类型，参数或者返回值
type Add func(a int, b int) int

func mainFunc2() {
	fmt.Println(process3(func(a int, b int) int {
		return a + b
	}))
}

func process3(adder Add) int {
	return adder(1, 2)
}

func mainfunc2() {
	// 用nextInt从byte读取所有数字
	// b := []byte("a11")
	// for i := 0; i < len(b); { // >> 循环输出 3, 11
	// 	x, i := nextInt(b, i)
	// 	fmt.Println(i, x)
	// }

	// 示例：从字符串读取数据并用到缓冲区
	reader := strings.NewReader("hello")

	buf := make([]byte, 5) // 准备长度为 5 的缓冲区
	n, err := ReadFull(reader, buf)

	fmt.Printf("读取字节数: %d\n", n)
	fmt.Printf("缓冲区内容: %q\n", string(buf))
	fmt.Printf("错误: %v\n", err)
}

/*
单引号 '0' 表示一个字符常量，其底层是一个字节，对应 ASCII 码值 48。
所以 '0' 实际上等于 48。
*/

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0' // 转换成对应的整数值
	}
	return x, i
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

type Reader interface {
	Read(buf []byte) (nr int, err error)
}

// “读取全部数据直到缓冲区填满或发生错误
func ReadFull(r *strings.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 生成签名
func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
		}
	}
	// 自定义密钥
	var secret = "123456789"

	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(secret))
	return sign
}
