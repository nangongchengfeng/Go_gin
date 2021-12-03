# Go_gin



编写



（1）编写TCP扫描器

# TCP扫描增强器

TCP扫描增强器实现原理，主要是使用TCP三次握手原理

TCP是比我们介绍的要复杂的多，但是我们只介绍一点基础知识。TCP的握手有三个过程。

首先，客户端发送一个 syn 的包，表示建立回话的开始。如果客户端收到超时，说明端口可能在防火墙后面，或者没有启用服务器

![img](https://img-blog.csdnimg.cn/img_convert/ba16919769276de3f44bb04257ac8f08.png)

第二，如果服务端应答 syn-ack 包，意味着这个端口是打开的，否则会返回 rst 包。最后，客户端需要另外发送一个 ack 包。从这时起，连接就已经建立。

![img](https://img-blog.csdnimg.cn/img_convert/a3ef55f9200aab9e8608092bb070b905.png)

![img](https://img-blog.csdnimg.cn/img_convert/2ec6eb21d371a22dd16150d93df230ad.png)

我们TCP扫描器第一步先实现单个端口的测试。使用标准库中的 net.Dial 函数，该函数接收两个参数：协议和测试地址（带端口号）

```go
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

```

