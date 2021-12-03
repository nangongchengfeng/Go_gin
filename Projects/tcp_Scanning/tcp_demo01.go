package main
 
import (
    "fmt"
    "net"
)
 
func main() {
	_, err := net.Dial("tcp", "www.baidu.com:80")
    if err == nil {
        fmt.Println("Connection successful")
    } else {
        fmt.Println(err)
    }
}