# 数组

[中文官方文档](https://blog.go-zh.org/go-slices-usage-and-internals)

#### 定义

一个包含n个同种类型t的有序集合
数组类型是值。与c语言不同，go的数组变量并不是表示指向第一个元素的指针，而是表示整个数组。
虽然数组看起来比较灵活，但是在go应用不多，更多的是切片

示例：

```go
package main
import "fmt"
func main()  {
    // 声明时不进行初始化
    var arr01 [2]string
    arr01[0] = "1"
    arr02[1] = "2"
    // 声明时进行初始化
    arr02 := [2]int{1,2}
    // 上面的声明也可以写作 arr02 := [...]int{1,2} 让go自行判断数组长度
    fmt.Println(arr01,arr02)
}
```

---

# 切片

#### 定义

就，数组切片，数组的一部分，左闭右开
但，切片属于纯引用，自身不储存数据。如果被切的那个数组变了，切片也会变。

示例：

```go
package main
import "fmt"
func main()  {
    arr = [10] int {1,2,3,4,5,6,7,8,9,10}
    // 完整声明进行切片
    var sli []int = arr[4:6]
    // 短声明进行切片
    sli2 := arr[3:5]
    fmt.Println(sli,sli2)
}
```

示例2：

```go
package main
import "fmt"
func main()  {
    // 这会创建一个数组
    arr := [10] int {}
    // 这会创建一个数组，然后创建一个指向整个数组的切片
    sli := [] int {0,0,0,0,0}
    // 下面这种方式：声明了一个包含结构体的数组，结构体的类型一并声明在一起
    sli2 := []struct{
    a int
    b bool
    }{
        {1,true},
        {2,false},
        {3,true},
        {4,false},
        {5,true},
    }
    // 可以改写成这种形式,就比较容易看懂了
    type s struct {
    a int
    b bool
    }
    sli3 := [] s {
        {1,true},
        {2,false},
        {3,true},
        {4,false},
        {5,true},
    }
    fmt.Println(arr,sli,sli2,sli3)
}
```

#### 默认行为

切片时，未给定的值会采用上下界自动判断

示例：

```go
package main
import "fmt"
func main()  {
    arr := []int {1,2,3,4,5,6,7,8,9,0}
    // 下面的输出完全一致
    fmt.Println(arr)
    fmt.Println(arr[:])
    fmt.Println(arr[0:])
    fmt.Println(arr[:10])
    fmt.Println(arr[0:10])
}
```

#### 容量与长度

长度：可以通过len()获取，表示切片中包含的元素数量。
容量：可以通过cap()获取，表示当前切片最大容纳的元素数量。
> 重点来了  
> 容量表示的是：*底层数组长度* 减去 *左指针走过的长度*  
> 如：底层数组为{1,2,3,4,5,6}，len=6 cap=6  
>> 第一次切片为[0:5]，左下标没有移动，右下标移动了5，len=5 cap=6 slice=[1 2 3 4 5 ]  
>> 第二次切片为[1:]，左下标移动了1，右下标没有移动，len=4 cap=5 slice=[2 3 4 5]  
>> 第三次切片为[1:]，左下标移动了1，右下标没有移动，len=4 cap=5 slice=[3 4 5]  
>> 第四次切片为[1:]，左下标移动了1，右下标没有移动，len=4 cap=5 slice=[4 5]  
>> 第五次切片为[1:]，左下标移动了1，右下标没有移动，len=4 cap=5 slice=[5]  
>
> 每次都在原有的切片基础上从左往右切，容量会一直变小，长度也会变小，从右往左切，只有长度变小。  
> (大概是因为实现问题？去掉头部的空间复杂度要大于去掉尾部的空间复杂度。)

#### 空切片

切片的零值是nil，虽然打出来是一对空括号，但是在判断中，空切片==nil

```go
package main
import "fmt"
func main()  { 
    var s []int
    fmt.Println(s,len(s),cap(s))
    if s == nil {
        fmt.Println("this slice is nil!")
    }
}
```

#### make

内建函数 make 可以创建一个数slice,map,channel

```go
package main
import "fmt"
func main()  {
    // 参数1是切片类型，参数2是长度，可选参数3为长度。创建后，会将各元素设置为默认值
    arr := make([]int,3)    // len=cap=3
    arr2 := make([]int,3,4) // len=3 ,cap=4
    fmt.Println(arr,arr2)
}
```

#### append

追加元素。如果底层数组的容量不足以支持追加后的切片大小，go会自动分配一个更大的数组给切片

```go
package main
import "fmt"
func main()  {
    var sli []int
    sli2 := []int{1,2,3,4}
    sli = append(sli , 1)
    sli = append(sli , 2)
    // 这个写法比较难受。解释：如果要增加的元素是数组或切片，则需要在后面加上...进行解包，然后逐一添加
    sli = append(sli , sli2...)
    fmt.Println(sli)
}
```

#### range

类似于Python中的enumerate，迭代一个数组或者切片，每次迭代返回index和value
其中 index一般来讲不需要，可以用下划线代替(没错，Python也这么干)
如果只需要只需要下标，则可以只接受range返回的一个参数，连下划线都可以省了

```go
package main
import "fmt"
func main()  {
    sli := []int {9,8,7,6,5,4,3,2,1,}
    for index, value := range sli {
        fmt.Println(index, value)
    }
}
```

练习题：  
实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。  
当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。  
图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。  
提示：需要使用循环来分配 [][]uint8 中的每个 []uint8；
请使用 uint8(intValue)在类型之间转换；你可能会用到 math 包中的函数。）

```go
package main
import (
"golang.org/x/tour/pic"
)
// 输入长度和宽度，返回一个二维切片
func Pic(dx, dy int) [][]uint8 {
    // 定义一下二维切片
    var res [][]uint8
    for i:=1 ; i<=dx ;i++ {
        // 定义第一纬度的元素，这个元素是一个uint8的切片
        var resX []uint8
        for j:=1 ;j<=dy ;j++ {
            // 把元素加到子切片中
            resX = append(resX,uint8(i*j))
        }
        // 把切片添加到二维切片中
        res = append(res,resX)
        
    }
    return res
}

func main() {
    // 显示图片
    pic.Show(Pic)
}
```
