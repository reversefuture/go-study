# 命名
## getter
不要加Get，用uppercase即可
```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```
## interface
按照约定，只包含一个方法的接口应当以该方法的名称加上 - er 后缀来命名，如 Reader、Writer、 Formatter、CloseNotifier 等
请将字符串转换方法命名为 String 而非 ToString。

# 分号
规则是这样的：若在新行前的最后一个标记为标识符（包括 int 和 float64 这类的单词）、数值或字符串常量之类的基本字面或以下标记之一
```go
break continue fallthrough return ++ -- ) }
```
则词法分析将始终在该标记后面插入分号

分号也可在闭括号之前直接省略

警告：无论如何，你都不应将一个控制结构（if、for、switch 或 select）的左大括号放在下一行

# Redeclaration and reassignment
if 语句不会执行到下一条语句时，亦即其执行体 以 break、continue、goto 或 return 结束时，不必要的 else 会被省略。
```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat() // This duplication is legal: err is declared by the first statement, but only re-assigned in the second
if err != nil {
    f.Close() //  Close() 可能返回错误，而被忽略
    return err
}
codeUsing(f, d)
```
修改：如果 Stat 成功而后续操作失败，Close 的错误也可能需要关注。
```go
f, err := os.Open(name)
if err != nil {
    return err
}

defer func() {
    closeErr := f.Close()
    if err == nil { // 只有在没有其他错误时才返回 Close 的错误
        err = closeErr // 覆盖原来err
    }
}()

d, err := f.Stat()
if err != nil {
    return err
}

codeUsing(f, d)
return nil
```
更多情况不用关注close错误，直接：
> defer f.Close()

# For
It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.
```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```
If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
```go
for key, value := range oldMap {
    newMap[key] = value
}
```
If you only need the first item in the range (the key or index), drop the second:
```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```
If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
```go
sum := 0
for _, value := range array {
    sum += value
}
```
String: breaking out individual Unicode code points by parsing the UTF-8 with Erroneous encodings consume one byte and produce the replacement **rune** U+FFFD.
```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
Finally, Go has no comma operator and ++ and -- are statements not expressions. Thus if you want to run multiple variables in a for you should use **parallel assignment** (although that precludes ++ and --).
```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```
# Switch
Use with logical operators:
```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```
Use with cases presented in comma-separated lists.
```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

由于 if 和 switch 可接受初始化语句， 因此用它们来设置局部变量十分常见。

if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}

# 函数
## 多返回值
## 命名结果参数
Go 函数的返回值或结果 “形参” 可被命名，并作为常规变量使用，就像传入的形参一样。 命名后，一旦该函数开始执行，它们就会被初始化为与其类型相应的零值； 若该函数执行了一条不带实参的 return 语句，则结果形参的当前值将被返回。

## 延迟 
Go 的 defer 语句用于预设一个函数调用（即推迟执行函数）， 该函数会**在执行 defer 的函数返回之前立即执行**

推迟诸如 Close 之类的函数调用有两点好处：第一， 它能确保你不会忘记关闭文件。如果你以后又为该函数添加了新的返回路径时， 这种情况往往就会发生。第二，它意味着 “关闭” 离 “打开” 很近， 这总比将它放在函数结尾处要清晰明了。
```go
// 内容返回文件的内容作为字符串。
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // 我们结束后就关闭了f

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append稍后讨论。
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // 如果我们回到这里，f就关闭了。
        }
    }
    return string(result), nil // 如果我们回到这里，f就关闭了。
}
```

### 参数实时求出但是函数LIFO执行
被推迟函数的实参（如果该函数为方法则还包括接收者）**在推迟执行时就会求值**， 而不是在调用执行时才求值。这样不仅无需担心变量值在函数执行时被改变， 同时还意味着单个已推迟的调用可推迟多个函数的执行。
```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```
**被推迟的函数按照后进先出（LIFO）的顺序执行**，因此以上代码在函数返回时会打印 4 3 2 1 0

这是一个非常经典且重要的 Go 语言问题，我们来**详细解释 `defer` 语句中函数参数的求值时机**，以及它与“推迟执行函数”的关系。

---

## 🔍 核心结论（先记住）：

> ✅ **`defer` 后面函数的实参（arguments）在 `defer` 语句执行时（即注册 `defer` 时）就会求值**，  
> ❌ 但函数本身的执行被推迟到外层函数 `return` 之前。

---

## 🧩 举个例子说明

```go
package main

import "fmt"

func main() {
    x := 10
    defer fmt.Println("deferred:", x) // ①
    x = 20
    fmt.Println("immediate:", x)     // ②
}
```

### 输出：
```
immediate: 20
deferred: 10
```

### 解释：
- `defer fmt.Println("deferred:", x)` 这行代码在执行时，会立即对参数 `x` 求值。
- 此时 `x` 是 `10`，所以 `fmt.Println` 被“绑定”了参数 `10`。
- 虽然后面 `x = 20`，但 `defer` 已经记住了当时的值。
- `defer` 函数在 `main` 函数结束前才执行，但参数早已确定。

---

## 📌 更清晰的例子：参数求值 vs 函数执行

```go
func f() {
    i := 1
    defer fmt.Println(i) // 输出：1，不是 2
    i++
    return
}
```

- `defer fmt.Println(i)` 执行时，`i` 是 `1`，所以参数被求值为 `1`。
- `i++` 不会影响已经求值的参数。
- 最终输出：`1`

---

## 🔄 如果你想“延迟求值”，怎么办？

使用**延迟执行一个匿名函数**（闭包）：

```go
func f2() {
	i := 1
	defer func() { //通过闭包引用变量, 闭包真实执行时i变了
		fmt.Println(i) // 输出：2
	}()
	i++
	return
}
```

### 原因：
- `defer` 推迟的是整个函数调用 `func() { ... }()`
- 参数没有显式传入，而是通过**闭包引用变量 `i`**
- 当 `defer` 执行时，`i` 已经是 `2` 了

> ⚠️ 注意：这是通过闭包“捕获变量”，不是传参！

