package logic

import (
	"context"
	"github.com/csd-admin664/go_zero_study/apps/product/rpc/product"

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
	info, err := l.svcCtx.ProductRPC.ProductInfo(l.ctx, &product.ProductInfoRequest{ProductId: req.ProductID})
	if err != nil {
		return nil, err
	}
	return &types.ProductDetailResponse{
		Product: &types.Product{
			Name:        info.Name,
			Description: info.Description,
			Price:       info.Price, Stock: info.Stock,
			Status:     info.Status,
			CreateTime: info.CreateTime,
		},
		Comments: []*types.Comment{
			{
				ID:         1,
				Content:    "测试评论",
				CreateTime: 1630000000,
				User: &types.User{
					ID:     1,
					Name:   "测试用户",
					Avatar: "https://avatar.csd-admin664.com/avatar.png",
				},
			},
		},
	}, nil
}
