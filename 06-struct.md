# 结构体

## 定义

一组字段的集合就是结构体

示例：

```go
package main
import "fmt"
// 使用 type 去定义一个结构体，类型为struct
type MyVertex struct{
    X int
    Y int
}

func main()  {
    // 类似于实例化
    v := MyVertex{1,2}
    fmt.Println(v)  // 输出 {1,2}
    // 可以通过 . 来访问属性
    fmt.Println(v.X)    // 输出 1
    fmt.Println(v.Y)    // 输出 2
    // 也可以给属性赋值
    v.X = 4
    fmt.Println(v.X)    // 输出 4
    // 也可以通过结构体指针来访问
    v2 := MyVertex{3,4}
    // p是指向v2的指针
    p := &v2
    // 可以通过*来访问属性
    fmt.Println((*p).X)
    // 因为上面那种写法很麻烦，所以go也支持隐式的间接引用
    fmt.Println(p.X)
    // 声明结构体骚操作
    var(
        // 直接分配结构体
        v3 = MyVertex{1,2}
        // 通过name:value的方式分配结构体，未指定值的属性会被自动赋予默认值
        v4 = MyVertex{X:1}
        // 同上
        v5 = MyVertex{Y:1}
        // 这样也可以，两个都是默认值
        v6 = MyVertex{}
        // 也可以直接取指针
        v7 = &MyVertex{4,5}
    )
    fmt.Println(v3,v4,v5,v6,v7) // v7会输出 &{4,5}，表示这是个指向结构体的指针
}
```