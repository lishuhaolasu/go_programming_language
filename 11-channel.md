# channel

## 定义

信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

```go
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
```

（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

```go
// make 了一个传递int类型的信道
ch := make(chan int)
```

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

```go
package main
import "fmt"
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
        fmt.Println(v)
    }
    c <- sum // 将和送入 c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    // 俩goroutine，你以为是一起执行？不，因为使用了信道，而且信道是阻塞的。所以，这里是分先后执行的。上面的print会打印 -9, 4, 0,7, 2, 8
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // 从 c 中接收

    fmt.Println(x, y, x+y)
}
```

## 缓冲区

信道可以带缓冲区。在定义信道的时候，增加第二个参数即可。  
第二个参数代表缓冲区大小，当填满这个缓冲区时，向信道发送信息才会阻塞

示例：

```go
package main

import "fmt"

func main() {
    // 信道缓冲区大小为2
    ch := make(chan int, 2)
    // 发送俩数据，没问题
    ch <- 1
    ch <- 2
    // 假如发送了仨数据，因为前俩数据无法取出，又加了一个数据进去，就发生了无法解决的阻塞，这就是死锁
    // ch <- 3
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

```

## range 和 close

可以使用close来关闭一个channel，虽然这个操作并不常用。  
和range 一起讲的原因就是：close经常和range一起出现。  
range操作channel的时候，会不停从channel中取值，直到channel被关闭。

注意，只有发送者才能关闭信道，接收者不可以。此外，从已关闭的信道中取值，会引发恐慌异常

```go
// 从信道取值，第一个是值，第二个是信道是否开启。如果信道关闭了，则ok的值为false
v , ok := <- c
```

```go
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    // 循环完成后，关闭信道
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    // 不断从信道取值
    for i := range c {
        fmt.Println(i)
    }
}
```

## select

不是sql里那个select，这个select是给goroutine和channel用的。  
用途是让goroutine选择一个没有阻塞的channel继续运行，如果都阻塞了，就等着，如果都不阻塞，就随机执行。

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        //  从下面的channel选择执行
        select {
        // 看看c 阻没阻塞，没阻塞就执行这个
        case c <- x:
            x, y = y, x+y
        //  看看quit阻没阻塞，没阻塞就执行这个
        case <-quit:
            fmt.Println("quit")
            return
        // 也有default啦
        default:
            fmt.Println("没准备好")
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    // 搞个协程
    go func() {
        for i := 0; i < 10; i++ {
            // 每次循环从channel里取个值
            fmt.Println(<-c)
        }
        // 循环完了给quit channel 传个值
        quit <- 0
    }()
    // 调用生产者
    fibonacci(c, quit)
}
```
