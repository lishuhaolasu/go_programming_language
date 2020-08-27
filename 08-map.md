# map

## 简介
映射  
映射是指一种将键映射到值的数据结构,类似于lua中的table,Python中的字典  
其构造为 map_var_name map[key_type]value_type
映射的零值为nil，nil映射没有键也不能添加键

## 声明与构造

直接上例子

```go
package main
import "fmt"

type MyStruct struct{
    x, y float64
}
// 定义一个变量，类型是映射
var m map[string]MyStruct
func main()  {
    // 生成一个map实体
    m = make(map[string]MyStruct)
    // 给这个实体添加键值对MyStruct
    m["1"] = MyStruct{123.1,321.1}
    fmt.Println(m)
    // 直接来一套
    var m2 = map[int]MyStruct{
        1:MyStruct{1.1,2.2},
        2:MyStruct{3.1,4.2},
        3:{5.1,1.2},
        5:{9.1,3.2},
    }
    fmt.Println(m2)
    // 如果值只是个类型名，可以省略声明。如上面有部分代码省略了MyStruct的声明，下面的代码省略了int的声明
    var m3 = map[int]int{
        1:2,
        2:3,
        3:4,
    }
    fmt.Println(m3)
}
```

## 修改

继续上例子

```go
package main
import "fmt"
func main()  {
    var m map[string] int
    // 插入元素
    m["one"] = 1
    m["two"] = 3
    // 修改元素
    m["two"] = 2
    // 获取元素
    ele := m["two"]
    fmt.Println(ele)
    fmt.Println(m)
    // 删除元素
    delete(m,"two")
    // 判断元素是否存在。第一个值是元素对应的值（没有返回对应类型的默认值），第二个值是元素是否存在。
    value, exist := m["one"]
    fmt.Println(value,exist) // 1 true
    value, exist = m["two"]
    fmt.Println(value,exist) // 2 false 
}
```

### 练习

又到了激动人心的练习时刻  
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。  
函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
那个golang.org导入进去的是个测试程序，可以在这里[做题](https://tour.go-zh.org/moretypes/23)

```go
package main
import "fmt"
import "strings"
import "golang.org/x/tour/wc"

type WordCountStruct struct{
        word string 
        count int
    } 

var resMap map[string]int

func WordCount(s string) map[string]int {
    wordlist := strings.Fields(s)
    resMap = make(map[string]int)
    for _,word := range wordlist {
        if _,exist := resMap[word]; exist {
            resMap[word] += 1
        } else {
            resMap[word] = 1
        }   
    }
    fmt.Println(resMap)
    return resMap
}

func main() {
    wc.Test(WordCount)
}

```
