# 流程控制

## if 判断

示例：

```go
package main
func main(){
    sum := 1
    for {
        // 无需小括号，但是大括号必须有
        if sum > 1000 {
            break
        }
        // if 也可以加上前置处理步骤，加个分号就行
        if x:=sum/2;x>500{
            break
        }
        sum += sum
        // go当然也和其他语言一样支持else
        //if {} else if {} else {}
    }
}
```

## switch case

示例

```go
package main
import "runtime"
import "fmt"
func main(){
    // 和C的switch case极其相似
    // 区别有：
    // 1：可以在switch处做一个预处理
    // 2：case不必是常量，也不必是整形
    // 3：case不需要显式break，go自动处理掉了。也就是说，不支持执行连续的case。断路
    // 4：switch也支持不写条件。这样可以写出优美的if-elseif
    switch os:= runtime.GOOS; os {
        case "darwin":
            fmt.Println("go runs on darwin")
        case "linux":
            fmt.Println("go runs on linux")
        default:
            fmt.Println("go runs on other")

    }
    a:=4
    switch {
        case 1:
            fmt.Println("number is 1")
        case 2:
            fmt.Println("number is 2")
        case 3:
            fmt.Println("number is 3")
        default:
            fmt.Printf("number is %d",a)
    }
}
```

## defer

说明：这玩意可以让语句在函数退出之后再执行  
defer会将所有未执行的命令压入一个栈内，命令会以后进先出的方式执行

示例：

```go
package main
import "fmt"
// 执行后，会输出world hello
func main(){
    defer fmt.Println("hello")
    fmf.Println("world")
}

func main2(){
    for i:=1 ;i <10;i++ {
        defer fmt.Println(i)
    }
}
```
