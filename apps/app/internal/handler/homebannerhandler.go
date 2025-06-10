package handler

import (
	"net/http"

	"github.com/csd-admin664/go_zero_study/apps/app/internal/logic"
	"github.com/csd-admin664/go_zero_study/apps/app/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 首页的banner
func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
