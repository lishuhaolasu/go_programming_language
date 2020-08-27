# 循环

golang 中只有一种循环：for

## for 循环

```go
package main
import "fmt"
func main(){
    sum := 0
    // 循环语句和c类似，但是无需加括号。循环体必须要用大括号包裹
    // 初始化语句和后处理语句也可以省略，程序会在判断语句为false时跳出
    for i:=1 ;i<10;i++ {
        sum += i
        fmt.Println(i)
    }
    fmt.Println(sum)
    sum2 := 1
    for ;sum<10;{
        sum2 += sum2
    }
}
```

```go
package main
import "fmt"
func main(){
    sum := 0
    // 去掉前置后置和分号，就形成了while循环
    for sum < 1000 {
        sum += 1
    }
    // 胆子再大点，循环条件和循环体都可以省略，变成真正的死循环。虽然看起来诡异，但是合法
    for {
    }
}
```

