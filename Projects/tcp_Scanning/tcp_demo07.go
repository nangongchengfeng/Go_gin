package main
 
import (
	"code/Projects/tcp_Scanning/golimit"
	"flag"
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
	hostname := flag.String("hostname", "", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	golimits := flag.Int("golimit", 1000, "the Program Concurrency")
	flag.Parse()
	ports := []int{}
	//timeout := time.Millisecond * 500
	g := golimit.NewGoLimit(*golimits)
	for port := *startPort; port <= *endPort; port++ {
		g.Add()
		go func(g *golimit.GoLimit, p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				log.Printf("端口: %d 已经开通", p)
				mutex.Unlock()
			}
			g.Done()
		}(g, port)
	}
	time.Since(startTime)
	cost := int(time.Since(startTime) / time.Second)
	fmt.Printf("opened ports: %v\n", ports)
 
	fmt.Printf("代码运行时长: %d S", cost)
}