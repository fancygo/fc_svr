package svr

import (
	"fmt"
	"github.com/fancygo/fc_svr/iface"
	"github.com/fancygo/fc_svr/utils"
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
	utils.GlobalConfig.Reload()
	fmt.Println(utils.GlobalConfig)
	s := &Server{
		Name:      utils.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalConfig.Host,
		Port:      utils.GlobalConfig.TcpPort,
		Router:    nil,
	}
	return s
}
