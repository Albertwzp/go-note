## channel分为有缓冲或者无缓冲通道，按make是否有第二个可选参数区分

## <-为通道格式
var ch chan int = make(chan int, 4)
ch <- i 向通道写入
i <- ch 从通道读取

## 通道依据是否还有空间决定阻塞


