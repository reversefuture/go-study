# Go 并发编程
Go 语言中的多线程编程主要通过 **goroutines** 和 **channels** 来实现。Go 语言没有直接的“线程”概念，而是使用 goroutines 来并发执行任务。
Go语言的设计哲学是“不要通过共享内存来通信，而要通过通信来共享内存”
调用完成后， 该 Go 协程也会安静地退出。

## 协程
在 go 关键字后面加一个函数，就可以创建一个线程，函数可以为已经写好的函数，也可以是匿名函数。
```go
func mainCon() {
	fmt.Println("main start")

	go func() {
		fmt.Println("goroutine")
	}()
  // 没有Sleep主程序会立马执行完退出。因为主进程在退出前 协程 才会有机会去执行（主进程在退出前不会等待全部 协程 执行完毕）
	// time.Sleep(1 * time.Second) // 让主线程休眠 1 秒钟，并发执行goroutine

	fmt.Println("main end")
}

// main start
// main end
```
- 为什么没有输出 goroutine
- 并发是不同的代码块交替执行，也就是交替可以做不同的事情。
- 并行是不同的代码块同时执行，也就是同时可以做不同的事情。
-  Go 语言的线程是并发机制，不是并行机制。main 函数是一个主线程，是因为主线程执行太快了，子线程还没来得及执行，所以看不到输出。
- 现在time.Sleep让主线程休眠 1 秒钟，再试试goroutine成功执行。

## routine调度
● 调度器：Go运行时包含一个调度器（scheduler），负责管理和调度goroutine。调度器会将goroutine分配到多个操作系统线程上执行，支持多核并行。
● M:N模型：Go使用M:N调度模型，即M个goroutine运行在N个操作系统线程上。调度器会根据负载情况动态调整goroutine的分配。
● 栈管理：Goroutine的栈是动态分配的，初始栈大小很小（通常为2KB），可以根据需要动态扩展和收缩

# channel 通道
并发编程的最大挑战源于数据的共享。 唯一的共享状态是通道
通道是协程之间用于传递数据的共享管道。换而言之，一个协程可以通过一个通道向另外一个协程传递数据。因此，在任意时间点，只有一个协程可以访问数据。
```go
// 声明不带缓冲的通道
ch1 := make(chan string)

// 声明带10个缓冲的通道
ch2 := make(chan string, 10)

// 声明只读通道
ch3 := make(<-chan string)

// 声明只写通道
ch4 := make(chan<- string)

cs := make(chan *os.File, 100)  // 指向文件的指针的缓冲信道

// 作为函数参数
func worker(c chan int) { ... }

// 写入 chan
ch1 := make(chan string, 10)
ch1 <- "a"

// 读取 chan
val, ok := <- ch1
// 或
val := <- ch1

// 关闭 chan
close(chan)
```
eg:
![file](../routine/forChan.go)

## chan共享内存
● 同步机制：Channel本身是**线程安全的**，通过内部的锁和条件变量实现同步。
● 阻塞和非阻塞：默认情况下，channel的发送和接收操作是**阻塞的**，直到有goroutine执行对应的接收或发送操作。可以通过select语句和default分支实现非阻塞操作。
● 缓冲机制：**Channel可以是有缓冲的或无缓冲的**。

### 无缓冲channel，阻塞（同步通信）
不带缓冲的通道，进和出都会阻塞。
● 发送操作：如果没有人接收数据，发送操作会阻塞。
● 接收操作：如果没有人发送数据，接收操作会阻塞。
![file](..\routine\testChanBLocked.go)

### 有缓冲channel，阻塞
带缓冲的通道，进一次长度 +1，出一次长度 -1。发送操作在缓冲区未满时不会阻塞，接收操作在缓冲区不为空时不会阻塞。
● 发送操作：如果缓冲区已满，发送操作会阻塞。
● 接收操作：如果缓冲区为空，接收操作会阻塞。
![file](..\routine\testChanBLocked.go)

### 非阻塞模式
在非阻塞模式下，发送和接收操作不会阻塞，而是立即返回。可以通过select语句和default分支实现非阻塞操作。
select 语句用于监听多个channel的发送或接收操作， 通常需要配合for循环。它会在其中一个channel准备好时执行对应的case分支然后退出。如果多个channel同时准备好，select 会随机选择一个执行
如果channel的发送操作无法立即完成，select语句会执行default分支。

![file](..\routine\testSelectDefault.go)
![file](..\routine\testSelectForloop.go)
- 在这个示例中，`select` 监听了两个通道 `ch1` 和 `ch2`，并**根据哪个通道先收到消息来执行相应的操作**。

eg2: https://learnku.com/docs/effective-go/2020/concurrent/6249
```go
var sem = make(chan int, MaxOutstanding)
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // 为该Go协程创建 req 的新实例。用相同的名字获得了该变量的一个新的版本， 以此来局部地刻意屏蔽循环变量
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

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
![image](../basics/concurrent.go)
```go
func (w *Worker) selectProcessTimeout(c chan int) {
	// 尝试发送数据到 c，等待最多 100ms，如果在这期间第一个 case 还不能执行，就走超时分支。
	for {
		select {
		case c <- rand.Int():
		case t := <-time.After(time.Millisecond * 100): // After返回一个t: chan time.Time，100ms 后会向这个通道写入当前时间。
			// case <-time.After(time.Millisecond * 100) // 不用t接收
			fmt.Println("timed out at", t)
		}
		time.Sleep(time.Millisecond * 50)
	}
}
```
> 第一个可用的通道被选择。
> 如果多个通道可用，随机选择一个。
> 如果没有通道可用，default 情况将被执行。
> 如果没有 default，select 将会阻塞。

## 最后
话虽如此，我仍然广泛使用 sync 和 sync / atomic 包中的各种同步原语。我觉得比较重要的是通过使用这两种方式比较舒适。我建议你首先关注通道，但是当你遇到一个需要短暂锁的简单示例时，请考虑使用互斥锁或读写互斥锁。


# 并发优化
## 固定数量协程
另一种管理资源的好方法就是启动固定数量的 handle Go 协程，一起从请求信道中读取数据。Go 协程的数量限制了同时调用 process 的数量
```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // 启动处理程序
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // 等待通知退出。
}
```

## 并行化最大利用cpu
- numCPU 常量值
```go
const numCPU = 4 // CPU 核心数

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // 缓冲区是可选的，但明显用上更好
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // 排空信道。
    for i := 0; i < numCPU; i++ {
        <-c    // 等待任务完成
    }
    // 一切完成
}
```
- 函数 [runtime.NumCPU](https://golang.org/pkg/runtime#NumCPU) 可以返回硬件 CPU 上的核心数量
> var numCPU = runtime.NumCPU()
-  runtime.GOMAXPROCS，设置当前最大可用的 CPU 数量，返回的是之前设置的最大可用的 CPU 数量。
- 注意不要混淆并发（concurrency）和并行（parallelism）的概念：并发：多个任务在重叠的时间段内执行，但不一定是同时进行。 而并行则是为了效率在多 CPU 上平行地进行计算。尽管 Go 的并发特性能够让某些问题更易构造成并行计算， 但 Go 仍然是种并发而非并行的语言
- 并发是关于“如何设计系统来处理多个事情”，而并行是关于“如何用更多资源同时做多个事情”。

## 漏桶缓存区限流设计
这里有个提取自 RPC 包的例子。 客户端 Go 协程从某些来源，可能是网络中循环接收数据。为避免分配和释放缓冲区， 它保存了一个空闲链表，使用一个带缓冲信道表示。若信道为空，就会分配新的缓冲区。 一旦消息缓冲区就绪，它将通过 serverChan 被发送到服务器。
```go
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // 若缓冲区可用就用它，不可用就分配个新的。
        select {
        case b = <-freeList:
            // 获取一个，不做别的。
        default:
            // 非空闲，因此分配一个新的。
            b = new(Buffer)
        }
        load(b)              // 从网络中读取下一条消息。
        serverChan <- b      // 发送至服务器。
    }
}
```
服务器从客户端循环接收每个消息，处理它们，并将缓冲区返回给空闲列表。
```go
func server() {
    for {
        b := <-serverChan    // 等待工作。
        process(b)
        // 若缓冲区有空间就重用它。
        select {
        case freeList <- b:
            // 将缓冲区放到空闲列表中，不做别的。
        default:
            // 空闲列表已满，保持就好。
        }
    }
}
```
客户端试图从 freeList 中获取缓冲区；若没有缓冲区可用， 它就将分配一个新的。服务器将 b 放回空闲列表 freeList 中直到列表已满，此时缓冲区将被丢弃，并被垃圾回收器回收。（select 语句中的 default 子句在没有条件符合时执行，这也就意味着 selects 永远不会被阻塞。）依靠带缓冲的信道和垃圾回收器的记录， 我们仅用短短几行代码就构建了一个限流漏桶。

## WaitGroup最佳实践!!
![file](..\routine\testWaitGroupPointer.go)

### 补充：
**`sync.WaitGroup`**: 管理多个 goroutine 的同步
   - 用于等待一组 goroutine 完成。
   - 示例：
     - `wg.Add(1)`：增加计数器。
     - `wg.Done()`：减少计数器。
     - `wg.Wait()`：等待计数器归零。

**goroutine通过定义的匿名函数 +() 立即执行**
```go
go func() {
    time.Sleep(time.Second * 1)
    ch1 <- "from ch1"
}()
```
go 关键字会启动一个新的goroutine，并在其中执行这个匿名函数
如果没有() 代码会阻塞在 <-ch1，因为没有任何数据发送到 ch1。
在 Go 中，匿名函数都是闭包：其实现在保证了**函数内引用变量的生命周期与函数的活动时间相同**。

---
### 输出结果
```
Goroutine 1 is done.
Received result from Goroutine 1.
Goroutine 2 is done.
Received result from Goroutine 2.
Goroutine 3 is done.
Received result from Goroutine 3.
All goroutines are finished.
```

## 避免goroutine泄漏
确保在不再需要goroutine时能够正确退出，避免goroutine泄漏。
可以通过context包或channel来通知goroutine退出

![file](..\routine\testContextWithTimeout.go)
CPU时间会在3个worker之间均匀分配，每个worker在每次Sleep期间会让出CPU，从而让其他worker有机会运行。
![file](..\routine\testContextWithValueCancelDeadline.go)

## context
`context`包提供了一种用于管理goroutine生命周期的机制

### `context` 的主要作用

1. **传递请求范围内的值**:
   - `context`可以用于在多个goroutine之间传递请求范围内的数据，例如用户身份、请求ID等。这些数据可以在整个请求处理链中共享，而不需要通过函数参数显式传递。

2. **取消信号**:
   - `context`可以携带一个取消信号，用于通知所有相关的goroutine停止它们的工作。这在处理超时或用户取消操作时非常有用。

3. **截止时间（Deadline）**:
   - `context`可以设置一个截止时间，当达到这个时间时，所有相关的goroutine都会收到取消信号。这对于限制操作的执行时间非常有用。

4. **超时控制**:
   - `context`可以设置一个超时时间，当超过这个时间时，所有相关的goroutine都会收到取消信号。这对于防止操作无限期地阻塞非常有用。

### `context` 的基本用法

1. **创建context**:
   - `context.Background()`: 创建一个空的context，通常作为根context使用。
   - `context.TODO()`: 创建一个空的context，通常用于尚未确定如何使用的场景。

2. **派生context**:
   - `context.WithCancel(parent Context)`: 创建一个带有取消功能的context。
   - `context.WithDeadline(parent Context, deadline time.Time)`: 创建一个带有截止时间的context。
   - `context.WithTimeout(parent Context, timeout time.Duration)`: 创建一个带有超时时间的context。
   - `context.WithValue(parent Context, key, val interface{})`: 创建一个带有键值对的context。

3. **使用context**:
   - `ctx.Done()`: 返回一个通道，当context被取消或超时时，该通道会被关闭。
   - `ctx.Err()`: 返回context被取消的原因（例如，`context.Canceled` 或 `context.DeadlineExceeded`）。
   - `ctx.Value(key)`: 从context中获取与键相关联的值。

### 示例代码
![file](..\routine\testContextWithTimeout.go)
![file](..\routine\testContextWithValueCancelDeadline.go)



### 代码解释

1. **`context.WithTimeout`**: 创建一个带有2秒超时的context。当超时时间到达时，context会被自动取消。

2. **`ctx.Done()`**: 在`worker`函数中，使用`select`语句监听`ctx.Done()`通道。当context被取消时，`ctx.Done()`通道会被关闭，`worker`函数会退出。

3. **`cancel()`**: 在`main`函数结束时，调用`cancel()`函数手动取消context，以确保所有相关的goroutine都能及时退出。

通过使用`context`，你可以更好地管理和控制goroutine的生命周期，避免资源泄漏和长时间阻塞。

## done 通道的用途
● 信号传递：chan struct{} 通常用于在 goroutine 之间传递信号，而不是传递实际的数据。例如，用来通知某个事件已经发生或某个任务已经完成。
● 关闭通道：关闭 done 通道（close(done)）可以用来广播信号，所有监听 done 通道的 goroutine 都会收到这个信号。

![file](..\routine\testChanDone.go)

## 会产生通道情况
### **1. `time.After`**
`time.After`返回一个通道，在指定的时间后会接收到一个值。

**示例：**
```go
select {
case <-time.After(2 * time.Second):
    fmt.Println("2 seconds elapsed")
}
```

---

### **2. `time.Tick`**
`time.Tick`返回一个通道，每隔指定的时间就会接收到一个值。

**示例：**
```go
ticker := time.Tick(1 * time.Second)
for range ticker {
    fmt.Println("Tick")
}
```

**注意**：`time.Tick`会创建一个永不关闭的通道，可能导致资源泄漏，通常建议使用`time.NewTicker`代替。

---

### **3. `os/signal.Notify`**
`os/signal.Notify`用于监听系统信号，并返回一个通道，当接收到指定信号时，通道会接收到该信号。

**示例：**
```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

select {
case sig := <-sigChan:
    fmt.Println("Received signal:", sig)
}
```

---

### **4. `context.WithTimeout` 或 `context.WithDeadline`**
`context.WithTimeout`和`context.WithDeadline`会返回一个带有超时或截止时间的`context`，其`Done()`通道会在超时或截止时间到达时关闭。

**示例：**
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-ctx.Done():
    fmt.Println("Context timeout")
}
```

---

### **5. `sync.WaitGroup` 的 `Wait` 方法（通过通道模拟）**
虽然`sync.WaitGroup`本身不返回通道，但可以通过通道模拟类似的行为。

**示例：**
```go
wg := sync.WaitGroup{}
done := make(chan struct{}) // struct{} 是一个空结构体，它不占用任何内存空间，通常用于传递信号而不是实际的数据。

wg.Add(1)
go func() {
    defer wg.Done()
    time.Sleep(2 * time.Second)
}()

go func() {
    wg.Wait()
    close(done)
}()

select {
case <-done:
    fmt.Println("All goroutines finished")
}
```

---

### **6. `chan` 的关闭**
手动关闭通道后，接收操作会立即返回通道类型的零值。

**示例：**
```go
ch := make(chan int)
go func() {
    time.Sleep(2 * time.Second)
    close(ch)
}()

select {
case val, ok := <-ch:
    if !ok {
        fmt.Println("Channel closed")
    } else {
        fmt.Println("Received:", val)
    }
}
```

---

### **7. `context.WithValue` 的派生**
虽然`context.WithValue`本身不返回通道，但可以结合`context.Done()`使用。
`context.Done()`返回一个通道（`<-chan struct{}`），当`context`被取消或超时时，该通道会被关闭。

**示例：**
```go
ctx := context.WithValue(context.Background(), "key", "value")
select {
case <-ctx.Done():
    fmt.Println("Context canceled")
}
```

---

### **8. `context.WithCancel` 的派生**
`context.WithCancel`返回一个可取消的`context`，其`Done()`通道在调用`cancel()`时关闭。

**示例：**
```go
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(2 * time.Second)
    cancel()
}()

select {
case <-ctx.Done():
    fmt.Println("Context canceled")
}
```

---

### **总结**
在Go中，以下方法或场景会产生通道：
1. `context.Done()`：用于监听`context`的取消或超时。
2. `time.After`：用于监听定时器。
3. `time.Tick`：用于监听周期性的定时器。
4. `os/signal.Notify`：用于监听系统信号。
5. 手动关闭通道：用于监听通道的关闭事件。
6. 使用`sync.WaitGroup`模拟通道：用于等待所有goroutine完成。
7. 其他场景：如`context.WithTimeout`、`context.WithDeadline`等。

根据具体的需求，可以选择合适的工具来监听事件或信号。

## defer
`defer` 是 Go 语言中的一个关键字，用于延迟执行某个函数或方法，它在声明时不会立刻去执行，而是在函数 return 后去执行的。`defer` 通常用于资源释放、日志记录、错误处理等场景。

### 注意事项
- `defer` 语句的执行顺序是“后进先出”（LIFO），即最后声明的 `defer` 语句会最先执行。
- `defer` 语句中的参数在声明时就已经被求值，而不是在执行时。

### 完整实例
![file](../basics/defer.go)


### 1. **资源释放**
在打开文件、数据库连接、网络连接等资源后，使用 `defer` 确保在函数结束时释放这些资源。

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // 确保文件在函数结束时关闭

    // 处理文件内容
    return nil
}
```

### 2. **解锁互斥锁**
在使用互斥锁（Mutex）时，使用 `defer` 确保在函数结束时解锁。

```go
var mu sync.Mutex

func safeIncrement() {
    mu.Lock()
    defer mu.Unlock() // 确保在函数结束时解锁

    // 安全的操作
}
```

### 3. **记录日志**
在函数执行过程中记录日志，使用 `defer` 确保日志在函数结束时被记录。

```go
func processData(data string) {
    defer log.Println("Processing complete") // 在函数结束时记录日志

    // 处理数据
}
```

### 4. **错误处理**
在函数执行过程中发生错误时，使用 `defer` 执行一些清理操作。

```go
func riskyOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()

    // 可能引发 panic 的操作
    return nil
}
```

### 5. **计时器**
使用 `defer` 和 `time` 包来记录函数的执行时间。

```go
func longRunningTask() {
    defer func(start time.Time) {
        log.Printf("Task took %v", time.Since(start))
    }(time.Now())

    // 执行耗时任务
}
```

### 6. **数据库事务**
在数据库事务中，使用 `defer` 确保在函数结束时提交或回滚事务。

```go
func updateDatabase(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback() // 确保在函数结束时回滚事务

    // 执行数据库操作

    return tx.Commit() // 如果成功则提交事务
}
```

### 7. **HTTP 请求处理**
在处理 HTTP 请求时，使用 `defer` 确保在函数结束时关闭响应体。

```go
func fetchData(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close() // 确保在函数结束时关闭响应体

    return io.ReadAll(resp.Body)
}
```

### 8. **自定义资源的清理**
如果你有自定义的资源需要清理，可以使用 `defer` 来确保它们在函数结束时被释放。

```go
func useCustomResource() {
    resource := acquireResource()
    defer releaseResource(resource) // 确保在函数结束时释放资源

    // 使用资源
}
```
