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

# interface
interface 是一种类型，它定义了一组方法签名。任何类型只要实现了这些方法，就自动实现了这个接口 —— 无需显式声明。

## Go 的多态是通过接口实现的。

## interface{} 可以表示任何类型
```go
var x interface{}
x = 42
x = "hello"
x = []int{1, 2, 3}

fmt.Println(x) // [1 2 3]
```
常用于：
>函数接收任意类型参数
> map[string]interface{} 表示 JSON 风格数据
```go
data := map[string]interface{}{
    "name": "Alice",
    "age":  30,
    "tags": []string{"go", "dev"},
}
```

# 继承
在 Go 语言中，**没有传统面向对象语言（如 Java、C++）中的“继承”机制**，Go 不支持类（class）和继承（inheritance）。但 Go 提供了 **组合（composition）** 和 **嵌入（embedding）** 机制，可以实现类似继承的效果，通常称为“**组合继承**”或“**结构体嵌入**”。

---

## 一、通过结构体嵌入实现“继承”效果

Go 使用 **匿名嵌入（Anonymous Embedding）** 来实现类似继承的功能。

### 示例：嵌入结构体

```go
package main

import "fmt"

// 定义一个基础结构体（父类）
type Animal struct {
    Name string
    Age  int
}

func (a *Animal) Speak() {
    fmt.Printf("I am %s, I am %d years old.\n", a.Name, a.Age)
}

// Dog 继承 Animal 的属性和方法
type Dog struct {
    Animal  // 匿名嵌入，相当于继承
    Breed   string
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Golden Retriever",
    }

    // 可以直接调用 Animal 的方法
    dog.Speak() // 输出: I am Buddy, I am 3 years old.

    // 也可以访问嵌入字段
    fmt.Println(dog.Name) // Buddy
}
```

### 特点：
- `Dog` 没有显式定义 `Name`、`Age` 和 `Speak()`，但可以直接使用。
- 这是 **组合优于继承** 的体现，Go 推荐这种方式。

---

## 二、方法重写（模拟多态）

Go 不支持方法重写语法，但可以通过“**同名方法覆盖**”实现类似效果。

```go
func (d *Dog) Speak() {
    fmt.Printf("I am %s the dog, woof!\n", d.Name)
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Golden Retriever",
    }

    dog.Speak() // 调用的是 Dog 的 Speak，而不是 Animal 的
}
```

> 注意：这并不是真正的“重写”，而是方法集的覆盖。调用时优先使用最外层的方法。

如果你想调用父类方法，可以显式调用：

```go
func (d *Dog) Speak() {
    d.Animal.Speak() // 调用“父类”方法
    fmt.Println("Woof!")
}
```

---

## 三、接口实现多态（Go 的多态方式）

Go 的多态是通过 **接口（interface）** 实现的，而不是继承。 **接口定义行为，而不是用结构继承**

```go
type Speaker interface {
    Speak() string
}

func MakeSound(s Speaker) {
    s.Speak()
}

func main() {
    animal := &Animal{Name: "Cat", Age: 2}
    dog := &Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Husky",
    }

    MakeSound(animal) // I am Cat, I am 2 years old.
    MakeSound(dog)    // I am Buddy the dog, woof!
}
```

只要实现了 `Speak()` 方法，就自动满足 `Speaker` 接口，无需显式声明“继承”。

---

## 四、总结：Go 中的“继承”特点

| 特性 | Go 实现方式 |
|------|-------------|
| 属性继承 | 通过匿名嵌入结构体 |
| 方法继承 | 嵌入结构体的方法自动可访问 |
| 方法重写 | 定义同名方法覆盖 |
| 多态 | 通过接口实现 |
| 多重继承 | 支持嵌入多个结构体（多重组合） |

### 多重嵌入示例：

```go
type HasLegs struct {
    Legs int
}

type Dog struct {
    Animal
    HasLegs
    Breed string
}

dog := Dog{
    Animal:  Animal{Name: "Max", Age: 4},
    HasLegs: HasLegs{Legs: 4},
    Breed:   "Poodle",
}
fmt.Println(dog.Legs) // 4
```
