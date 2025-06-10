package user

import (
	"context"

	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/svc"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginRes, err error) {
	// todo: add your logic here and delete this line

	return
}
