# 错误处理
 https://learnku.com/docs/the-little-go-book/error_handling/3321

# 尽管 Go 有一个垃圾回收器，一些资源仍然需要我们显式地释放他们。
```go
func main() {
	file, err := os.Open("go.sum")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 无论什么情况，在函数返回之后（本例中为 main() ），defer 将被执行
	defer file.Close()
	// 读取文件
}
```
# 带初始化的 if
```go
	if _, err := process2(); err != nil {
		fmt.Print(err)
	}
```

# 格式化
> go fmt ./...

# 空接口
Go 没有继承，也没有这样一个超类。不过他确实有一个没有任何方法的空接口： interface{}。因为空接口没有方法，可以说所有类型都实现了空接口，并且由于空接口是隐式实现的，因此**每种类型都满足空接口契约**。
```go
func add(a interface{}, b interface{}) interface{} {
  ...
}
```
为了将一个接口变量转化为一个显式的类型，又可以用 .(TYPE)：
```go
return a.(int) + b.(int)
```
如果底层类型不是 int，上面的结果将是 error。

你也可以访问强大的类型转换：
```go
switch a.(type) {
  case int:
    fmt.Printf("a is now an int and equals %d\n", a)
  case bool, string:
    // ...
  default:
    // ...
}
```

# 字符串和字节数组
字符串和字节数组是紧密相关的。我们可以轻松地在他们之间转换：
```go
stra := "the spice must flow"
byts := []byte(stra)
strb := string(byts)
```

这种转换方式在各种类型之间是通用的。一些函数显式地需要一个 int32 或者 int64 或者它们的无符号部分。你可能发现你必须这样做：
>int64(count)

然而，当它涉及到字节和字符串时，这可能是你经常做的事情。一定记着当你使用 []byte(X) 或者 string(X) 时，你实际上**创建了数据的副本**。这是必要的，因为**字符串是不可变的**。


那些由 Unicode 码点 runes 构成的字符串，如果你获取字符串的长度，你可能不能得到你期望的。下面的结果是 3：
```go
	s := "椒"
	fmt.Println(len(s))                    //3
	fmt.Println(len([]byte(s)))            //3，拷贝了新数组
	fmt.Println(len([]rune(s)))            // 输出: 1 ✅ 正确的“字符个数”，拷贝了新数组
	fmt.Println(utf8.RuneCountInString(s)) // 输出: 1 ✅ 推荐（更高效，不分配内存）
```
如果你用 range 迭代一个字符串，你将得到 runes，而不是字节。当然，当你将字符串转换为 []rune 类型时，你将得到正确的数据。或者用utf8.RuneCountInString

# 函数类型
>type Add func(a int, b int) int

它可以用在任何地方 -- 作为字段类型，参数或者返回值。
```go
type Add func(a int, b int) int

func main() {
  fmt.Println(process(func(a int, b int) int{
      return a + b
  }))
}

func process(adder Add) int {
  return adder(1, 2)
}
```

# **类型断言（Type Assertion）** 
是一种从 `interface{}`（或任何接口类型）中提取其**底层具体类型值**的操作。

因为接口变量可以存储多种类型的值，当你知道它当前“实际”是什么类型时，就可以用类型断言来“取出来”。

---

## 🔹 一、基本语法

```go
x.(T)
```
- 在运行时检查接口中保存的动态类型是否为 T
- 只能用于 接口类型
- 不能用于两个非接口类型之间的转换
---

## ✅ 二、两种写法

### 1. **直接断言（可能 panic）**

```go
value := x.(int)
```

- 如果 `x` 的实际类型确实是 `int`，返回值。
- 如果不是，**程序 panic（崩溃）**

📌 示例：

```go
var x interface{} = 42
v := x.(int)        // 成功，v = 42
fmt.Println(v)
```

```go
var x interface{} = "hello"
v := x.(int)        // ❌ panic: interface is string, not int
```

---

### 2. **安全断言（推荐）——带两个返回值**

```go
value, ok := x.(T)
```

- `value`：如果类型匹配，是转换后的值；否则是 `T` 的零值
- `ok`：布尔值，`true` 表示类型匹配，`false` 表示不匹配

📌 示例：

```go
var x interface{} = "hello"

if v, ok := x.(int); ok {
    fmt.Println("x 是整数:", v)
} else {
    fmt.Println("x 不是整数") // 会输出这行
}
```

这样不会 panic，适合用于不确定类型时的安全检查。

---

## 🔹 三、使用场景

### ✅ 场景1：从 `interface{}` 中取出具体值

```go
func printValue(v interface{}) {
    if str, ok := v.(string); ok {
        fmt.Println("字符串:", str)
    } else if n, ok := v.(int); ok {
        fmt.Println("整数:", n)
    } else {
        fmt.Println("未知类型")
    }
}

printValue("hello")  // 字符串: hello
printValue(100)      // 整数: 100
```

---

### ✅ 场景2：与 `map[string]interface{}` 配合（如解析 JSON）

```go
data := map[string]interface{}{
    "name":  "Alice",
    "age":   30,
    "active": true,
}

name := data["name"].(string)           // 已知是 string
age, ok := data["age"].(int)            // 安全断言
if ok {
    fmt.Println("年龄:", age)
}
```

---

### ✅ 场景3：配合 `switch` 使用 —— 类型 switch
```go
func describe(x interface{}) {
    switch v := x.(type) {
    case string:
        fmt.Printf("字符串: %s\n", v)
    case int:
        fmt.Printf("整数: %d\n", v)
    case bool:
        fmt.Printf("布尔: %t\n", v)
    case nil:
        fmt.Println("nil")
    default:
        fmt.Printf("未知类型: %T\n", v)
    }
}
```

> `x.(type)` 是类型断言在 `switch` 中的特殊语法，专门用于判断类型。

---


# 类型转换（Type Conversion）

### ✅ 用途：
将一种类型 **显式转换为另一种类型**，前提是 Go 允许这种转换（如数值类型、切片 ↔ 数组指针、字符串 ↔ 字节切片等）。

### 📌 语法：
```go
T(x)
```

### ⏱️ 编译时 / 运行时行为：
- 大部分转换在编译时完成
- 转换前后类型必须是 **可兼容的**

### ✅ 适用场景：
- 数值类型转换：`int(float64)`
- 字符串和字节切片互转：`[]byte("hello")`, `string([]byte{72, 101})`
- 指针类型转换（需 `unsafe` 包，谨慎使用）

### 🚫 限制：
- 不能转换任意类型（比如 `int` 不能直接转 `string`）
- 结构体之间不能直接转换，除非字段完全一致且类型别名

### ✅ 示例：
```go
f := 3.14
i := int(f)                    // float64 → int

s := "hello"
b := []byte(s)                 // string → []byte

b2 := []byte{72, 101, 108, 108, 111}
s2 := string(b2)               // []byte → string
```

---

# 反射（Reflection）

### ✅ 用途：
在 **运行时动态地** 获取变量的类型信息和值，并可以操作它，甚至调用方法、修改字段（如果可寻址）。

### 📌 核心包：
```go
import "reflect"
```

### 📌 主要函数：
```go
reflect.TypeOf(i)   // 获取类型
reflect.ValueOf(i)  // 获取值
```

### ⏱️ 运行时行为：
- 完全在运行时进行
- 性能较低（比类型断言和转换慢很多）
- 非常灵活，但也容易出错

### ✅ 适用场景：
- 编写通用库（如 ORM、序列化库 json/xml）
- 处理未知结构的数据
- 实现“泛型前”的通用函数（Go 1.18 前常见）

### 🚫 限制：
- 代码复杂、易错
- 性能开销大
- 不能违反 Go 的访问规则（如访问私有字段）

### ✅ 示例：
```go
import "reflect"

var x interface{} = "hello"
t := reflect.TypeOf(x)   // string
v := reflect.ValueOf(x)  // hello

fmt.Println("类型:", t)
fmt.Println("值:", v.String())
```

更复杂示例：修改变量（需传指针）

```go
i := 10
v := reflect.ValueOf(&i) // 传指针才能修改
v.Elem().SetInt(20)      // 修改 i 的值为 20
fmt.Println(i)           // 输出 20
```

---

# 三者对比表

| 特性 | 类型断言 | 类型转换 | 反射 |
|------|----------|----------|--------|
| **用途** | 从接口提取具体类型 | 类型间显式转换 | 运行时动态获取类型和值 |
| **语法** | `x.(T)` | `T(x)` | `reflect.ValueOf(x)` |
| **性能** | 较快 | 极快（编译时） | 慢（运行时解析） |
| **安全性** | 不安全（可能 panic），推荐用 `ok` 形式 | 安全（编译时检查） | 易出错，需小心使用 |
| **使用场景** | 处理 `interface{}` | 数值、字符串、切片等转换 | 通用库、序列化、动态调用 |
| **能否操作任意类型** | ❌ 只能用于接口 | ❌ 仅支持预定义可转换类型 | ✅ 几乎可以操作任何类型 |
| **是否需要知道类型** | ✅ 需要指定 `T` | ✅ 需要指定目标类型 | ❌ 可以在运行时发现类型 |
| **是否修改值** | ❌ 不能直接修改 | ✅ 可以（如转换后赋值） | ✅ 可以（如果可寻址） |

---

## 如何选择？

| 你想做什么 | 推荐方式 |
|-----------|----------|
| 从 `interface{}` 中取 `int` 或 `string` | ✅ 类型断言 `x.(int)` |
| 把 `string` 转成 `[]byte` | ✅ 类型转换 `[]byte(s)` |
| 判断一个接口变量是哪种类型，并分别处理 | ✅ 类型 switch `switch v := x.(type)` |
| 写一个能打印任何类型变量的函数 | ✅ 反射 `fmt.Printf("%#v", x)` 或 `reflect` |
| 实现一个通用的校验器（如验证结构体字段） | ✅ 反射 |
| 提高性能关键路径上的类型处理 | ❌ 避免反射，优先用类型断言或泛型 |

---

## 现代 Go 的替代方案：泛型（Go 1.18+）

从 Go 1.18 开始，引入了 **泛型**，很多原本需要反射或类型断言的场景，现在可以用更安全、高效的方式实现：

```go
func Print[T any](v T) {
    fmt.Println(v)
}

Print("hello") // T = string
Print(42)      // T = int
```

✅ 优势：
- 编译时检查
- 无运行时开销
- 比反射清晰安全


