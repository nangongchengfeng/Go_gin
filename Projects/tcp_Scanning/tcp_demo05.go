package main
 
import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)
 
var wg sync.WaitGroup
var mutex sync.Mutex
 
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
	startTime := time.Now()
	ports := []int{}
	timeout := time.Millisecond * 500
	for port := 1; port <= 65000; port++ {
		go func(p int) {
			opened := isOpen("www.baidu.com", p, timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				log.Printf("端口: %d 已经开通", p)
				mutex.Unlock()
			}
		}(port)
	}
	time.Since(startTime)
	cost := int(time.Since(startTime) / time.Second)
	fmt.Printf("opened ports: %v\n", ports)
 
	fmt.Printf("代码运行时长: %d S", cost)
}