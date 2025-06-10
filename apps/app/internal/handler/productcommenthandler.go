package handler

import (
	"net/http"

	"github.com/csd-admin664/go_zero_study/apps/app/internal/logic"
	"github.com/csd-admin664/go_zero_study/apps/app/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 商品评论
func ProductCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductCommentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProductCommentLogic(r.Context(), svcCtx)
		resp, err := l.ProductComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
