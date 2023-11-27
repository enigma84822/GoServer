package main

import "GoZinx/znet"

func main() {
	s := znet.NewServer("[zinx V0.2]")
	s.Server()
}
