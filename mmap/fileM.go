package main

import (
       "fmt"
       "os"
)

func main() {
     f, _ := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0644)
     _, _ = f.WriteAt([]byte("abcdefg"), 0)
     buff := make([]byte, 3)
     _, _ = f.ReadAt(buff, 4)
     _ = f.Close()
     fmt.Println(string(buff))
}
