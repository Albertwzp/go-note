package main

import (
       "fmt"
       "time"
       "context"
)

func reqTask(ctx context.Context, name string) {
     for {
     	 select {
	 	case <-ctx.Done():
		     fmt.Println("stop", name)
		     return
		default:
		     fmt.Println(name, "send request")
		     time.Sleep(1 * time.Second)
	 }
     }
}
func main() {
     ctx, cancel := context.WithCancel(context.Background())
     go reqTask(ctx, "worker")
     time.Sleep(3 * time.Second)
     cancel()
     time.Sleep(3 * time.Second)
}
