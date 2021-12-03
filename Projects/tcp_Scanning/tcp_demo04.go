package main
 
import (
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
    ports := []int{}
 
    wg := &sync.WaitGroup{}
    timeout := time.Millisecond * 200
    for port := 1; port < 100; port++ {
        wg.Add(1)
        go func(p int) {
            opened := isOpen("www.baidu.com", p, timeout)
            if opened {
                ports = append(ports, p)
            }
            wg.Done()
        }(port)
    }
 
    wg.Wait()
    fmt.Printf("opened ports: %v\n", ports)
}