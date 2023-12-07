package main

import (
	"GoZinx/ziface"
	"GoZinx/znet"
	"fmt"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Printf("Call PingRouter Handle\n")
	// 1.读取客户端的数据
	fmt.Printf("recv from client msgId=%d data=%s\n", request.GetMsgID(), request.GetData())
	// 2.回写ping
	err := request.GetConnection().SendMsg(200, []byte("ping... ping... ping..."))
	if err != nil {
		fmt.Printf("send msg err:%s\n", err)
	}
}

// hello test 自定义路由
type HelloRouter struct {
	znet.BaseRouter
}

// test Handle
func (this *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Printf("Call HelloRouter Handle\n")
	// 1.读取客户端的数据
	fmt.Printf("recv from client msgId=%d data=%s\n", request.GetMsgID(), request.GetData())
	// 2.回写ping
	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome to Zinx!"))
	if err != nil {
		fmt.Printf("send msg err:%s\n", err)
	}
}

func main() {
	// 1,创建server
	s := znet.NewServer("[zinx V0.8]")

	// 2.添加自定义Router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	// 3.启动服务
	s.Server()
}
