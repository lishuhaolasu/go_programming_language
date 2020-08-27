# 变量

## 关键字

哪个语言没有关键字。go也有关键字，一共25个  
关键字各有各的用处。  
和其他语言一样，变量名不能是关键字

|关键字列表|||||
|---|---|---|---|---|
| break | default| func| interface| select|  
| case | defer | go | map | struct |
| chan | else | goto | package | switch |
| const | fallthrough | if | range | type |
| continue | for | import | return | var |

## 预定义标识符

go还有36个预定义标识符，这些要么是数据类型，要么是内置函数  
也不要用这些标识符作为变量名

|预定义标识符||||||
|---|---|---|---|---|---|
| append | bool | byte | cap | close | complex |
| complex64 | complex128 | uint16 | copy | false | float32 |
| float64 | imag | int | int8 | int16 | uint32 |
| int32 | int64 | iota | len | make | new |
| nil | panic | uint64 | print | prinrln | real |
| recover | string | TRUE | uint | uint8 | uintprt |

## 数据类型类型

基本分类

> int  
>> int  int8  int16  int32  int64  
>> uint uint8 uint16 uint32 uint64  
>
> float32 float64  
> complex32 complex64  
> string  
> bool  

别名类型  
> rune (int32)  
> byte (uint8)  
---
补充说明  ：  
*零值(默认值)*
> int -> 0  
> string -> ""  
> bool -> false  

## 类型转换

```go
package main
import "fmt"
// 类型转换
var a = 1123
var b = float32(a)
// 如果实参不是float类型，则会有一个隐式的类型转换
func swap(a float) float {
    return a * 0.2
}

func main()  {
    // 这样也可以
    c := string(a)
    fmt.println(c)
}
```

## 声明

### 声明方式一：var(不初始化)

示例：

```go
package main
import "fmt"
// var可以一次性声明多个变量，在最后处添加变量类型
var a, b, c, d int  
func main()  {
    // 当然，声明一个变量也是OK的
    var i bool
    // 打印局部变量和全局变量
    fmt.Println(a,b,c,d,i)
}

```

### 声明方式二：var(类型推导)

示例：

```go
package main
import "fmt"
// 可以直接给变量赋值，这样就不用声明变量的类型了。编译器会自动获取变量的类型
var a = "1"
var b = 2
var c = false
// 可以同时赋值一大堆
var d,e,f = "1",2,true
// 如果想增加类型声明，则用下面这种方式声明
var g int = 3
// 甚至可以这样
var (
    h int = 123
    i = 3
    k = false
    l = "a"
)
func main()  {
    fmt.Println(a,b,c,d,e)
}
```

### 声明方式三：短声明(:=)

示例：  
> 注意：这种声明方式只能在函数内使用。函数外的部分，必须使用var，func等关键字声明

```go
package main
import "fmt"
func main()  {
    // 声明单个变量
    a := 1
    // 声明多个变量
    b,c := "false", true
    fmt.Println(a,b,c)
}
```

### 声明方式四：const

示例：
> 注意：常亮不可使用短声明

```go
package main
const a int = 1
const b bool = false
const c = "yes"
```

## 类型断言

go的类型断言用起来很方便的样子  
用法和map相似，但又些许不同  
map: value,exist := map_value[key]  
val_type :  value , ok := value.(type_name)  

```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"
    // 如果是string类型，就返回原值
    s := i.(string)
    fmt.Println(s)
    // 如果使用俩参数接收，则第二个参数为bool，意为是否为该类型
    s, ok := i.(string)
    fmt.Println(s, ok)
    // 用俩参数接收，即便类型不对，也不会报错
    f, ok := i.(float64)
    fmt.Println(f, ok)
    //  用一个参数接收就会报错,官网给的例子不明确，下面给个明确的
    k = i.(float64) // 报错(panic)
    fmt.Println(k)
}

```

## 类型选择

又是一个骚操作  
go允许使用value.(type)来判断类型(这里的type就是那个关键字type，而不是任意类型)  
然而，这么美好的功能只能用在switch语句中

```go
package main

import "fmt"

func do(i interface{}) {
    // 这种用法只能在switch语句中。即便类型不对，也不会报错，极其牛逼
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)              // Twice 21 is 42
    do("hello")         // "hello" is 5 bytes long
    do(true)            // I don't know about type bool!
}

```
