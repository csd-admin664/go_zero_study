// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package handler

import (
	"net/http"

	"github.com/csd-admin664/go_zero_study/apps/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 购物车
				Method:  http.MethodPost,
				Path:    "/v1/cartList",
				Handler: CartListHandler(serverCtx),
			},
			{
				// 分类
				Method:  http.MethodPost,
				Path:    "/v1/category",
				Handler: CategoryListHandler(serverCtx),
			},
			{
				// 抢购
				Method:  http.MethodGet,
				Path:    "/v1/flashsale",
				Handler: FlashSaleHandler(serverCtx),
			},
			{
				// 首页的banner
				Method:  http.MethodGet,
				Path:    "/v1/home/banner",
				Handler: HomeBannerHandler(serverCtx),
			},
			{
				// 订单列表
				Method:  http.MethodPost,
				Path:    "/v1/order/list",
				Handler: OrderListHandler(serverCtx),
			},
			{
				// 商品评论
				Method:  http.MethodPost,
				Path:    "/v1/product/comment",
				Handler: ProductCommentHandler(serverCtx),
			},
			{
				// 商品详情
				Method:  http.MethodPost,
				Path:    "/v1/product/detail",
				Handler: ProductDetailHandler(serverCtx),
			},
			{
				// 推荐
				Method:  http.MethodPost,
				Path:    "/v1/recommend",
				Handler: RecommendHandler(serverCtx),
			},
		},
	)
}
