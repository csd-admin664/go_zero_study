// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: order.proto

package server

import (
	"context"

	"github.com/csd-admin664/go_zero_study/apps/order/rpc/internal/logic"
	"github.com/csd-admin664/go_zero_study/apps/order/rpc/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) Orders(ctx context.Context, in *order.OrdersRequest) (*order.OrdersResponse, error) {
	l := logic.NewOrdersLogic(ctx, s.svcCtx)
	return l.Orders(in)
}
