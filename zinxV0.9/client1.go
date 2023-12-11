package main

import (
	"GoZinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

/**
模拟客户端
*/

func main() {
	fmt.Printf("client 1 start ...\n")
	time.Sleep(1 * time.Second)

	// 1. 连接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Printf("clinet start err exit!\n")
		return
	}

	// 2. 调用Write 写数据
	for {
		// 发送封包的消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(1, []byte("ZinxV0.9 client 1 test message")))
		if err != nil {
			fmt.Printf("pack err:%s\n", err)
			break
		}

		// 写数据
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Printf("write err:%s\n", err)
			break
		}

		// 读数据
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Printf("read head err:%s\n", err)
			break
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Printf("client unpack msghead err:%s\n", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Printf("read msg data err:%s\n", err)
				break
			}
			fmt.Printf("---> Recv server msg id=%d len=%d data=%s\n", msg.Id, msg.DataLen, msg.Data)
		}

		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
