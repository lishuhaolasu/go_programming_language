# goroutine

## 定义

goroutine是由 Go 运行时管理的轻量级线程

```go
// 创建一个协程
go func_name(param1, param2, ...)
// 创建一个协程
go func_name(param2, param3, ...)
// 执行所有协程(包括这条)
func_name(param4, param5, ...)
```

示例：

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    go say("1231231")
    say("hello")
}
```

## 
