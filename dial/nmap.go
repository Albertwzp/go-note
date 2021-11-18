package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func main() {
	hostname := flag.String("hostname", "", "host to scan")
	sport := flag.Int("start-port", 1, "begin to scan")
	eport := flag.Int("end-port", 65535, "end scan")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	flag.Parse()

	ports := []int{}
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	for port := *sport; port <= *eport; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mx.Lock()
				ports = append(ports, p)
				mx.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Printf("Opened ports: %v\n", ports)
}
