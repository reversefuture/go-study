# 数据
## 使用 new 关键字分配内存
Go 提供了两种分配原语，即内建函数 new 和 make

## new
这是个用来分配内存的内建函数， 但与其它语言中的同名函数不同，它不会 初始化 内存，只会将内存 **置零**。 也就是说，**new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址（指针），也就是一个类型为 *T 的值**。 每种类型的零值就不必进一步初始化了。

### 比如bytes.Buffer， sync.Mutex
eg1
```go
var buf bytes.Buffer
buf.Write([]byte("hello"))

buf := new(bytes.Buffer)
buf.Write([]byte("hello"))
```
- 为什么不需要buf := bytes.NewBuffer(nil)
因为 bytes.Buffer 的零值就是一个空但可用的缓冲区。

它内部的字段（如 []byte 切片）初始为 nil，但在 Write 方法中会自动处理并进行扩容（通过 make 或 append）。

eg2
```go
var mu sync.Mutex
mu.Lock()
// ...
mu.Unlock()
// 或者
mu := new(sync.Mutex)
mu.Lock()
// ...
mu.Unlock()
```
- 为什么不需要 mu.Init()？
因为 sync.Mutex 的零值就是一把已解锁的互斥锁。

eg3
```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```
SyncedBuffer 类型的值也是在声明时就分配好内存就绪了。后续代码中， p 和 v 无需进一步处理即可正确工作。
```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

## 构造函数和复合字面量
复合字面来简化实例属性赋值
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```
实际上，每当获取一个复合字面的地址时，都将为一个新的实例分配内存， 因此我们可以将上面的最后两行代码合并
```go
    return &File{fd, name, nil, 0}
```

复合字面的字段必须按顺序全部列出。但如果以 字段:值 对的形式明确地标出元素，初始化字段时就可以按任何顺序出现，未给出的字段值将赋予零值。
```go
    return &File{fd: fd, name: name}
```

少数情况下，若复合字面不包括任何字段，它将创建该类型的零值。**表达式 new(File) 和 &File{} 是等价的。**返回一个指向新分配的 File 类型零值的指针（即 *File）

复合字面量可用于创建数组、切片和映射，但只有数组和映射可以使用 key: value 的形式进行初始化，其中数组的 key 是索引，映射的 key 是键；而切片不能使用整数作为键来初始化。
```go
const (
	Enone  = iota
	Eio    = iota
	Einval = iota
)

a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"} // ...让编译器自动推断数组长度
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"} // 新建slice
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```

切片的复合字面量只能是:
> []T{value1, value2, ...}

## 使用 make 分配
内建函数 make(T, args) 的目的不同于 new(T)。它只用于创建 slice、map 和 channel，并返回类型为 T（而非 *T）的一个已初始化 （而非置零）的值,不返回指针。 出现这种用差异的原因在于，这三种类型本质上为引用数据类型，它们在使用前必须初始化

> make([]int, 10, 100)
会分配一个具有 100 个 int 的数组空间，接着创建一个长度为 10， 容量为 100 并指向该数组中前 10 个元素的切片结构。（生成切片时，其容量可以省略，更多信息见切片一节。） 与此相反，new([]int) 会返回一个指向新分配的，已置零的切片结构， 即一个指向 nil 切片值的指针。

## 数组
以下为数组在 Go 和 C 中的主要区别。在 Go 中，

- 数组是值。将一个数组赋予另一个数组会复制其所有元素。
- 特别地，若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
- 数组的大小是其类型的一部分。类型 [10]int 和 [20]int 是不同的。
```go
func modify(arr [3]int) {
    arr[0] = 999
    fmt.Println("函数内:", arr)
}

func main() {
    a := [3]int{1, 2, 3}
    fmt.Println("调用前:", a)
    modify(a)
    fmt.Println("调用后:", a)
}
// 调用前: [1 2 3]
// 函数内: [999 2 3]
// 调用后: [1 2 3]
```

数组为值的属性很有用，但代价高昂；若你想要 C 那样的行为和效率，你可以传递一个指向该数组的指针。
```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1} // 让编译器自动推断数组的长度
x := Sum(&array)  // Note the explicit address-of operator
```

## 切片
切片通过对数组进行封装，为数据序列提供了更通用、强大而方便的接口。
切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。 
若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这可以理解为传递了底层数组的指针。
Read 函数可接受一个切片实参 而非一个指针和一个计数；切片的长度决定了可读取数据的上限
>func (f *File) Read(buf []byte) (n int, err error)

只要切片不超出底层数组的限制，它的长度就是可变的，只需将它赋予其自身的切片即可。 切片的容量可通过内建函数 cap 获得
```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // 重新分配
        // 为未来的增长,双重分配所需的内容.
        newSlice := make([]byte, (l+len(data))*2)
        // copy函数是预先声明的，适用于任何切片类型。
        copy(newSlice, slice) // 复制原始slice到扩容后的newSlice
        slice = newSlice // 调整切片的长度到 (l+len(data))*2。
    }
    slice = slice[0:l+len(data)] // 调整切片的长度到 l + len(data)。
    copy(slice[l:], data) // 从slice[l:]原长度位置开始的“空位”开始拷贝data
    return slice
}
```
尽管 Append 可修改 slice 的元素，但**切片自身（其运行时数据结构包含指针、长度和容量）是通过值传递的**。

## 二维切片
```go
type Transform [3][3]float64  // 一个 3x3 的数组，其实是包含多个数组的一个数组。
type LinesOfText [][]byte     // 包含多个字节切片的一个切片。
```

有时必须分配一个二维数组。一种就是独立地分配每一个切片；而另一种就是只分配一个数组， 将各个切片都指向它。
```go
// 分配底层切片.
picture := make([][]uint8, YSize) // y每一行的大小
//循环遍历每一行
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```
现在是一次分配，对行进行切片：
```go
// 分配底层切片
picture := make([][]uint8, YSize) //  每 y 个单元一行。准备了 YSize 个“行指针”，但每行都还是空的nil。
// 一次性分配所有像素数据。一块连续的内存区域，非常高效
pixels := make([]uint8, XSize*YSize) // 指定类型[]uint8, 即便图片是 [][]uint8.
//循环遍历图片所有行，从剩余像素切片的前面对每一行进行切片。
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```
if i=0:
- pixels 当前是全部像素：长度 XSize*YSize
- pixels[:XSize] → 第 0 行（前 XSize 个像素）
- pixels[XSize:] → 剩下的 (YSize-1)*XSize 个像素
- picture[0] = pixels[:XSize]     // 第0行指向前XSize个像素
- pixels = pixels[XSize:]         // pixels 缩短，去掉已分配的行

最终结果
- picture 是一个二维切片：YSize 行，每行 XSize 列
- 所有行共享同一块连续内存（最初由 pixels 分配）
- 所有像素在内存中连续，缓存命中率高，访问速度快，GC 压力小

# 映射
其键可以是任何相等性操作符支持的类型， 如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、结构以及数组。 

若试图通过映射中不存在的键来取值，就会返回与该映射中项的类型对应的零值。int：0， bool: false

有时你需要区分某项是不存在还是其值为零值。你可以使用多重赋值的形式来分辨这种情况。
```go
if seconds, ok := timeZone[tz]; ok {
    return seconds
}
// 仅需判断映射中是否存在某项
_, present := timeZone[tz]
```

要删除映射中的某项，可使用内建函数 delete
> delete(timeZone, "PDT")  // 现在是标准时间


# print
## fmt.Sprintln
- fmt.Sprintln 是标准库函数，功能是：
- 接收多个参数（...interface{}）。
- 将它们转换为字符串，并用空格连接，最后加上换行符。
- 返回一个字符串（string），而不是直接打印到控制台。

## fmt.Fprint 
fmt.Fprint 一类的格式化打印函数可接受任何实现了 io.Writer 接口的对象作为第一个实参；变量 os.Stdout 与 os.Stderr 都是人们熟知的例子
```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```

```go
func main() {
    // 整数
    fmt.Printf("十进制: %d\n", 255)
    fmt.Printf("十六进制: %x\n", 255)   // ff
    fmt.Printf("大写十六进制: %X\n", 255) // FF
    fmt.Printf("二进制: %b\n", 255)     // 11111111
    fmt.Printf("八进制: %o\n", 255)     // 377

    // 浮点数
    fmt.Printf("浮点数: %f\n", 3.1415926)
    fmt.Printf("科学计数: %e\n", 123456.0)
    fmt.Printf("自动格式: %g\n", 0.000123)

    // 字符串
    fmt.Printf("字符串: %s\n", "Golang")
    fmt.Printf("带引号: %q\n", "Hello\nWorld")

    // 指针
    x := 10
    fmt.Printf("地址: %p\n", &x)

    // 类型和默认值
    fmt.Printf("类型: %T\n", x)
    fmt.Printf("值: %v\n", x)
    fmt.Printf("Go 语法: %#v\n", x)

    // 宽度和精度
    fmt.Printf("右对齐: |%10d|\n", 42)     // |        42|
    fmt.Printf("左对齐: |%-10d|\n", 42)    // |42        |
    fmt.Printf("补零: %08d\n", 42)         // 00000042
    fmt.Printf("小数: %.2f\n", 3.14159)    // 3.14
}
```


若你想控制自定义类型的默认格式，只需为该类型定义一个具有 String() string 签名的方法
```go
func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c) //7/-2.35/"abc\tdef"
}
```
请勿通过调用 Sprintf 来构造 String 方法，因为它会无限递归你的 String 方法
要解决这个问题也很简单：将该实参转换为基本的字符串类型，它没有这个方法。
```go
type MyString string

func (m MyString) String() string {
	//return fmt.Sprintf("MyString=%s", m) // 错误：会无限递归
	return fmt.Sprintf("MyString=%s", string(m)) // 可以：注意转换
}
```

将打印例程的实参直接传入另一个这样的例程。Printf 的签名为其最后的实参使用了 ...interface{} 类型，这样格式的后面就能出现任意数量，任意类型的形参了。
>  func Printf(format string, v ...interface{}) (n int, err error) {}

在 Printf 函数中，v 看起来更像是 []interface{} 类型的变量，但如果将它传递到另一个变参函数中，它就像是常规实参列表了。如下：
```go
// Println 通过 fmt.Println 的方式将日志打印到标准记录器
func Println(v ...interface{}) { //...interface{}: 可变参数，可以传入任意数量、任意类型的参数。
	log.Output(2, fmt.Sprintln(v...)) // Output takes parameters (int, string)
}
```

... 形参类似用法：求最小值例子
```go
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // 最大的 int == math.MaxInt
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

## append
append 会在切片末尾追加元素并返回结果
> func append(slice []T, elements ...T) []T
你无法在 Go 中编写一个类型 T 由调用者决定的函数。这也就是为何 append 为内建函数的原因：它需要编译器的支持。
```go
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	y := []int{4, 5, 6}
	x = append(x, y...) // 同上
	fmt.Println(x)
```

## ... spread operator
在函数定义中：
func f(args ...T)
表示可变参数（variadic function）

在函数调用中：
f(slice...)
将切片展开为多个参数（unpack）

# 初始化
## 常量
 常量只能是数字、字符（符文）、字符串或布尔值。
 由于编译时的限制， 定义它们的表达式必须也是可被编译器求值的常量表达式。例如 1<<3 就是一个常量表达式，而 math.Sin(math.Pi/4) 则不是，因为对 math.Sin 的函数调用在运行时才会发生。
## 枚举
枚举常量使用枚举器 iota 创建。由于 iota 可为表达式的一部分，而表达式可以被隐式地重复，这样也就更容易构建复杂的值的集合了。
```go
type ByteSize float64

const (
    _           = iota // 通过赋予空白标识符来忽略第一个值
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```
自定义一个 String 方法：为打印时自动格式化任意值提供了可能性，即便是作为一个通用类型的一部分
```go
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
	var testSize ByteSize
	testSize = 10000
	var testSize2 ByteSize = 1000
	testSize3 := ByteSize(100) // 短类型声明+强制类型转换
	fmt.Println(testSize) // 9.77KB
```

## iota
iota 是 Go 中在 const 块内从 0 开始自动递增的“常量计数器”，常用于定义枚举值、位标志等。 
```go
const (
    a = iota  // 0
    b = iota  // 1
    c = iota  // 2
)
// 等价于
const (
    a = 0
    b = 1
    c = 2
)
```

简化写法（省略重复 iota）
```go
const (
    Sunday = iota // 0
    Monday// 1（隐式 = iota）
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday //6
)
```

## init 函数
最后，每个源文件都可以通过定义自己的无参数 init 函数来设置一些必要的状态。每个文件都可以拥有多个 init 函数。

只有该包中的所有变量声明都通过它们的初始化器求值后 init 才会被调用， 而包中的变量只有在所有已导入的包都被初始化后才会被求值。
```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath 可通过命令行中的 --gopath 标记覆盖掉。
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```

# 方法
## 指针 vs. 值
正如 ByteSize 那样，我们可以为任何已命名的类型（除了指针或接口）定义方法； 接收者可不必为结构体。eg:
```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // 主体与上面定义的Append函数完全相同。
}
```
我们仍然需要该方法返回更新后的切片。为了消除这种不便，我们可通过重新定义该方法， 将一个指向 ByteSlice 的指针作为该方法的接收者， 这样该方法就能重写调用者提供的切片了。
```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // 主体同上，只是没有返回值
    *p = slice
}
```
使用：
```go
	var b ByteSlice
	//fmt.Fprintf 在写入后通常会调用其目标的 String() 方法来格式化输出
	fmt.Fprintf(&b, "This hour has %d days\n", 7) // 将&b后面参数合并format后写入&b
	fmt.Print(string(b))                          // This hour has 7 days
```
以指针或值为接收者的区别在于：**值方法可通过指针和值调用， 而指针方法只能通过指针来调用**。

**指针方法可以修改接收者；通过值调用它们会导致方法接收到该值的副本， 因此任何修改都将被丢弃**

不过有个方便的例外：若该值是可寻址的， 那么该语言就会自动插入取址操作符来对付一般的通过值调用的指针方法。在我们的例子中，变量 b 是可寻址的，因此我们只需通过 b.Write 来调用它的 Write 方法，编译器会将它重写为 (&b).Write

**如果一个值是 可寻址的，意味着你可以获取它的内存地址，即可以写 &x。**。 eg:
```go
// 1. 普通变量
x := 10
px := &x // ✅ 可寻址

// 2. 数组元素
arr := [3]int{1, 2, 3}
p := &arr[0] // ✅ 可寻址

// 3. 切片元素
s := []int{1, 2, 3}
p = &s[1] // ✅ 可寻址

// 4. 结构体字段
type Person struct {
    Name string
}
var p1 Person
pn := &p1.Name // ✅ 可寻址

// 5. 指针解引用
pp := &p1
pp.Name = "Alice" // 等价于 (*pp).Name
```

# 接口
Go 中的接口为指定对象的行为提供了一种方法。每种类型都能实现多个接口
```go
type Sequence []int

// sort.Interface所需的方法。
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy方法返回Sequence的复制
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// 打印方法-在打印前给元素排序
func (s Sequence) String() string {
    s = s.Copy() // 复制s，不要覆盖参数本身
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop空间复杂度是O(N²)；将在下个例子中修复它
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

在 Go 程序中，为访问不同的方法集而进行类型转换的情况非常常见. 简化上面：
```go
type Sequence []int

// 打印方法-在打印之前对元素进行排序
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort() // IntSlice attaches the methods of Interface to []int
    return fmt.Sprint([]int(s))
}
```

## 接口转换与类型断言
类型选择 是类型转换的一种形式：它接受一个接口，在选择 （switch）中根据其判断选择对应的情况（case）， 并在某种意义上将其转换为该种类型。
```go
type Stringer interface {
    String() string
}

var value interface{} // Value 由调用者提供
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

类型断言。类型断言接受一个接口值， 并从中提取指定的明确类型的值。
```go
value.(typeName)
```
而其结果则是拥有静态类型 typeName 的新值。该类型必须为该接口所拥有的具体类型， 或者该值可转换成的第二种接口类型。
```go
str := value.(string)
```
但若它所转换的值中不包含字符串，该程序就会以运行时错误崩溃。为避免这种情况， 需使用 “逗号，ok” 惯用测试它能安全地判断该值是否为字符串：
```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

## 通用性
若某种现有的类型仅实现了一个接口，且除此之外并无可导出的方法，则该类型本身就无需导出

在这种情况下，构造函数应当返回一个接口值而非实现的类型。例如在 hash 库中，crc32.NewIEEE 和 adler32.New 都返回接口类型 hash.Hash32

## 接口和方法
我们可以为除指针和接口以外的任何类型定义方法，同样也能为一个函数写一个方法。一个很直观的例子就是 http 包中定义的 Handler 接口

为一个函数写一个方法
```go
// HandlerFunc 类型是一个适配器，
// 它允许将普通函数用做HTTP处理程序。
// 若 f 是个具有适当签名的函数，
// HandlerFunc(f) 就是个调用 f 的处理程序对象。
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```
HandlerFunc 是个具有 ServeHTTP 方法的类型， 因此该类型的值就能处理 HTTP 请求。

# 空白标识符
## 多个参数赋值中的空白标识符
```go
if _, err := os.Stat(path); os.IsNotExist(err) { // _被安全丢弃， IsNotExist用来判断错误类型
    fmt.Printf("%s does not exist\n", path)
}
```
## 未使用的导入和变量
```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf  // 用于调试，结束时删除。
var _ io.Reader    // 用于调试，结束时删除。

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

## 为辅助作用而导入
有时导入某个包只是为了其辅助作用， 而没有任何明确的使用。例如，在 net/http/pprof 包的 init 函数中记录了 HTTP 处理程序的调试信息。只需将包重命名为空白标识符
```go
import _ "net/http/pprof"
```

## 接口检查
在 Go 中，使用 `_`（下划线）进行**接口检查**是一种常见的**编译期接口实现检查**技巧，主要用于确保某个结构体确实实现了某个接口。

---

### ✅ 目的
Go 的接口是 **隐式实现**的（不需要 `implements` 关键字），因此有时你可能误以为某个结构体实现了某个接口，但实际上没有。  
使用 `_` 可以在 **编译时** 静默检查是否实现了接口，避免运行时出错。

---

### ✅ 常见写法：使用 `_` 进行接口实现检查

```go
var _ InterfaceName = (*StructType)(nil) // StructType类型的nil指针
```

或

```go
var _ InterfaceName = StructType{}
```

> 这行代码的含义是：**声明一个匿名变量 `_`，其类型为 `InterfaceName`，并用 `*StructType(nil)` 或 `StructType{}` 初始化**。  
> 如果 `*StructType` 没有实现 `InterfaceName`，编译器会报错。

---

### 🔍 示例说明

假设我们有：

```go
type Writer interface {
    Write([]byte) (int, error)
}

type MyWriter struct{}

func (mw *MyWriter) Write(data []byte) (int, error) {
    // 实现逻辑
    return len(data), nil
}
```

#### ✅ 正确的接口检查写法：

```go
var _ Writer = (*MyWriter)(nil)  // 推荐：检查指针是否实现
```

### ❗ 为什么用 `(*MyWriter)(nil)` 而不是 `MyWriter{}`？

因为方法接收者可能是指针类型（如 `func (mw *MyWriter)`），此时只有指针类型实现了接口，值类型（`MyWriter`）不实现。

所以通常建议用指针形式检查：

```go
var _ Writer = (*MyWriter)(nil)  // 检查 *MyWriter 是否实现 Writer
```

如果写成：

```go
var _ Writer = MyWriter{}  // 检查 MyWriter 值类型是否实现
```

而 `Write` 方法是 `func (mw *MyWriter)`，就会编译失败。


### ✅ 实际项目中的常见用法

通常放在文件末尾或类型定义附近，作为“静态断言”：

```go
type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // ...
}

// 静态检查：*Handler 是否实现了 http.Handler 接口
var _ http.Handler = (*Handler)(nil)
```

---

### ⚠️ 注意事项

- 使用 `_` 不会分配内存或产生运行时开销，只是编译期检查。
- 这种变量不会出现在程序中，仅用于类型检查。
- 推荐在库或关键组件中使用，提高代码健壮性。

### 💡 小技巧：多个接口检查

```go
var (
    _ io.Reader = (*MyReader)(nil)
    _ io.Writer = (*MyWriter)(nil)
    _ http.Handler = (*MyHandler)(nil)
)
```

# 内嵌 Embedding
Go 并不提供典型的，类型驱动的子类化概念，但通过将类型内嵌到结构体或接口中， 它就能 “借鉴” 部分实现。

## 类继承内嵌
![file](../basics/structEmbeddingInherit.go)


内嵌的类型（如 Person）不需要指定字段名，直接写类型名。
外层结构体可以直接访问内嵌结构体的字段和方法。
如果有多个内嵌结构体有同名方法，调用时需显式指定。 
字段或方法 X 会隐藏该类型中更深层嵌套的其它项 X。若相同的嵌套层级上出现同名冲突，通常会产生一个错误
支持多层内嵌。

## 内嵌接口
io 包也导出了一些其它接口，以此来阐明对象所需实现的方法。 例如 io.ReadWriter 就是个包含 Read 和 Write 的接口。我们可以通过显示地列出这两个方法来指明 io.ReadWriter， 但通过将这两个接口内嵌到新的接口中显然更容易且更通用，就像这样：
```go
// ReadWriter 接口结合了 Reader 接口 和 Writer 接口
type ReadWriter interface {
    Reader
    Writer
}
```

区分内嵌与子类的重要手段。当内嵌一个类型时，该类型的方法会成为外部类型的方法， 但当它们被调用时，该方法的接收者是内部类型，而非外部的。

# 错误
按照约定，错误的类型通常为 error，这是一个内置的简单接口
```go
type error interface {
    Error() string
}
```
也可以自定义：
```go
// PathError 记录错误、执行的操作和文件路径
type PathError struct {
    Op string    // "open", "unlink" 等等对文件的操作
    Path string  // 相关文件的路径
    Err error    // 由系统调用返回
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```
若调用者关心错误的完整细节，可使用类型选择或者类型断言来查看特定错误，并抽取其细节
```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // 恢复一些空间。
        continue
    }
    return
}
```

## panic
`panic` 是一个内置函数，用于**触发运行时错误（运行时恐慌）**，导致程序中断正常流程并开始**恐慌（panic）和恢复（recover）机制**。用法：

```go
panic(v interface{})
```

- `v` 可以是任意类型的数据（字符串、错误、结构体等），通常用于描述错误信息。
- 调用 `panic` 后，程序会停止当前函数的执行，并开始**回溯调用栈（unwind stack）**，执行所有已注册的 `defer` 函数。
- 如果没有通过 `recover` 捕获该 panic，程序最终会崩溃并打印错误信息和调用栈。


### 🔹 `panic` 的执行流程

1. 调用 `panic(...)`。
2. 当前函数停止执行。
3. 所有已 `defer` 的函数按**后进先出（LIFO）**顺序执行。
4. 如果某个 `defer` 函数中调用了 `recover()`，可以捕获 panic，恢复正常执行。
5. 如果没有 `recover`，程序崩溃，打印 panic 信息和堆栈跟踪。

### 🔹 结合 `defer` 和 `recover` 捕获 panic
![file](../basics/panic.go)


### 🔹 `panic` 的常见使用场景

1. **不可恢复的错误**（如配置错误、初始化失败）：
   ```go
   if err != nil {
       panic("failed to load config: " + err.Error())
   }
   ```

2. **程序逻辑错误**（如断言失败）：
   ```go
   if user == nil {
       panic("user cannot be nil")
   }
   ```

3. **库内部错误保护**：
   有些库会在遇到严重错误时 panic，提示使用者调用方式错误。

---

### 🔹 `panic` vs `error`

| 特性         | `error`                          | `panic`                            |
|------------|----------------------------------|------------------------------------|
| 用途        | 可预期的错误，应被处理               | 不可预期或严重错误，程序可能无法继续     |
| 是否必须处理 | 是（建议）                         | 否（但可 recover 捕获）              |
| 性能开销     | 低                                | 高（涉及栈展开）                     |
| 推荐场景     | 文件读写、网络请求、用户输入校验等       | 程序初始化失败、逻辑断言、严重状态不一致等 |

> ✅ 原则：**用 `error` 处理常规错误，用 `panic` 处理真正异常的情况**。

---

### 🔹 小贴士

- `recover()` 只能在 `defer` 的函数中调用，否则返回 `nil`。
- `panic` 可以被多次 `recover`，只要在不同的 `defer` 层中处理。
- 在 Go 的 `main` 函数中未捕获的 panic 会导致整个程序退出。

---

### 🔹 进阶示例：嵌套 defer 和 recover

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Caught panic:", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }
    fmt.Println("Result:", a/b)
}
```

调用 `safeDivide(10, 0)` 会输出：

```
Caught panic: division by zero
```
