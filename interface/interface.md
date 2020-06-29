## 用接口来模拟多态
> 使用这种方式模拟多态本质上是通过组合属性变量（Name）和接口变量（Fruitable）来做到的，属性变量是对象的数据，而接口变量是对象的功能，将它们组合到一块就形成了一个完整的多态性的结> 使用这种方式模拟多态本质上是通过组合属性变量（Name）和接口变量（Fruitable）来做到的，属性变量是对象的数据，而接口变量是对象的功能，将它们组合到一块就形成了一个完整的多态性的结构体体
```go
package main

import "fmt"

type Fruitable interface {
    eat()
}

type Fruit struct {
    Name string  // 属性变量
    Fruitable  // 匿名内嵌接口变量
}

func (f Fruit) want() {
    fmt.Printf("I like ")
    f.eat() // 外结构体会自动继承匿名内嵌变量的方法
}

type Apple struct {}

func (a Apple) eat() {
    fmt.Println("eating apple")
}

type Banana struct {}

func (b Banana) eat() {
    fmt.Println("eating banana")
}

func main() {
    var f1 = Fruit{"Apple", Apple{}}
    var f2 = Fruit{"Banana", Banana{}}
    f1.want()
    f2.want()
}
```
## 接口组合继承
```go
type Smellable interface {
  smell()
}
type Eatable interface {
  eat()
}

// 如下两种形式等效
type Fruitable interface {
  Smellable
  Eatable
}
type Fruitable interface {
  smell()
  eat()
}
```
## 接口变量赋值
**浅拷贝: 数据内存的复制,与之相反的则是指针变量的拷贝,所指内存共享.**

  > 变量赋值: 浅拷贝
  > 切片赋值: 拷贝切片头
  > 字符串赋值: 拷贝字符串头部

