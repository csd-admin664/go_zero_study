package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/app/user/cmd/rpc/internal/svc"
	"github.com/csd-admin664/go_zero_study/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.UserRegisterReq) (*pb.UserRegisterRes, error) {
	// todo: add your logic here and delete this line

	return &pb.UserRegisterRes{
		Accesstonken: "123123123123TOlke",
	}, nil
}
