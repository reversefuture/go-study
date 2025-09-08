package main

// go test -bench=.
// 运行特定测试
// go test -bench=BenchmarkStringBuilder
import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// 不要使用 + 和 fmt.Sprintf 操作字符串，用 bytes.NewBufferString ！
// bytes.Buffer 是一个高效的字节缓冲区，支持动态增长，常用于构建字符串或字节流，避免频繁内存分配。

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
