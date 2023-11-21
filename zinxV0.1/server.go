package main

import "GoZinx/znet"

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.Server()
}
