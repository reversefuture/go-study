package main

import (
	"fmt"
)

func mainAppend() {
	arr := [...]float64{1.1, 3, 4.1} // 让编译器自动推断数组的长度
	fmt.Println(sum(&arr))           //8.2

	s := []byte{1, 2, 3} // len=3, cap=3
	d := []byte{4, 5}    // 要追加的数据
	s = Append(s, d)     // 调用函数
	fmt.Println(s)       //[1 2 3 4 5]

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x) //[1 2 3 4 5 6]

	y := []int{4, 5, 6}
	x = append(x, y...) // 功能同上
	fmt.Println(x)      //[1 2 3 4 5 6 4 5 6]

	// 自定义 append
	var b ByteSlice
	//fmt.Fprintf 在写入后通常会调用其目标的 String() 方法来格式化输出
	fmt.Fprintf(&b, "This hour has %d days\n", 7) // 将&b后面参数合并format后写入&b
	fmt.Print(string(b))                          // This hour has 7 days

}

func sum(arr *[3]float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	return
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // 重新分配
		// 为未来的增长,双重分配所需的内容.
		newSlice := make([]byte, (l+len(data))*2)
		// copy函数是预先声明的，适用于任何切片类型。
		copy(newSlice, slice) // 复制原始slice到扩容后的newSlice
		slice = newSlice      // 调整切片的长度到 (l+len(data))*2。
	}
	slice = slice[0 : l+len(data)] // 调整切片的长度到 l + len(data)。
	copy(slice[l:], data)          // 从slice[l:]原长度位置开始的“空位”开始拷贝data
	return slice
}

type ByteSlice []byte

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// 主体同上，只是没有返回值
	*p = slice
}

// 将一个指向 ByteSlice 的指针作为该方法的接收者， 这样该方法就能重写调用者提供的切片了。
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l+len(data) > cap(slice) { // 重新分配
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
	return len(data), nil // 满足标准的 io.Writer 接口
}

func (s *ByteSlice) String() string {
	// res := ""
	// for _, v := range s { // range s 实际上是在遍历指针本身（不合法），这会导致编译错误。
	// 	res += v
	// }
	return string([]byte(*s))
}
