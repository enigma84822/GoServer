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

// test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Printf("Call Router PreHandle\n")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Printf("Call back before ping error\n")
	}
}

// test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Printf("Call Router Handle\n")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping...\n"))
	if err != nil {
		fmt.Printf("Call back ping... ping... error\n")
	}
}

// test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Printf("Call Router PostHandle\n")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Printf("Call back after ping... error\n")
	}
}

func main() {
	// 1,创建server
	s := znet.NewServer("[zinx V0.3]")

	// 2.添加自定义Router
	s.AddRouter(&PingRouter{})
	// 3.启动服务
	s.Server()
}
