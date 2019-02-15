package svr

import (
	"github.com/fancygo/fc_svr/iface"
)

type Request struct {
	conn iface.IConn
	data []byte
}

func (this *Request) GetConn() iface.IConn {
	return this.conn
}

func (this *Request) GetData() []byte {
	return this.data
}
