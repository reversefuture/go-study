# 协程 
类似于一个线程，但是由 Go 而不是操作系统预定。在 协程 中运行的代码可以与其他代码同时运行。我们来看一个例子：
```go
package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("start")
  go process() //启动协程
  // 没有Sleep主程序会立马执行完退出。因为主进程在退出前 协程 才会有机会去执行（主进程在退出前不会等待全部 协程 执行完毕）
  time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
  fmt.Println("done")
// 匿名函数启动
  go func() {
    fmt.Println("processing2")
    }()
}

func process() {
  fmt.Println("processing")
}
```

协程 易于创建且开销很小。最终多个 协程 将会在同一个底层的操作系统线程上运行。这通常也称为 M:N 线程模型，因为我们有 M 个应用线程（ 协程 ）运行在 N 个操作系统线程上。结果就是，一个 协程 的开销和系统线程比起来相对很低（几 KB）。在现代的硬件上，有可能拥有数百万个 协程 。

# 同步
从变量中读取变量是唯一安全的并发处理变量的方式。 你可以有想要多少就多少的读取者， 但是写操作必须要得同步。 有太多的方法可以做到这个了，包括使用一些依赖于特殊的 CPU 指令集的真原子操作。然而，常用的操作还是使用互斥量（译者注：mutex）:
```go
var (
  counter = 0
  lock sync.Mutex
)

func main() {
  for i := 0; i < 20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond * 10)
}

func incr() {
  lock.Lock()
  defer lock.Unlock()
  counter++
  fmt.Println(counter)
}
```
互斥量序列化会锁住锁下的代码访问。因为默认的的 sync.Mutex 是未锁定状态，这儿我们就得先定义 lock sync.Mutex。

使用这样粗糙的锁操作（覆盖着大量代码的锁操作）确实很诱人，这就违背了我们当初进行并发编程的初心了

使用单个锁时，这没有问题，但是如果你在代码中使用两个或者更多的锁，很容易出现一种危险的情况，当协程 A 拥有锁 lockA ，想去访问锁 lockB ，同时协程 B 拥有锁 lockB 并需要访问锁 lockA

实际上我们使用一个锁时也有可能发生死锁的问题，就是当我们忘记释放它时。 但是这和多个锁引起的死锁行为相比起来，这并不像多锁死锁那样危险（因为这真的 很难发现）
```go
var (
	lock2 sync.Mutex
)

func main() {
	go func() { lock2.Lock() }() // 未释放锁
	time.Sleep(time.Millisecond * 10)
	lock2.Lock() // 死锁
}
```

到现在为止还有很多并发编程我们没有看到过。 首先，有一个常见的锁叫**读写互斥锁**。它主要提供了两种锁功能：一个锁定读取和一个锁定写入。它的区别是允许多个同时读取，同时确保写入是独占的。在 Go 中，**sync.RWMutex** 就是这种锁。另外 sync.Mutex 结构不但提供了 Lock 和 Unlock 方法 ，也提供了 RLock 和 RUnlock 方法；其中 R 代表 Read.。虽然读写锁很常用，它也给开发人员带来了额外的负担：我们不但要关注我们正在访问的数据，还要注意如何访问。



##  `sync.RWMutex` 详解（中文）
- [https://pkg.go.dev/sync#RWMutex](https://pkg.go.dev/sync#RWMutex)

`sync.RWMutex` 是 Go 语言中用于**读写互斥锁**的同步机制。它允许 **多个读操作同时进行**，但写操作是**独占的**——也就是说，当有写操作时，其他读和写都必须等待。

这在**读多写少**的场景中非常有用，能显著提高并发性能，相比普通的 `sync.Mutex`（无论读写都只能一个 goroutine 访问），`RWMutex` 更高效。

---

### 📚 主要方法

`sync.RWMutex` 提供了四类方法：

| 方法 | 用途 | 说明 |
|------|------|------|
| `Lock()` | 获取写锁 | 写者调用，会阻塞所有其他读/写操作 |
| `Unlock()` | 释放写锁 | 必须与 `Lock()` 成对使用 |
| `RLock()` | 获取读锁 | 读者调用，多个读者可同时持有 |
| `RUnlock()` | 释放读锁 | 必须与 `RLock()` 成对使用 |

---

### ✅ 使用示例：并发读写共享数据

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    data = make(map[string]int)
    rwMutex sync.RWMutex
    wg sync.WaitGroup
)

// 读操作
func reader(id int) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        rwMutex.RLock()           // 获取读锁
        value := data["counter"]
        rwMutex.RUnlock()         // 释放读锁
        fmt.Printf("Reader %d 读取值: %d\n", id, value)
        time.Sleep(100 * time.Millisecond)
    }
}

// 写操作
func writer(id int) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        rwMutex.Lock()             // 获取写锁（独占）
        data["counter"]++
        fmt.Printf("Writer %d 将值增加为: %d\n", id, data["counter"])
        rwMutex.Unlock()           // 释放写锁
        time.Sleep(150 * time.Millisecond)
    }
}

func main() {
    data["counter"] = 0

    // 启动 3 个读 goroutine
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go reader(i)
    }

    // 启动 2 个写 goroutine
    for i := 1; i <= 2; i++ {
        wg.Add(1)
        go writer(i)
    }

    wg.Wait()
    fmt.Println("最终值为:", data["counter"])
}
```

**输出示例：**
```
Reader 1 读取值: 0
Writer 1 将值增加为: 1
Reader 2 读取值: 1
Reader 3 读取值: 1
Writer 2 将值增加为: 2
...
最终值为: 6
```

---

### ⚠️ 注意事项

1. **读写锁不能嵌套或升级**
   - 不能在持有 `RLock()` 的情况下再调用 `Lock()`（会死锁）。
   - 没有“读锁升级为写锁”的机制。

2. **必须成对调用**
   - 每次 `RLock()` 必须对应一个 `RUnlock()`，否则会 panic。
   - 同样，`Lock()` 和 `Unlock()` 也要配对。

3. **写操作优先**
   - 一旦有写者在等待，新的读者会被阻塞，防止写者“饿死”。

4. **性能权衡**
   - `RWMutex` 内部比 `Mutex` 复杂，单次操作开销略大。
   - 只有在**读远多于写**时才推荐使用。

5. **不可重入**
   - 同一个 goroutine 不能多次 `RLock()` 或 `Lock()`。


---

### 💡 最佳实践

- 使用 `defer` 确保解锁：

```go
rwMutex.RLock()
defer rwMutex.RUnlock()
// 安全读取数据
```

```go
rwMutex.Lock()
defer rwMutex.Unlock()
// 安全写入数据
```

- 对于简单的并发 map，也可以考虑使用 `sync.Map`，但 `RWMutex` 更灵活可控。

---

### 🆚 `RWMutex` vs `Mutex`

| 特性 | `Mutex` | `RWMutex` |
|------|--------|-----------|
| 多个读同时进行 | ❌ | ✅ |
| 写操作独占 | ✅ | ✅ |
| 适合读多写少 | ❌ | ✅ |
| 开销 | 小 | 稍大 |
| 使用复杂度 | 简单 | 中等 |

---

# 通道
并发编程的最大挑战源于数据的共享。 唯一的共享状态是通道

通道在共享不相关数据的情况下，让并发编程变得更健壮。通道是协程之间用于传递数据的共享管道。换而言之，一个协程可以通过一个通道向另外一个协程传递数据。因此，在任意时间点，只有一个协程可以访问数据。

>c := make(chan int)

这个通道的类型是 chan int。因此，要将通道传递给函数，我们的函数签名看起来是这个样子的：
>func worker(c chan int) { ... }

通道只支持两个操作：接收和发送。可以这样往通道发送一个数据：
>CHANNEL <- DATA

这样从通道接收数据：
>VAR := <-CHANNEL

接收和发送操作是阻塞的。也就是，当我们从一个通道接收的时候， goroutine 将会直到数据可用才会继续执行。类似地，当我们往通道发送数据的时候，goroutine 会等到数据接收到之后才会继续执行。
```go
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    worker := &Worker{id: i}
    go worker.process(c)
  }

  for {
    c <- rand.Int() // 向c发送数据并等待
    time.Sleep(time.Millisecond * 50)
  }
}

type Worker struct {
  id int
}

func (w *Worker) process(c chan int) {
  for { // 循环一值读取c中数据
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
  }
}
```

在某些情况下，你可能需要担心数据被处理掉，这个时候就需要开始阻塞客户端。通道内建这种缓冲容量：
>c := make(chan int, 100)

## 缓冲通道
缓冲通道不会增加容量，他们只提供待处理工作的队列，以及处理突然飙升的任务量的好方法。

通过查看通道的 len 来理解缓冲通道是什么：
```go
for {
  c <- rand.Int()
  fmt.Println(len(c))
  time.Sleep(time.Millisecond * 50)
}
```
你可以看到通道长度一直增加直到满了

## Select
语法上，select 看起来有一点像 switch。使用它，我们提供当通道不能发送数据的时候处理代码
接下来，改变我们的 for 循环：
```go
// 每个 worker 每隔 50ms 尝试向 c 发送一个随机数，如果 c 已满（无法立即发送），就走 default 分支并打印 "dropped"
func (w *Worker) selectProcess(c chan int) {
	for {
		select {
		case c <- rand.Int(): // 如果通道已满或即使未满但调度不及时，case 分支不能立即执行，就会走 default。
		default: //不等待，打印 "dropped"。
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50) //休息：同步、阻塞当前 goroutine 的操作。 在这期间，该 goroutine 进入等待状态，不会被调度执行，相当于yield。
	}
}
```
select 的主要目的是管理多个通道，select 将阻塞直到第一个通道可用。如果没有通道可用，如果提供了 default ，那么他就会被执行。如果多个通道都可用了，随机挑选一个。

## 超时
我们看过了缓冲消息以及简单地将他们丢弃。另一个通用的选择是去超时。我们将阻塞一段时间，但不会永远。

用time.After修改发送器：
```go
	// 尝试发送数据到 c，等待最多 100ms，如果在这期间第一个 case 还不能执行，就走超时分支。
	for {
		select {
		case c <- rand.Int():
			// case <-time.After(time.Millisecond * 100): //返回一个 chan time.Time，100ms 后会向这个通道写入当前时间。
			// 	fmt.Println("timed out")
		case t := <-time.After(time.Millisecond * 100): // 接收After返回的值
			fmt.Println("timed out at", t)
		}
		time.Sleep(time.Millisecond * 50)
	}
```
> 第一个可用的通道被选择。
> 如果多个通道可用，随机选择一个。
> 如果没有通道可用，default 情况将被执行。
> 如果没有 default，select 将会阻塞。

![image](../basics/concurrent.go)

## 最后
话虽如此，我仍然广泛使用 sync 和 sync / atomic 包中的各种同步原语。我觉得比较重要的是通过使用这两种方式比较舒适。我建议你首先关注通道，但是当你遇到一个需要短暂锁的简单示例时，请考虑使用互斥锁或读写互斥锁。
