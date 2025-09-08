在 Go 语言中，查找字符串中是否包含匹配的子串有多种方法。以下是**常用方法的汇总、使用方式、性能对比和适用场景分析**。

---

## ✅ 一、常用查找子串的方法

| 方法 | 包/函数 | 特点 |
|------|--------|------|
| `strings.Contains` | `strings` | 简单精确匹配 |
| `strings.Index` | `strings` | 返回位置，-1 表示未找到 |
| `strings.ContainsAny` | `strings` | 是否包含任意一个字符 |
| `strings.ContainsRune` | `strings` | 是否包含某个 rune |
| `strings.ContainsAny` | `strings` | 是否包含任意字符 |
| `strings.HasPrefix` / `HasSuffix` | `strings` | 前缀/后缀匹配 |
| 正则表达式 `regexp.MatchString` | `regexp` | 支持复杂模式 |
| `index/suffixarray` | `index/suffixarray` | 多次搜索大文本高效 |
| `bytes.Contains` | `bytes` | 对 `[]byte` 高效操作 |

---

### 1. `strings.Contains(s, substr string) bool`
**用途**：判断字符串 `s` 是否包含子串 `substr`。

```go
result := strings.Contains("gopher", "go") // true
```

- ✅ 最常用、最直观
- ✅ 性能好（底层使用 Boyer-Moore 启发式）
- ❌ 仅支持精确匹配

---

### 2. `strings.Index(s, substr string) int`
**用途**：返回子串首次出现的位置，`-1` 表示未找到。

```go
pos := strings.Index("hello world", "world") // 6
if pos != -1 {
    fmt.Println("found at", pos)
}
```

- ✅ 可获取位置信息
- ✅ 比 `Contains` 稍慢一点（返回位置）
- ✅ 适合需要定位的场景

> `strings.Contains(s, substr)` 实际上是 `strings.Index(s, substr) >= 0` 的封装。

---

### 3. `strings.ContainsAny(s, chars string) bool`
**用途**：判断 `s` 是否包含 `chars` 中**任意一个字符**。

```go
strings.ContainsAny("hello", "aeiou") // true (包含 e, o)
strings.ContainsAny("bcdfg", "aeiou") // false
```

- ❗ 注意：不是子串，是字符集合
- ✅ 适合过滤含特定字符的字符串

---

### 4. `strings.HasPrefix(s, prefix)` / `HasSuffix(s, suffix)`
**用途**：判断前缀或后缀。

```go
strings.HasPrefix("https://example.com", "https://") // true
strings.HasSuffix("data.txt", ".txt")                // true
```

- ✅ 高效（直接比较前/后 n 字节）
- ✅ 常用于 URL、文件名处理

---

### 5. 正则表达式：`regexp.MatchString(pattern, s)`
**用途**：使用正则进行模糊/模式匹配。

```go
matched, _ := regexp.MatchString(`\d+`, "age: 25") // true（包含数字）
```

- ✅ 支持复杂模式（邮箱、数字、模糊匹配等）
- ❌ 性能较差（编译正则 + 执行）
- ❌ 错误处理必须检查 `error`

> 推荐预编译正则表达式以提高性能：

```go
var numReg = regexp.MustCompile(`\d+`)
if numReg.MatchString("age: 25") { ... }
```

---

### 6. `index/suffixarray`：适用于**大文本 + 多次搜索**
**用途**：构建后缀数组，支持快速多次查找。

```go
idx, _ := suffixarray.New([]byte(largeText))
positions := idx.Lookup([]byte("search"), -1) // 所有匹配位置
found := len(positions) > 0
```

- ✅ 多次搜索极快（O(m log n) 每次）
- ⏱ 构建慢，内存高
- ✅ 适合：搜索引擎、日志分析等场景
- ❌ 单次搜索不推荐

---

### 7. `bytes.Contains([]byte, []byte) bool`
**用途**：对 `[]byte` 类型进行子串查找。

```go
data := []byte("hello world")
found := bytes.Contains(data, []byte("world"))
```

- ✅ 高效，避免字符串转换开销
- ✅ 适合处理二进制数据或已为 `[]byte` 的场景
- ✅ 性能略优于 `strings.Contains`（当数据本就是 `[]byte`）

