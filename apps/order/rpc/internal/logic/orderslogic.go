package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/apps/order/rpc/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	// todo: add your logic here and delete this line

	return &order.OrdersResponse{}, nil
}
