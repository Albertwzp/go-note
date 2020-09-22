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
		     fmt.Println("Stop", name, ctx.Err())
		     return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
	 }
     }
}

func main() {
     ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
     go reqTask(ctx, "GoroutieA")
     go reqTask(ctx, "GoroutieB")

     time.Sleep(3 * time.Second)
     fmt.Println("Before cancel")
     cancel()
     time.Sleep(3 * time.Second)
}
