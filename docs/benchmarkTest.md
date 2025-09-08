在 Go 语言中，**基准测试（Benchmarking）** 是性能测试的核心工具，用于精确测量函数、方法或代码段的执行效率，包括：

- 执行时间（纳秒/操作）
- 内存分配（字节/操作、分配次数）
- 并发性能（可选）

Go 的基准测试集成在标准库 `testing` 中，通过 `go test` 命令运行，非常轻量、易用、标准化。

---

## 🧩 一、基准测试函数的基本结构

```go
func BenchmarkXxx(b *testing.B) {
    // 准备工作（不计时）
    b.ResetTimer() // 从此开始计时

    for i := 0; i < b.N; i++ {
        // 被测试的代码
    }
}
```

### 关键点：

- 函数名必须以 `Benchmark` 开头。
- 参数是 `*testing.B`，框架通过它控制循环次数 `b.N`。
- `b.N` 是 Go 自动调整的（100, 1000, 10000, ...），直到获得稳定测量。
- `b.ResetTimer()` 用于排除初始化开销。
- 可选：`b.StopTimer()` / `b.StartTimer()` 控制计时区间。

---

## 📦 二、常用辅助方法

| 方法 | 作用 |
|------|------|
| `b.ResetTimer()` | 重置时间和内存计数器，从当前位置开始计时 |
| `b.StopTimer()` | 暂停计时（用于准备数据等） |
| `b.StartTimer()` | 恢复计时 |
| `b.ReportAllocs()` | 强制报告内存分配统计（即使为0） |
| `b.SetBytes(n)` | 设置每次操作处理的字节数，输出中会显示 `MB/s` |

---

## 🧪 三、实战示例：不同字符串拼接方式对比

```go
// file: string_bench_test.go
package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// 1. 使用 + 拼接
func BenchmarkStringPlus(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 10; j++ {
			s += "golang"
		}
		_ = s // 防止编译器优化掉
	}
}

// 2. 使用 fmt.Sprintf
func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 10; j++ {
			s = fmt.Sprintf("%s%s", s, "golang")
		}
		_ = s
	}
}

// 3. 使用 strings.Builder（推荐）
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 10; j++ {
			builder.WriteString("golang")
		}
		_ = builder.String()
	}
}

// 4. 使用 bytes.Buffer
func BenchmarkStringBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < 10; j++ {
			buf.WriteString("golang")
		}
		_ = buf.String()
	}
}

// 5. 使用 strings.Join
func BenchmarkStringJoin(b *testing.B) {
	b.ResetTimer()
	parts := make([]string, 10)
	for i := range parts {
		parts[i] = "golang"
	}
	for i := 0; i < b.N; i++ {
		s := strings.Join(parts, "")
		_ = s
	}
}
```

---

## ▶️ 四、运行基准测试

### 1. 运行所有基准测试：

```bash
go test -bench=.
```

### 2. 运行特定测试：

```bash
go test -bench=BenchmarkStringBuilder
```

### 3. 显示内存分配：

```bash
go test -bench=. -benchmem
```

### 4. 输出详细结果到文件：

```bash
go test -bench=. -benchmem -benchtime=3s > bench.txt
```

> `benchtime=3s`：每个测试最少运行 3 秒（默认 1 秒），结果更稳定。

---

## 📊 五、典型输出解读

```text
BenchmarkStringPlus-8          1000000     1050 ns/op     512 B/op     10 allocs/op
BenchmarkStringSprintf-8        500000     2400 ns/op    1024 B/op     20 allocs/op
BenchmarkStringBuilder-8      10000000      120 ns/op       0 B/op      0 allocs/op
BenchmarkStringBuffer-8       10000000      130 ns/op       0 B/op      0 allocs/op
BenchmarkStringJoin-8         20000000       80 ns/op       0 B/op      0 allocs/op
```

- `-8`：使用 8 个 CPU（GOMAXPROCS）
- `1000000`：循环次数（b.N）
- `1050 ns/op`：每次操作平均耗时 1050 纳秒
- `512 B/op`：每次操作分配 512 字节内存
- `10 allocs/op`：每次操作发生 10 次堆内存分配

✅ **结论**：`strings.Join` 最快，`strings.Builder` 最灵活且高效，避免使用 `+` 或 `fmt.Sprintf` 在循环中拼接字符串。

---

## 🚀 六、进阶技巧

### 1. 并发基准测试

```go
func BenchmarkConcurrent(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // 被测试代码
        }
    })
}
```

### 2. 防止编译器优化（无副作用代码被优化掉）

```go
result := someFunc()
_ = result // 或使用全局变量
var sink interface{}
sink = result // “沉没”结果
```

### 3. 子基准测试（分组测试）

```go
func BenchmarkGroup(b *testing.B) {
    tests := []struct {
        name string
        fn   func()
    }{
        {"fast", func() { /* ... */ }},
        {"slow", func() { /* ... */ }},
    }

    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                tt.fn()
            }
        })
    }
}
```

输出：

```text
BenchmarkGroup/fast-8
BenchmarkGroup/slow-8
```

---

## ⚠️ 七、常见陷阱

| 陷阱 | 说明 | 修正 |
|------|------|------|
| 忘记 `b.ResetTimer()` | 初始化时间被计入 | 在准备后调用 `ResetTimer()` |
| 未使用结果 | 编译器可能优化掉整个循环 | 用 `_ = result` 或全局 sink |
| 测试有副作用 | 每次迭代状态累积，影响结果 | 在循环内重置状态 |
| 未启用内存统计 | 看不到内存分配情况 | 加 `-benchmem` 参数 |

---

## ✅ 八、最佳实践

1. **命名清晰**：`BenchmarkXxxYyy`，如 `BenchmarkJSONMarshalLargeStruct`
2. **隔离测试**：每次迭代应是独立的、无状态的
3. **报告内存**：始终使用 `-benchmem` 查看内存开销
4. **多次运行**：使用 `-count=N` 获取平均值，减少波动
5. **对比版本**：用 `benchstat` 工具比较优化前后的性能差异

> 安装 benchstat：  
> `go install golang.org/x/perf/cmd/benchstat@latest`

---

## 📌 九、推荐工具

- `benchstat`：对比多组基准测试结果，计算差异和显著性
- `pprof`：配合基准测试做性能剖析（CPU、内存）
- `go-cmp` + 自定义测试：用于更复杂的性能回归测试

---

## ✅ 总结一句话：

> **Go 基准测试 = `BenchmarkXxx(b *testing.B)` + `for i := 0; i < b.N; i++` + `go test -bench=. -benchmem`**

它是你优化性能、防止退化、选择最佳实现的“量化武器”。

---

如需我为你生成针对特定场景（如 JSON 序列化、数据库查询、并发 map 等）的基准测试模板，欢迎随时告诉我！🎯