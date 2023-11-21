package main

import (
	"fmt"
	"net"
	"time"
)

/**
模拟客户端
*/

func main() {
	fmt.Printf("client start ...\n")
	time.Sleep(1 * time.Second)

	// 1. 连接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Printf("clinet start err exit!\n")
		return
	}

	// 2. 调用Write 写数据
	for {
		_, err := conn.Write([]byte("Hello zinx V0.1.. !"))
		if err != nil {
			fmt.Printf("write conn err:%s\n", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read buf err:%s\n", err)
			return
		}

		fmt.Printf("server call back:%s, cnt = %d\n", buf, cnt)

		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
