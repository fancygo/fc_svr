package iface

import (
	"net"
)

type IConn interface {
	Start()
	Stop()
	GetTCPConn() *net.TCPConn
	GetConnID() uint32
	RemoteAdr() net.Addr
	Send(data []byte) error
	SendBuff(data []byte) error
}

type HandFunc func(*net.TCPConn, []byte, int) error
