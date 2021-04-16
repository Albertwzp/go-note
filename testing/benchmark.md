
-bench='Fib$' #匹配指定测试用例
-cpu=2,4 #改变 GOMAXPROCS,默认等于 CPU 核数
-benchtime=[5s|30x] #指定测试时间，默认1s，或者指定测试次数
-count=3 #指定测试轮数
-benchmem #显示度量内存分配的次数
b.ResetTimer() #重置定时器
b.StopTimer() & b.StartTimer() #暂停计时和开始计时
