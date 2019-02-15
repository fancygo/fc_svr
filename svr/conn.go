package svr

import (
	"fmt"
	"github.com/fancygo/fc_svr/iface"
	"net"
)

type Conn struct {
	*net.TCPConn
	ConnID int

	isClosed bool

	ExitBuffChan chan bool

	Router iface.IRouter
}

func NewConn(conn *net.TCPConn, connID int, router iface.IRouter) *Conn {
	c := &Conn{
		TCPConn:      conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

func (this *Conn) StartReader() {
	fmt.Println("reader goroutine is running")
	defer fmt.Println(this.RemoteAddr().String(), " conn reader exit")
	defer this.Stop()

	for {
		buf := make([]byte, 512)
		_, err := this.TCPConn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			this.ExitBuffChan <- true
			continue
		}

		req := &Request{
			conn: this,
			data: buf,
		}

		go func(request iface.IRequest) {
			this.Router.PreHandle(request)
			this.Router.Handle(request)
			this.Router.PostHandle(request)
		}(req)
	}
}

func (this *Conn) Start() {
	go this.StartReader()

	for {
		select {
		case <-this.ExitBuffChan:
			return
		}
	}
}

func (this *Conn) Stop() {
	if this.isClosed {
		return
	}
	this.isClosed = true
	this.TCPConn.Close()
	this.ExitBuffChan <- true
	close(this.ExitBuffChan)
}

func (this *Conn) GetTCPConn() *net.TCPConn {
	return this.TCPConn
}

func (this *Conn) GetConnID() int {
	return this.ConnID
}

func (this *Conn) RemoteAddr() net.Addr {
	return this.TCPConn.RemoteAddr()
}

func (this *Conn) Send(data []byte) error {
	return nil
}

func (this *Conn) SendBuff(data []byte) error {
	return nil
}
