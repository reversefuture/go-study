package main

import (
	"fmt"
)

/*
为什么我们想要一个指针指向值而不是直接包含该值呢？这归结为 Go 中传递参数到函数的方式：镜像复制, 这样更省内存
* 可以用来:声明 + 解引用

var p *int // 我是一个指针（声明）
p = &x     // 我指向 x 的地址
y := *p    // 我要取 p 指向的值
var p *int = &x    // 声明指针 p，指向 x

new(T) 返回的是 *T —— 指向 T 类型零值的指针。
*/

type Saiyan struct {
	Name  string
	Power int
}

func mainPinter() {

	goku := Saiyan{"Power", 9000}
	Super(goku)             // Super 修改了原始值 goku 的复制版本而非本身
	fmt.Println(goku.Power) // 9000

	goku2 := &Saiyan{"Power", 9000} // 用了 & 操作符以获取值的地址（
	Super2(goku2)
	fmt.Println(goku2.Power) // 19000

	goku3 := &Saiyan{"Goku3", 9001}
	goku3.Super3()
	fmt.Println(goku3.Power) // 将会打印出 19001
}

func Super(s Saiyan) {
	s.Power += 10000
}

func Super2(s *Saiyan) { // *X 意思是 指向类型 X 值的指针 。
	s.Power += 10000
}

// 我们可以把一个方法关联在一个结构体上
func (s *Saiyan) Super3() { // *Saiyan 类型是 Super3 方法的接受者。
	s.Power += 10000
}
