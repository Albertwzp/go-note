#
### 零值结构体和nil结构体的区别
```go
type Circle struct {
    x int
    y int
    Radius int
}
var c *Circle = new(Circle)
var c *Circle = nil
```
### slice string & dict
```go
// 24bytes
type slice struct {
  array unsafe.Pointer  // 底层数组的地址
  len int // 长度
  cap int // 容量
}
// 16bytes
type string struct {
  array unsafe.Pointer // 底层数组的地址
  len int
}

type hmap struct {
  count int
  ...
  buckets unsafe.Pointer  // hash桶地址
  ...
}
```
