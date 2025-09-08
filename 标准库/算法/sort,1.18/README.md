# sort

让目标对象实现特定接口，以支持排序。</br>
内部实现了 `QuickSort`、`HeapSort`、`InsertionSort`、`SymMerge`算法。

&nbsp;

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []int{5, 2, 6, 3, 1, 4}
    sort.Ints(s)
    fmt.Println(s)
}
```

&nbsp;

## Slice

通过自定义函数，选择要比较的内容，或改变次序。

&nbsp;

```go
func Slice(x any, less func(i, j int) bool)
func SliceStable(x any, less func(i, j int) bool)

func SliceIsSorted(x any, less func(i, j int) bool)
```

&nbsp;

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []struct{
        id   int
        name string
    }{
        {5, "a"},
        {2, "b"},
        {6, "c"},
        {3, "d"},
        {1, "e"},
        {4, "f"},
    }

    sort.Slice(s, func(i, j int) bool {
        return s[i].id > s[j].id
    })

    fmt.Println(s)
}

// [{6 c} {5 a} {4 f} {3 d} {2 b} {1 e}]
```

&nbsp;

## interface

避开辅助函数，实现排序接口。

```go
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int

    // Less reports whether the element with index i 
    // must sort before the element with index j.
    Less(i, j int) bool

    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

func Sort(data Interface)       // 不稳定排序，不保证相等元素原始次序不变
func Stable(data Interface)     // 稳定排序，相等元素原始次序不变
```

```go
package main

import (
    "fmt"
    "sort"
)

type Data struct {
    text string
    index int
}

type Queue []Data

func (q Queue) Len() int {
    return len(q)
}

func (q Queue) Less(i, j int) bool {
    return q[i].index < q[j].index
}

func (q Queue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}


func main() {
    q := Queue{
        {"d", 3},
        {"c", 2},
        {"e", 4},
        {"a", 0},
        {"b", 1},
    }

    fmt.Println(sort.IsSorted(q))

    sort.Sort(q)
    fmt.Println(q, sort.IsSorted(q))
}

// false
// [{a 0} {b 1} {c 2} {d 3} {e 4}] true
```

&nbsp;

## Search

排序过后的数据，可用 `Search` 执行二分搜索(binary search)。</br>

返回 `[0,n)` 之间，`f() == true`的最小索引序号。 </br>

可用来查找有序插入位置。如找不到，则返回 `n`。

```go
func Search(n int, f func(int) bool) int
```

```go
package main

import (
    "fmt"
    "sort"
)

type Data struct {
    text string
    index int
}

type Queue []Data

func (q Queue) Len() int {
    return len(q)
}

func (q Queue) Less(i, j int) bool {
    return q[i].index < q[j].index
}

func (q Queue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}

func main() {
    q := Queue{
        {"d", 3},
        {"c", 2},
        {"e", 4},
        {"a", 0},
        {"b", 1},
    }

    sort.Sort(q)
    fmt.Println(q)

    i := sort.Search(len(q), func(index int) bool {
        return q[index].index > 6
    })

    fmt.Println("index > 6: ", i)

    i = sort.Search(len(q), func(index int) bool {
        return q[index].index >= 3
    })

    fmt.Println("index >= 3:", i)

    s := make(Queue, len(q) + 1)
    copy(s, q[:i])
    copy(s[i+1:], q[i:])
    s[i] = Data{"a3", 3}

    fmt.Println(s)
}

/*
[{a 0} {b 1} {c 2} {d 3} {e 4}]
index > 6:  5
index >= 3: 3
[{a 0} {b 1} {c 2} {a3 3} {d 3} {e 4}]
*/
```

&nbsp;

## Reverse

辅助函数 `Reverse` 返回一个将 `Less` 参数对调的包装对象。</br>
如此，判断结果就正好相反，实现倒序。

```go
// sort.go

type reverse struct {// embed sort的Interface，定义reverse 
    Interface
}

func (r reverse) Less(i, j int) bool {
    return r.Interface.Less(j, i)
    // return !r.Interface.Less(i,j ) //同上
}
// 将一个实现了 sort.Interface 的数据包装成一个“按反向顺序比较”的新对象，从而实现降序排序。
func Reverse(data Interface) Interface {
    return &reverse{data}
}
```

```go
package main

import (
    "fmt"
    "sort"
)

type Data struct {
    text string
    index int
}

type Queue []Data

func (q Queue) Len() int {
    return len(q)
}

func (q Queue) Less(i, j int) bool {
    return q[i].index < q[j].index
}

func (q Queue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}



func main() {
    q := Queue{
        {"d", 3},
        {"c", 2},
        {"e", 4},
        {"a", 0},
        {"b", 1},
    }

    sort.Sort(sort.Reverse(q))
    fmt.Println(q)
}

// [{e 4} {d 3} {c 2} {b 1} {a 0}]
```

> `Reverse` 是一个**装饰器（decorator）**，它通过包装 `Interface` 并反转 `Less` 方法，让 `sort.Sort` 能够按**从大到小**的顺序排序。


# slice pointer
当你把一个切片作为值传递时，**传递的是这个结构体的副本**，但它里面的 `data` 指针仍然指向**同一个底层数组**。

即使是值接收者，只要类型是 slice、map、channel，都可以“修改”其指向的数据。

Swap 使用值接收者是因为 切片是“引用-like”类型，即使被复制，仍然指向同一个底层数组，因此 x[i], x[j] = x[j], x[i] 能真正交换元素顺序，无需指针。这是 Go 切片设计的便利之处

## ### `Float64Slice` 是什么？

```go
type Float64Slice []float64
```

它是一个切片类型。而 Go 的切片结构本质上是这样的（简化）：

```go
struct {
    data *float64  // 指向底层数组
    len  int
    cap  int
}
```

## ❓ 那什么时候需要指针接收者？

| 类型 | 是否需要指针接收者来修改数据 |
|------|------------------------|
| `[]T`（切片） | ❌ 不需要（值接收者即可修改元素） |
| `map[T]V` | ❌ 不需要 |
| `chan T` | ❌ 不需要 |
| `struct`（普通结构体） | ✅ 需要指针才能修改字段 |
| `string`, `int`, `array` 等值类型 | ✅ 需要指针 |

📌 切片、map、channel 是“**头包含指针**”的类型，值拷贝后仍能操作同一份数据。

---

## go 方法集规则回顾
> T: 所有接收者为T的方法
> *T: 所有接收者为T和*T的方法
```go
data := sort.IntSlice(nums)        // data 是 IntSlice 类型
ptr := &data                       // ptr 是 *IntSlice

var _ sort.Interface = data        // ❌ 编译错误！IntSlice 没实现 Swap
var _ sort.Interface = ptr         // ✅ 正确，*IntSlice 实现了全部方法
```