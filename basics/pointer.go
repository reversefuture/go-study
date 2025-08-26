package main

import (
	"fmt"
)

/*
之前讨论了内置类型，比如整数和字符串。既然现在我们要讨论结构，那么我们需要把讨论范围扩展到指针。
许多时候，我们并不想让一个变量直接关联到值，而是让它的值为一个指针，通过指针关联到值。一个指针就是内存中的一个地址；指针的值就是实际值的地址。这是间接地获取值的方式。形象地来说，指针和实际值的关系就相当于房子和指向该房子的方向之间的关系。
为什么我们想要一个指针指向值而不是直接包含该值呢？这归结为 Go 中传递参数到函数的方式：镜像复制
这样更省内存

- 应该是值还是指向值的指针呢？ 这儿有两个好消息，首先，无论我们讨论下面哪一项，答案都是一样的：
局部变量赋值
结构体指针
函数返回值
函数参数
方法接收器
- 需要保留原来值不需要。如果你不确定，那就用指针咯。

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
