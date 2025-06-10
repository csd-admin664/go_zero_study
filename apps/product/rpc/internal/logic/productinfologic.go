package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/apps/product/rpc/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoLogic {
	return &ProductInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductInfoLogic) ProductInfo(in *product.ProductInfoRequest) (*product.ProductItem, error) {

	return &product.ProductItem{
		ProductId:  123123,
		Name:       "测试商品",
		Price:      123.12,
		Stock:      123,
		Status:     1,
		CreateTime: 123123,
	}, nil
}
