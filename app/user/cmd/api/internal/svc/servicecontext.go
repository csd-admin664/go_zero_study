package svc

import (
	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/config"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User // 初始化后的 RPC 客户端接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
