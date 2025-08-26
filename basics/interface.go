package main

import (
	"fmt"
)

// Logger 接口
type Logger interface {
	Log(message string)
}

// Logger2 接口：支持可变参数打印
// ...interface{} 表示 “接收任意数量、任意类型的参数”。 interface{}：空接口，表示“任何类型”。...：可变参数（variadic parameter）
type Logger2 interface {
	Print(v ...interface{})
}

// 具体实现：StdLogger 同时实现两个接口
type StdLogger struct{}

func (s *StdLogger) Log(message string) {
	fmt.Println("LOG:", message)
}

func (s *StdLogger) Print(values ...interface{}) {
	for i, v := range values {
		fmt.Printf("第%d个: 值=%v, 类型=%T\n", i+1, v, v)
	}
}

// Server 使用组合：持有 Logger，匿名嵌入 Logger2
type Server struct {
	logger  Logger // 显式字段， 不依赖具体日志实现，而是依赖 Logger 接口
	Logger2        // 匿名嵌入，可以直接调用 Print
}

// 示例：函数接收接口，实现解耦
func process(logger Logger) {
	logger.Log("处理中...")
}

func main() {
	// 创建日志实例
	logger := &StdLogger{} // *StdLogger 实现了 Logger 和 Logger2， 指针类型，共用

	// 创建 Server 实例
	server := Server{ // 值类型，会拷贝整个结构体
		logger:  logger, // 显式字段赋值
		Logger2: logger, // 匿名字段也要显式赋值（Go 不支持隐式推导）
	}

	// 使用显式字段
	server.logger.Log("Server is running")

	// 使用匿名嵌入接口（可以直接调用其方法）
	server.Print("Hello", 42, true, []string{"a", "b"})

	// 测试多态
	process(server.logger)
}

// LOG: Server is running
// 第1个: 值=Hello, 类型=string
// 第2个: 值=42, 类型=int
// 第3个: 值=true, 类型=bool
// 第4个: 值=[a b], 类型=[]string
// LOG: 处理中...

type ConsoleLogger struct {
	prefix string
}

// (l ConsoleLogger) 是 Log方法的接收者（receiver）。Log 就是 ConsoleLogger 类型的“成员方法”，可以通过该类型的实例来调用。
// 这其实可以理解为：一个隐藏了第一个参数的函数，这个参数就是 (l ConsoleLogger)
func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func (l *ConsoleLogger) SetPrefix(p string) {
	l.prefix = p // 修改字段 → 必须用指针接收者
}
