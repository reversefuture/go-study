# Go 并发编程
Go 语言中的多线程编程主要通过 **goroutines** 和 **channels** 来实现。Go 语言没有直接的“线程”概念，而是使用 goroutines 来并发执行任务。
Go语言的设计哲学是“不要通过共享内存来通信，而要通过通信来共享内存”

## routine调度
● 调度器：Go运行时包含一个调度器（scheduler），负责管理和调度goroutine。调度器会将goroutine分配到多个操作系统线程上执行，支持多核并行。
● M:N模型：Go使用M:N调度模型，即M个goroutine运行在N个操作系统线程上。调度器会根据负载情况动态调整goroutine的分配。
● 栈管理：Goroutine的栈是动态分配的，初始栈大小很小（通常为2KB），可以根据需要动态扩展和收缩

## chan共享内存
● 同步机制：Channel本身是**线程安全的**，通过内部的锁和条件变量实现同步。
● 阻塞和非阻塞：默认情况下，channel的发送和接收操作是**阻塞的**，直到有goroutine执行对应的接收或发送操作。可以通过select语句和default分支实现非阻塞操作。
● 缓冲机制：**Channel可以是有缓冲的或无缓冲的**。

### 无缓冲channel，阻塞
无缓冲channel的特点是：发送和接收操作必须同时进行，否则会阻塞。
● 发送操作：如果没有人接收数据，发送操作会阻塞。
● 接收操作：如果没有人发送数据，接收操作会阻塞。
![file](..\routine\testChanBLocked.go)

### 有缓冲channel，阻塞
有缓冲channel的特点是：发送操作在缓冲区未满时不会阻塞，接收操作在缓冲区不为空时不会阻塞。
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


## WaitGroup最佳实践
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
`defer` 是 Go 语言中的一个关键字，用于延迟执行某个函数或方法，直到包含它的函数返回为止。`defer` 通常用于资源释放、日志记录、错误处理等场景。

### 注意事项
- `defer` 语句的执行顺序是“后进先出”（LIFO），即最后声明的 `defer` 语句会最先执行。
- `defer` 语句中的参数在声明时就已经被求值，而不是在执行时。

### 完整实例


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
