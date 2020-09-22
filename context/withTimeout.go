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
     ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
     go reqTask(ctx, "workerA")
     go reqTask(ctx, "workerB")
     time.Sleep(3 * time.Second)
     fmt.Println("Before cancel")
     cancel()
     time.Sleep(3 * time.Second)
}
