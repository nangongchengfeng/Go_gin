package main
 
import (
	"fmt"
	"net"
	"sync"
	"time"
)
 
func isOpen(host string, port int) bool {
	time.Sleep(time.Millisecond * 3)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err == nil {
		_ = conn.Close()
		return true
	}
 
	return false
}
 
func main() {
	ports := []int{}
	wg := &sync.WaitGroup{}
	for port := 1; port < 50000; port++ {
		wg.Add(1)
		port := port
		go func() {
			opened := isOpen("www.baidu.com", port)
			if opened {
				ports = append(ports, port)
			}
			wg.Done()
		}()
 
	}
	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}