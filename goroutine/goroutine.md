#### 1. goroutine 特性
 - 主协程不会等待子协程结束再退出，需要同步状态
 - 协程入口函数defer定义函数增加recover语句恢复协程内部异常,阻断传播到主协程导致程序崩溃
 - 设置线程数: runtime.GOMAXPROCS(n int) 
   - n <= 0 查询线程数
   - n > 0 设置线程数
 - 查询协程数: runtime.NumGrouptine()
