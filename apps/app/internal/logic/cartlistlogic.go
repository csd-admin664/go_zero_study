package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/apps/app/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 购物车
func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartListLogic) CartList(req *types.CartListRequest) (resp *types.CartListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
