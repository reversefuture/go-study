# statement VS expression
## Expressions（表达式）
● 含义：通过计算产生一个具体的值。
● 特点：  
  ○ 可以是变量、字面量、函数调用、运算符操作等。
  ○ 能作为其他表达式或语句的组成部分。

## Statements（语句）
● 含义：表示要执行的一个动作，不一定产生值。
● 特点：
  ○ 控制程序流程（如 if, for, switch）。
  ○ 声明变量、函数。
  ○ 表达式可以转换为语句（称为 "表达式语句"，但要求表达式有副作用）。

表达式statement关注值，语句expression关注动作。表达式常作为语句的组成部分，而语句控制程序的执行逻辑。

Go 允许将有副作用的表达式单独作为语句（称为 "表达式语句"），例如：
- 函数调用（如 fmt.Println()，利用其副作用而非返回值）
- 通道发送/接收操作（如 ch <- 1）
- 自增/自减（如 i++）
但纯计算表达式不能直接作为语句：
>x + y        // ❌ 错误：表达式的结果未被使用，无法单独成句

# return
The return expression list may be empty if the function's result type specifies names for its result parameters. function may assign values to them as necessary.
```go
func (devnull) Write(p []byte) (n int, _ error) { // n, _ act as local variables
    n = len(p) // assign values to the retrn expression list
    return // returns the values of these variables.
}
```
A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) **with the same name as a result parameter is in scope at the place of the return**.
```go
func f(n int) (res int, err error) {
    if _, err := f(n-1); err != nil { // err is shadowed by the local variable err defined by if.
        return  // invalid return statement: err is shadowed
        //return res, err  // Fix: 显式返回 res 和 err, 或者重命名err为e.
    }
    return
}
```
## 局部变量遮蔽（Variable Shadowing）出现场景
1. 短变量声明（:=）
2. if、for、switch 等块作用域
3. 函数参数
4. defer 和 go 语句
在 defer 和 go 语句中，函数参数的值在语句执行时确定，但局部变量仍然可能遮蔽外部变量：
```go
var x int = 10
func main() {
    x := 20// shadowed
    defer func() {
        fmt.Println(x) // 输出 20，使用遮蔽后的 x
    }()
}
```

# go 关键字
### **Go 中的特殊关键字：`fallthrough`**
1. **`fallthrough`**：
   - **用途**：在 `switch` 语句中，强制执行下一个 `case`，无论条件是否匹配。
   - **特点**：
     - Go 的 `switch` 默认会匹配一个 `case` 后退出，但 `fallthrough` 会继续执行下一个 `case`。
     - 与 Java 的 `switch` 行为不同，Java 的 `switch` 默认会贯穿执行多个 `case`，除非使用 `break`。
   - **示例**：
     ```go
     switch x {
     case 1:
         fmt.Println("x is 1")
         fallthrough
     case 2:
         fmt.Println("x is 2") // 即使 x 不等于 2，也会执行
     }
     ```

### **其他 Go 特有的关键字**：
   - `defer`：用于延迟执行函数调用，通常用于资源清理。
   - `go`：用于启动一个新的 goroutine（轻量级线程）。
   - `range`：用于遍历数组、切片、map 或通道。
   - `select`：用于多路复用通道操作。
   - `chan`：用于定义通道类型。
   - `interface`：定义接口类型（与 JS 的 `interface` 类似，但是go也可以用做js type）。

---

### **Java 中有而 Go 中没有的关键字**
1. **`class`**：
   - Java 是面向对象的语言，所有代码都必须在类中定义。
   - Go 没有类的概念，但可以通过结构体和方法实现类似的功能。

2. **`extends` 和 `implements`**：
   - Java 使用 `extends` 表示类的继承，`implements` 表示接口的实现。
   - Go 没有继承机制，但可以通过嵌入结构体和接口组合实现类似的功能。

3. **`public`、`protected`、`private`**：
   - Java 使用这些关键字控制访问权限。
   - Go 使用大小写区分访问权限：首字母大写的标识符是公开的，小写的是私有的。

4. **`static`**：
   - Java 使用 `static` 定义类级别的变量或方法。
   - Go 没有 `static` 关键字，但可以使用包级别的变量和函数实现类似的功能。

5. **`synchronized`**：
   - Java 使用 `synchronized` 实现线程同步。
   - Go 使用 goroutine 和通道（`chan`）实现并发和同步。

---


# Redeclaration and reassignment
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat() // This duplication is legal: err is declared by the first statement, but only re-assigned in the second
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)

# For
It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.

for key, value := range oldMap {
    newMap[key] = value
}

If you only need the first item in the range (the key or index), drop the second:

for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:

sum := 0
for _, value := range array {
    sum += value
}

String: breaking out individual Unicode code points by parsing the UTF-8 with Erroneous encodings consume one byte and produce the replacement **rune** U+FFFD.
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}

Finally, Go has no comma operator and ++ and -- are statements not expressions. Thus if you want to run multiple variables in a for you should use parallel assignment (although that precludes ++ and --).

// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}

# Switch
Use with logical operators:
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

Use with cases presented in comma-separated lists.

func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}