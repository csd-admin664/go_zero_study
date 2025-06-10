package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/app/user/cmd/rpc/internal/svc"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.UserLoginReq) (*pb.UserLoginRes, error) {

	return &pb.UserLoginRes{
		Accesstonken:      "123",
		Expire:            123,
		RefreshtokenAfter: 123,
	}, nil
}
