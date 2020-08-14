package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn,_ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square.", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
