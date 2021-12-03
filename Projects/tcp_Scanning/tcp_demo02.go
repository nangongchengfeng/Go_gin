package main
 
import (
    "fmt"
    "net"
)
 
func main() {
    for port := 80; port < 100; port++ {
        conn, err := net.Dial("tcp", fmt.Sprintf("www.baidu.com:%d", port))
        if err == nil {
            conn.Close()
            fmt.Println("Connection successful")
        } else {
            fmt.Println(err)
        }
    }
}