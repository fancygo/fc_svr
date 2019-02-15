package main

import (
	"fmt"
	"github.com/fancygo/fc_svr/iface"
	"github.com/fancygo/fc_svr/svr"
)

type PingRouter struct {
	svr.BaseRouter
}

func (this *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("call prehandle")
	_, err := request.GetConn().GetTCPConn().Write([]byte("before ping ...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("call handle")
	_, err := request.GetConn().GetTCPConn().Write([]byte("ping ...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func (this *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("call psthandle")
	_, err := request.GetConn().GetTCPConn().Write([]byte("after ping ...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func main() {
	fmt.Println("vim-go")
	s := svr.NewServer("game_svr")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
