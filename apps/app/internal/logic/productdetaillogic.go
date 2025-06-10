package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/apps/app/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品详情
func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ProductDetailRequest) (resp *types.ProductDetailResponse, err error) {

	return
}
