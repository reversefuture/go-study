# 同步,锁
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

实际上我们使用一个锁时也有可能发生死锁的问题，就是当我们忘记释放它时。
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

这在**读多写少**的场景中非常有用，能显著提高并发性能

### 📚 主要方法

`sync.RWMutex` 提供了四类方法：

| 方法 | 用途 | 说明 |
|------|------|------|
| `Lock()` | 获取写锁 | 写者调用，会阻塞所有其他读/写操作 |
| `Unlock()` | 释放写锁 | 必须与 `Lock()` 成对使用 |
| `RLock()` | 获取读锁 | 读者调用，多个读者可同时持有 |
| `RUnlock()` | 释放读锁 | 必须与 `RLock()` 成对使用 |

### ✅ 使用示例：并发读写共享数据
![file](../routine/rwmutex.go)

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
