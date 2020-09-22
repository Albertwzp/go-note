package main

import (
	"fmt"
	"time"
	"context"
)

type Options struct{ Interval time.Duration }

func reqTask(ctx context.Context, name string) {
     for {
     	 select {
	 	case <-ctx.Done():
		     fmt.Println("Stop", name)
		     return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
	 }
     }
}

func main() {
     ctx, cancel := context.WithCancel(context.Background())
     vCtx := context.WithValue(ctx, "options", &Options{1})

     go reqTask(vCtx, "workerA")
     go reqTask(vCtx, "workerB")

     time.Sleep(3 * time.Second)
     cancel()
     time.Sleep(3 * time.Second)
}
