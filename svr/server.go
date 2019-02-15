package svr

import (
	"errors"
	"fmt"
	"github.com/fancygo/fc_svr/iface"
	"net"
	"time"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iface.IRouter
}

func CallBackToClient(conn *net.TCPConn, data []byte, count int) error {
	fmt.Println("CallbackToClient ...")
	if _, err := conn.Write(data[:count]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallbackToClient error")
	}
	return nil
}

func (this *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(this.IPVersion, fmt.Sprintf("%s:%d", this.IP, this.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listenner, err := net.ListenTCP(this.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", this.IPVersion, "err", err)
			return
		}

		fmt.Println("start server ", this.Name, " now listenning")

		cid := int(0)

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			dealConn := NewConn(conn, cid, this.Router)
			cid++
			go dealConn.Start()
		}
	}()
}

func (this *Server) Stop() {
	fmt.Println("stop server")
}

func (this *Server) Serve() {
	this.Start()
	//阻塞，不应该用sleep
	for {
		time.Sleep(10 * time.Second)
	}
}

func (this *Server) AddRouter(router iface.IRouter) {
	this.Router = router
	fmt.Println("Add Router")
}

func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7002,
		Router:    nil,
	}
	return s
}
