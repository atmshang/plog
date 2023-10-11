package plog

import (
	"log"
)

// 不影响fatal和panic
var enable = true

func SetEnable(e bool) {
	enable = e
}

func GetEnable() bool {
	return enable
}

func Printf(format string, v ...any) {
	if enable {
		log.Printf(format, v)
	}
}

func Println(v ...any) {
	if enable {
		log.Println(v)
	}
}

func Fatalln(v ...any) {
	log.Fatalln(v)
}

func Panic(v ...any) {
	log.Panic(v)
}

/**
在Go语言中，`Fatalln`和`Panic`都被用来在遇到无法处理的错误时结束程序，但它们在处理方式和用途上有一些区别：

- `log.Fatalln`: 这个函数在打印传入的参数（通过调用`Println`）后调用`os.Exit(1)`来立即结束程序。这意味着`defer`语句**不会**被执行。

- `panic`: 这个函数停止当前函数的正常执行，开始运行在该函数中注册的所有`defer`语句。然后返回到该函数的调用者，并在调用者中做同样的事情，直到它回溯到当前goroutine的起点。在回溯过程中，程序将打印出一个错误消息和堆栈跟踪。在所有的`defer`语句被调用后，程序会以非零的状态码退出。

因此，你应该在以下情况中使用它们：

- 如果你想要立即结束程序，并且不需要清理任何资源（比如关闭文件或数据库连接），你可以使用`log.Fatalln`。

- 如果你需要在程序结束前执行一些清理工作（比如运行`defer`语句），或者你想提供一个堆栈跟踪来帮助调试，你应该使用`panic`。

请注意，尽管这两个函数在处理致命错误时非常有用，但在可能的情况下，你应该尽量避免使用它们。在Go中，推荐的错误处理方式是返回错误给调用者，并让调用者决定如何处理错误。
*/
