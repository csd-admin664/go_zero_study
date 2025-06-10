package logic

import (
	"context"

	"github.com/csd-admin664/go_zero_study/apps/app/internal/svc"
	"github.com/csd-admin664/go_zero_study/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 首页的banner
func NewHomeBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerLogic {
	return &HomeBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBannerLogic) HomeBanner() (resp *types.HomeBannerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
