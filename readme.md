# plog

`plog`是一个简单的Go语言日志库，它提供了基本的日志功能并允许用户根据需要开启和关闭日志功能。

## 功能

- `Printf`：类似于`log.Printf`，但可以通过`SetEnable`开启和关闭。
- `Println`：类似于`log.Println`，但可以通过`SetEnable`开启和关闭。
- `Fatalln`：类似于`log.Fatalln`，它在打印传入的参数后立即结束程序，并且`defer`语句**不会**被执行。
- `Panic`：类似于`log.Panic`，它停止当前函数的正常执行并开始运行在该函数中注册的所有`defer`语句。然后返回到该函数的调用者，并在调用者中做同样的事情，直到它回溯到当前goroutine的起点。在所有的`defer`语句被调用后，程序会以非零的状态码退出。
- `SetEnable`：设置是否开启日志。
- `GetEnable`：获取当前日志是否开启。

## 使用方法

```go
import "github.com/atmshang/plog"

func main() {
    plog.Println("This is a test log.")
    plog.SetEnable(false)
    plog.Println("This log will not be printed.")
    plog.SetEnable(true)
    plog.Fatalln("This is a fatal log.")
}
```

## 注意事项
尽管`Fatalln`和`Panic`在处理致命错误时非常有用，但在可能的情况下，你应该尽量避免使用它们。在Go中，推荐的错误处理方式是返回错误给调用者，并让调用者决定如何处理错误。