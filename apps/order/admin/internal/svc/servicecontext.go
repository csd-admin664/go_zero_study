package svc

import (
	"github.com/csd-admin664/go_zero_study/apps/order/admin/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
