# 字符串是不可变的

## 🔒 一、Go 字符串为什么是不可变的？（Immutable）

在 Go 中，**字符串是只读的、不可变的字节序列**。一旦创建，就不能修改其内容。

### ✅ 为什么设计成不可变？

这是 Go 的有意设计，主要原因包括：

### 1. **安全性（Security）**
不可变字符串避免了意外或恶意修改。例如：
```go
func process(s string) {
    // 你不需要担心 s 被其他地方修改
}
```

### 2. **并发安全（Goroutine Safe）**
多个 goroutine 可以同时读同一个字符串，**无需加锁**，因为没人能改它。

```go
go func() { fmt.Println(s) }()
go func() { fmt.Println(s) }()
// 安全！字符串不会变
```

### 3. **性能优化**
- 字符串可以**共享底层内存**（比如切片、子串）
- 不需要频繁拷贝（除非你显式操作）
- 哈希值可以缓存（用于 map 的 key）

### 4. **字符串可作为 map 的 key**
因为不可变，Go 才能安全地把 `string` 用作 `map[string]T` 的键。

```go
m := map[string]int{
    "hello": 1,
}
```

如果字符串可变，修改键会导致 map 崩溃。

---

# 用 slice 获取字符串时发生了什么？

比如：

```go
s := "Hello, 世界"
sub := s[0:5] // "Hello"
```

这叫做 **字符串切片（string slicing）**。

### 🧠 实际发生了什么？

Go 的字符串在底层是一个结构体（类似）：

```go
type stringStruct {
    ptr *byte  // 指向底层字节数组的第一个字节
    len int    // 字符串长度
}
```

当你做 `s[0:5]` 时：

✅ Go **不会拷贝数据**，而是：
- 共享原来的字节数组
- 创建一个新的字符串头（string header），指向原数组的某一段

### 🖼 图示：

```
原始字符串 s: "Hello, 世界"
内存: [H][e][l][l][o][,][ ][ä][½][½][é][˜][8][¤]...
       ↑                    ↑
       |                    |
      s[0]                s[7] 开始是“世”的 UTF-8 编码

sub := s[0:5] → 指向 [H][e][l][l][o]
                     ↑         ↑
                     |         |
                   新 ptr    新 len=5
```

👉 所以 `sub` 和 `s` **共享同一块内存**！

---

## ⚠️ 潜在问题：内存泄漏（Memory Leak）

因为共享内存，即使原字符串很大，只要一个小子串还活着，**整个底层数组都不会被释放**。

### 示例：

```go
func bad() string {
    large := strings.Repeat("x", 1000000) + "Hello"
    return large[len(large)-5:] // "Hello"
}
```

返回的 `"Hello"` 只有 5 字节，但它仍然指向一个 1_000_005 字节的大数组！只要这个 `"Hello"` 存活，**整个大内存都不能释放**。

### ✅ 解决方案：显式拷贝

```go
func good() string {
    large := strings.Repeat("x", 1000000) + "Hello"
    sub := large[len(large)-5:]
    return string([]byte(sub)) // 强制拷贝，断开引用
}
```

现在返回的字符串指向**新的内存块**，原大数组可被 GC 回收。

---

## 🔁 修改字符串的唯一方式：创建新字符串

因为字符串不可变，你不能这样做：

```go
s := "Hello"
s[0] = 'h' // ❌ 编译错误！cannot assign
```

✅ 正确做法：转成 `[]rune` 或 `[]byte`，修改后再转回来：

```go
s := "Hello"
bytes := []byte(s)
bytes[0] = 'h'
newS := string(bytes) // "hello"
```

或者处理 Unicode：

```go
runes := []rune("世界")
runes[0] = '人'
newS := string(runes) // "人界"
```

---

## 📊 总结对比

| 操作 | 是否拷贝数据 | 说明 |
|------|---------------|------|
| `s[i:j]` | ❌ 不拷贝（共享内存） | 高效，但可能导致内存泄漏 |
| `string([]byte(s))` | ✅ 拷贝 | 断开共享，安全 |
| `[]rune(s)` | ✅ 拷贝并解码 UTF-8 | 转为 Unicode 码点数组 |
| `len(s)` | ❌ | 返回字节长度（不是字符数） |
| `utf8.RuneCountInString(s)` | ❌ | 返回字符（rune）个数 |

---

## ✅ 最佳实践

1. ✅ 使用 `s[i:j]` 提取子串是高效的，适合临时使用。
2. ❌ 如果子串很小但原串很大，且子串生命周期长 → **用 `string([]byte(sub))` 显式拷贝**。
3. ✅ 处理中文/emoji 用 `[]rune`。
4. ❌ 不要试图修改字符串原地内容。
5. ✅ 字符串做 map key、并发读取都安全。


# string vs byte[]
`[]byte("hello")` 和 `"hello"` 在 Go 中都表示字符串内容，但它们的 **类型和用途** 完全不同。下面是详细的对比：

---

### ✅ 1. 类型不同

| 表达式 | 类型 | 说明 |
|--------|------|------|
| `"hello"` | `string` | 字符串常量，不可变（immutable） |
| `[]byte("hello")` | `[]byte`（字节切片） | 可变的字节序列，底层是 `uint8` 的切片 |

---

### ✅ 2. 是否可变（Mutability）

- `"hello"` 是 **不可变的**：
  ```go
  s := "hello"
  // s[0] = 'H'  // ❌ 编译错误！不能修改字符串中的字符
  ```

- `[]byte("hello")` 是 **可变的**：
  ```go
  b := []byte("hello")
  b[0] = 'H'  // ✅ 合法，现在 b 是 []byte("Hello")
  fmt.Println(string(b)) // 输出: Hello
  ```

---

### ✅ 3. 内存使用和性能

| 情况 | 说明 |
|------|------|
| `string` | 通常用于存储和传递文本，只读，安全，适合做 map 的 key |
| `[]byte` | 适合需要修改内容的场景，比如网络传输、加密、编码转换等 |

⚠️ 转换有开销：
```go
[]byte("hello")  // string → []byte：会复制数据，O(n) 时间
string(b)        // []byte → string：也会复制（在大多数情况下）
```

---

### ✅ 4. 使用场景举例

#### ✅ 使用 `string`（大多数情况）：
```go
name := "Alice"
fmt.Println("Hello", name)
// 用作 map key
users := map[string]int{"Alice": 25, "Bob": 30}
```

#### ✅ 使用 `[]byte`：
```go
// 修改内容
data := []byte("hello")
data[0] = 'H' // → "Hello"

// 传给需要字节切片的函数
json.Unmarshal([]byte(`{"name":"Tom"}`), &person)

// 高频拼接（配合 bytes.Buffer）
var buf bytes.Buffer
buf.Write([]byte("hello"))
buf.Write([]byte("world"))
```

### ✅ 总结对比表

| 特性 | `"hello"` (`string`) | `[]byte("hello")` (`[]byte`) |
|------|------------------------|-------------------------------|
| 类型 | `string` | `[]uint8` |
| 是否可变 | ❌ 不可变 | ✅ 可变 |
| 是否可修改某个字符 | ❌ 不行 | ✅ 可以 `b[0] = 'H'` |
| 转换是否复制数据 | ✅ 是（转换时） | ✅ 是（转换时） |
| 适合场景 | 显示、存储、map key | 修改、网络、加密、I/O |
| 是否可做 map key | ✅ 可以 | ❌ `[]byte` 不能做 map key（因为 slice 不可比较） |

---

### 💡 小技巧：什么时候用哪个？

- ✅ 普通文本处理、打印、配置：用 `string`
- ✅ 需要修改内容、性能敏感的底层操作：用 `[]byte`
- ✅ 调用 `Compare([]byte, []byte)` 这类函数：必须用 `[]byte(...)`

---

### 🔄 相互转换

```go
s := "hello"
b := []byte(s)     // string → []byte

newS := string(b)  // []byte → string
```

> 注意：这两次转换都会**复制底层数据**，不是视图。
