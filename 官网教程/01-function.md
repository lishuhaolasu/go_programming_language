# 函数

## 基本入门

声明：func  
参数：任意，但是需要声明形参类型和返回类型  
示例:  

```go
package test
// 每个形参都声明类型
func add(a int, b int) int {
    return a + b
}
// 形参相同类型，可以省略前面的定义
func add(a, b int) int {
    return a + b
}
// 返回多个参数
func swap(a,b int) (int,int){
    return b,a
}
// 骚操作了，可以预先声明返回值，在函数中可以直接返回对应的值，这颗糖有点好吃
func split(a,b int) (x,y int){
    x = a * 5/12
    y = b - a
    return
}
```

## 高级功能


### 第一类型值

函数也是值，如同Python、JAVA、lua等语言一样，函数也是第一类型值，可以像值一样传递

```go
package main
import "fmt"
// 这个map用来映射函数名与函数,映射函数时，函数本身也是一个type，定义不能错，否则会报错
var calcMap map[string]func(x,y float64)float64
// 这个结构体存的也是函数名和函数
type calcStruct struct {
    add func(x,y float64)float64
    sub func(x,y float64)float64
    multiply func(x,y float64)float64
    divide func(x,y float64)float64
}
// 这个
func add(x ,y float64) float64 {
    ret := x + y
    return ret
}   
func sub(x,y float64) float64 {
    ret := x - y
    return ret
}
func multiply(x,y float64) float64 {
    ret := x * y
    return ret
}
func divide(x,y float64) float64 {
    ret := x/y 
    return ret
}
var (
    a = add
    s = sub
    m = multiply
    d = divide
)
func main() {
    fmt.Println(a(1,2))
    fmt.Println(s(1,2))
    fmt.Println(m(1,2))
    fmt.Println(d(1,2))
    calcMap = make(map[string]func(x,y float64)float64)
    calcMap["add"] = add
    calcMap["sub"] = sub
    calcMap["multiply"] = multiply
    calcMap["divide"] = divide
    fmt.Println(calcMap["add"](1,2))
    fmt.Println(calcMap["sub"](1,2))
    fmt.Println(calcMap["multiply"](1,2))
    fmt.Println(calcMap["divide"](1,2))
}

```

### 闭包

go 当然也有闭包  
闭包是一个函数值，包含函数及函数所引用的外部变量，这个外部变量和函数是绑定在一起的  
废话不多说，上例子

```go
package main
import "fmt"

func count() func(int) int {
    sum := 0
    return func (i int) int {
        sum += i
        return sum
    }   
}
func main() {
    pos1,post := count() ,count()
    var ret1,ret2 int
    for i:=1 ;i<=100 ;i++ {
        ret1 = pos1(i)
        ret2 = post(2*i)
    }
    fmt.Println(ret1,ret2)
}

```

#### 斐波那契数列

这是个实战

```go
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
    now := 0
    last := -1
    return func() int {
        // 这几个if-else是为了给前3项赋值，如果不要前三项，则可以直接跳过
        if now == 0 && last == -1 {
            now = 1
            return 0
        } else if last == -1 && now == 1 {
            last = 0
            now = 1
            return 1
        } else if last == 0 && now == 1{
            last = 1
            now = 1
            return 1
        }

        now , last = now+last,now
        return now
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
```

### 方法

go 莫得类，但是可以为结构体类型定义方法  

#### 值接受者

go允许在定义函数的声明过程中指定一个接收者，这个接受者可以像参数一样使用  
值接收者就是其中一种，它的接受者是一个值，方便，但是功能有限，可以先做了解 
上例子，例子说明一切

```go
package main
import "fmt"
import "math"

// 先定义个结构体
type MyStruct struct{
    X , Y float64
}

// 定义个函数，前面括号内的内容相当于Python中的self，将自动接收s作为接收者
func (s MyStruct) Abs() float64 {
    return math.Sqrt(s.X * s.X + s.Y * s.Y)
    }
// 写成正常函数就是这样
func Abs2(v MyStruct) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main()  {
    s := MyStruct{3,4}
    // 可以作为方法调用了
    fmt.Println(s.Abs())
    // 注意，Abs这个函数并不是结构体的一部分，结构体并不包含这个属性
    fmt.Println(s)      // 输出：{3 4}
    // 这么调用是等效的
    fmt.Println(Abs2(s))
}
```

结构体是一种类型。既然结构体可以声明方法，其他的结构当然也可以  
虽然定义方法很方便，但是有go有自己限制  
- 类型定义和方法定义要在同一个包内  
- 不能为內建类型定义方法  

虽然第二点看起来有点坑，但是可以绕过这种限制。  
上例子

```go
package main
import "fmt"
import "math"
// 声明一个类型，这个类型就是float64
type newFloat float64

func (f newFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
func main() {
    f := newFloat(math.SqrtE)
    fmt.Println(f.Abs())
}
```

### 指针接受者

说完值接受者，当然要说指针接受者了  
上例子，有对比才有伤害

```go
package main
import "fmt"
import "math"
type MyStruct struct{
    X , Y float64
}
func (m MyStruct) Abs() float64{
    return math.Sqrt(m.X*m.X + m.Y*m.Y)
}
func (m MyStruct) scale1(f float64) {
    m.X = m.X * f
    m.Y = m.Y * f
}
func (m *MyStruct) scale2(f float64) {
    m.X = m.X * f
    m.Y = m.Y * f
}
func main() {
    m := MyStruct{3,4}
    m.scale1(10)
    fmt.Println(m.Abs())    // 输出 5
    m.scale2(10)
    fmt.Println(m.Abs())        // 输出50
}
```

可以发现，上面的例子执行后，分别输出了5和50     code
产生的原因是：如果接收者是值参数，则在调用时，go会创建一个结构体的副本对这个结构体进行操作  
用Python来解释就是：scale1在调用时重新实例化了一个匿名实例，而调用scale2的时候是对self进行的操作  

```go
package main
import "fmt"
type Vertex struct {
    X, Y float64
}
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
func main() {
    v := Vertex{3, 4}
    v.Scale(2)
    ScaleFunc(&v, 10)
    // ScaleFunc(v, 10) 报错
    p := &Vertex{4, 3}
    p.Scale(3)
    ScaleFunc(p, 8)
    fmt.Println(v, p)
}
```

这是官方给的例子，有神坑。  
如果以普通函数定义(ScaleFunc)，这时参数v只能是指针，而不能是值  
但如果是以接受者形式定义(Scale)，则即可以是值，也可以是指针。  
Scale在以值形式调用的时候(v.Scale)实际上是被go的编译器解释为了(&v).Scale  
同样的事情发生在接收者为值参数的情况下。  
如果接收者为值(v Vertex)，那么无论v是指针还是值，都可以正常使用，值会自动解释成(*v).Scale，而以函数形式定义就会报错

---

## 总结

建议使用指针接收者。原因如下  

- 无论传递的是值还是指针，程序都能自动处理
- 方法能够修改指向的值
- 由于使用值类型会导致底层重新复制该结构，所以在大型结构体中，指针接收者的效率更高

当然，选择值接收者并没有错，确保自己使用的过程中能管理好即可。  
注意：

- *类型方法都应该有接收者*
- *不要在一个类型中混用两种接收者*
