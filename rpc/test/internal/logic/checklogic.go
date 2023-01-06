package logic

import (
	"context"

	"rapid-development-microservices/rpc/test/internal/svc"
	"rapid-development-microservices/rpc/test/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *test.CheckReq) (*test.CheckResp, error) {
	// todo: add your logic here and delete this line

	return &test.CheckResp{}, nil
}
