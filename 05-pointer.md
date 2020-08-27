# 指针

## 定义

> 指针保存了值的内存地址  
> 直接打印指针时，会输出当前的保存的地址，而不是值  
> 如果想取指针的底层值，则需要使用*进行取值
> 指针的默认值是nil

## 声明

```go
package main
import  "fmt"
func main(){
    a, b := 1,2
    // 定义一个指向a的指针。如果打印上面的指针，则会输出地址（并不是所有的打印都是地址）
    i := &a
    fmt.Println(i)
    // 通过*取出对应的值
    fmt.Println(*i)
    // 通过指针进行赋值
    *i = 21
    fmt.Println(i)
    fmt.Println(*i)
    i = &b
    fmt.Println(i)
    fmt.Println(*i)
}
```
