# interface

学啥不好学接口
呵，在不声明接口的情况下，go支持智能指针，下面那个例子中使用v和&v效果一致  
但是使用接口的情况下，只能声明啥用啥  
为啥要这么设计呢？原因很简单：要支持面向对象中的*重载*。如果传参可以智能确定类型，那么重载功能就相当于废了
第二个例子中，显示了方法的重载

```go
package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat 实现了 Abser
    a = &v // a *Vertex 实现了 Abser

    // 下面一行，v 是一个 Vertex（而不是 *Vertex）
    // 所以没有实现 Abser。
    // a = v

    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

```go
package main

import (
    "fmt"
    "math"
)
// 定义接口
type I interface {
    M()
}
type T struct {
    S string
}
type F float64
// 这就是重载(实际上，go不支持重载，但是这种搞法和重载简直一样)
// 这个是结构体T的M方法
func (t *T) M() {
    fmt.Println("这是指针类型的调用！")
    fmt.Println(t.S)
}
//  这个是 类型F的M 方法
func (f F) M() {
    fmt.Println("这是值类型的调用！")
    fmt.Println(f)
}
func main() {
    var i I
    i = &T{"Hello"}
    describe(i)
    i.M()
    i = F(math.Pi)
    describe(i)
    i.M()
}
func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

## 底层值为nil的接口

即使接口底层值为nil，方法仍然会被nil接收者调用。
即使接口底层值为nil，接口自身是不为nil的，所以调用也没问题，但是在处理过程中要避免问题

```go
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}
// 在其他语言中，如果指针为空，会报空指针异常，但是go可以通过这种方式避开这种异常
func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I
    var t *T
    i = t
    describe(i)
    i.M()
    i = &T{"hello"}
    describe(i)
    i.M()
}
func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

## nil接口值

nil接口值表示接口自身都是nil，更别说底层了。  
这种接口既不保存值，也不保存类型，用了就报错

```go
package main

import "fmt"

type I interface {
    M()
}

func main() {
    var i I
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}

```

## 空接口

指定了*零个*方法的接口称为空接口  
这样： interface{}  
空接口可以指定任何值和类型  
因为空接口啥也没实现，而任何类型至少实现了*零个*方法  
所以可以随便搞  
看起来很没用？不不不，当要处理未知类型的值时，这个特性特别有用

```go
package main

import "fmt"

func main() {
    var i interface{}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
}

func describe(i interface{}) {
    // println就可接口任意类型为i的接口值，因为i啥都没实现
    fmt.Printf("(%v, %T)\n", i, i)
}
```

## Stringer

stringer 接口的定义如下

```go
type Stringer interface {
    String() string
}
```

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}
// 在调用println的时候，实际上会调用这个方法，然后打印值就会变化
func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a) // Arthur Dent (42 years)
    fmt.Println(z)  // Zaphod Beeblebrox (9001 years)
}
```

### 练习

通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。  
例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。

```go
package main

import "fmt"

type IPAddr [4]byte

func (ia IPAddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d",ia[0],ia[1],ia[2],ia[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}

```

## error

这也是个built-in 的接口  
定义如下：

```go
type error interface {
    Error() string
}
```

```go
package main

import (
    "fmt"
    "time"
    "strconv"
)

// 定义这个结构体，用来存error信息
type MyError struct {
    When time.Time
    What string
}

//  实现Error接口，接收者为 Myerror结构体，返回一个string
func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

// 定义一个函数，用来调用。如果返回的错误码不是nil，则表示运行失败
// 因为 Myerror实现了error，所以这里函数体返回MyError，定义是返回error
func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    // 不算是个正常调用，这个直接返回了错误
    if err := run(); err != nil {
        fmt.Println(err)
    }
    //  正常调用，第二个返回值是错误码，第一个返回值是处理结果
    i, err2 := strconv.Atoi("42")
    if err2 != nil {
        fmt.Printf("couldn't convert number: %v\n", err2)
        return
    }
    fmt.Println(i,err2)
}

```

### error练习

从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。
Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

```go
package main

import (
    "fmt"
    "math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f",float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0 , ErrNegativeSqrt(x)
    }
    return math.Sqrt(x), nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-100))
}

```

## reader

内置接口reader，有一大堆实现方式  
其中，io.Reader接口有个read方法

```go
func (T) Read(b []byte) (n int, err error)
```

示例：

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    // 实例化一个reader
    r := strings.NewReader("Hello, Reader!")
    // 定义一个切片，长度为1
    b := make([]byte, 1)
    // 死循环
    for {
        // 使用strings的read方法读取指定长度
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        // 当返回的错误码为io.EOF时，break掉
        if err == io.EOF {
            break
        }
    }
}

```

## image

图像
image 包定义了 Image 接口：

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

注意: Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。
color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。

```go
package main

import (
    "fmt"
    "image"
)

func main() {
    m := image.NewRGBA(image.Rect(0, 0, 100, 100))
    fmt.Println(m.Bounds())
    fmt.Println(m.At(100, 1).RGBA())
}
```
