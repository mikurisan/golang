### [Goroutines](https://go.dev/tour/concurrency/1)

使用 `go func()` 启动一个 goroutine, 此时 func 的 params 的 evaluation 发生在当前 goroutine, 执行发生在一个新的 goroutine.

Goroutines 在同一个 address space 中执行, 因而需要保持对 shared memery 访问的 synchronized. 在包 `sync` 中提供了一些 primitives. 

### [Channels](https://go.dev/tour/concurrency/2)

Channels 是一种类型管道(typed conduit), 使用 `<-` 进行 values 的 send 和 receive, arrow 的 direction 表示了 data flows 的 direction.

默认情况下, send 或 receive 有一方没有 ready, 那么另一方就会 block 直到另一方 ready.

### [Buffered Channels](https://go.dev/tour/concurrency/3)

Buffered Channels 含有指定类型长度的缓冲区, 当缓冲区 full 时 sends 会 block, 当缓冲区 emptys 时 receive.

