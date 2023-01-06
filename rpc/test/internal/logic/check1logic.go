package logic

import (
	"context"

	"rapid-development-microservices/rpc/test/internal/svc"
	"rapid-development-microservices/rpc/test/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type Check1Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheck1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Check1Logic {
	return &Check1Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Check1Logic) Check1(in *test.CheckReq1) (*test.CheckResp1, error) {
	// todo: add your logic here and delete this line

	return &test.CheckResp1{}, nil
}
