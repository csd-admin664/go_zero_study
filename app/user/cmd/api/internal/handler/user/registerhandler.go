package user

import (
	"net/http"

	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/logic/user"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/svc"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// register
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
