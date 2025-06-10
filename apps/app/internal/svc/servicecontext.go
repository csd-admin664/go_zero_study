package svc

import (
	"github.com/csd-admin664/go_zero_study/apps/app/internal/config"
	"github.com/csd-admin664/go_zero_study/apps/order/rpc/orderClient"
	"github.com/csd-admin664/go_zero_study/apps/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   orderclient.Order
	ProductRPC productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
