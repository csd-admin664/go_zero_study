package user

import (
	"context"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/rpc/user"

	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/svc"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/api/internal/types"

	"github.com/jinzhu/copier" // 添加 copier 导入
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) (resp *types.UserRegisterRes, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.UserRegisterReq{
		Password: req.Password,
		Phone:    req.Phone,
	})

	if err != nil {
		return nil, err
	}

	var resp1 types.UserRegisterRes
	_ = copier.Copy(&resp1, registerResp)

	return &resp1, nil

}
