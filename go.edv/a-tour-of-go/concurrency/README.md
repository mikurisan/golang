### [Goroutines](https://go.dev/tour/concurrency/1)

使用 `go func()` 启动一个 goroutine, 此时 func 的 params 的 evaluation 发生在当前 goroutine, 执行发生在一个新的 goroutine.

Goroutines 在同一个 address space 中执行, 因而需要保持对 shared memery 访问的 synchronized. 在包 `sync` 中提供了一些 primitives. 

### [Channels](https://go.dev/tour/concurrency/2)

Channels 是一种类型管道(typed conduit), 使用 `<-` 进行 values 的 send 和 receive, arrow 的 direction 表示了 data flows 的 direction.

默认情况下, send 或 receive 有一方没有 ready, 那么另一方就会 block 直到另一方 ready.

### [Buffered Channels](https://go.dev/tour/concurrency/3)

Buffered Channels 含有指定类型长度的缓冲区, 当缓冲区 full 时 sends 会 block, 当缓冲区 emptys 时 receive.

### [Range and Close](https://go.dev/tour/concurrency/4)

sender 可以使用 `close` 关闭一个 channel 表示不会继续发送 value 了. receiver 可以使用 `v, ok := <-ch` 用第 2 个参数测试 channel 是否已经 close.

循环 `for i := range c` 将会终止直到 channel 关闭.

值得注意的是, 最好仅有 sender 关闭 channel, 并且除非必须这么做, 否则这不是必要的.

### [Select](https://go.dev/tour/concurrency/5)

`select` 可以让 goroutine 等待多个 communicaion operations 直到有任一 `case` 满足时就执行, 如果有多个满足, 将会随机选择一个执行.

### [Default Selection](https://go.dev/tour/concurrency/6)

`default` 会在没有 case 满足的情况下执行, 小心使用以避免引起 block.

### [sync.Mutex](https://go.dev/tour/concurrency/9)

使用 `sync.Mutex` 使得一个 variable 在同一时刻只能被一个 goroutine 访问, 这也叫"互斥(mutual exclustion)", 提供这种机制的数据结构叫做 *mutex*.

`sync.Mutext` 的方法 `Lock` 和 `Unlock` 提供了 mutual exclustion 的机制.
