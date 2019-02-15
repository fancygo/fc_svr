package svr

import (
	"github.com/fancygo/fc_svr/iface"
)

type BaseRouter struct {
}

func (this *BaseRouter) PreHandle(req iface.IRequest) {
}

func (this *BaseRouter) Handle(req iface.IRequest) {
}

func (this *BaseRouter) PostHandle(req iface.IRequest) {
}
