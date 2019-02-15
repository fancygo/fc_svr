package iface

type IRequest interface {
	GetConn() IConn
	GetData() []byte
}
