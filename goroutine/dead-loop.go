package main

import "fmt"
import "time"
import "runtime"

func main() {
    fmt.Println(runtime.GOMAXPROCS(0))
//    runtime.GOMAXPROCS(25)
    fmt.Println("run in main goroutine")
    n := 24
    for i:=0; i<n; i++ {
        go func() {
            fmt.Println("dead loop goroutine start")
            for {}  // 死循环
        }()
    }
    for {
        time.Sleep(time.Second)
        fmt.Println("main goroutine running")
    }
}
