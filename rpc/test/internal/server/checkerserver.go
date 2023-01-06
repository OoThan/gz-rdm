// Code generated by goctl. DO NOT EDIT!
// Source: test.proto

package server

import (
	"context"

	"rapid-development-microservices/rpc/test/internal/logic"
	"rapid-development-microservices/rpc/test/internal/svc"
	"rapid-development-microservices/rpc/test/test"
)

type CheckerServer struct {
	svcCtx *svc.ServiceContext
	test.UnimplementedCheckerServer
}

func NewCheckerServer(svcCtx *svc.ServiceContext) *CheckerServer {
	return &CheckerServer{
		svcCtx: svcCtx,
	}
}

func (s *CheckerServer) Check(ctx context.Context, in *test.CheckReq) (*test.CheckResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}

func (s *CheckerServer) Check1(ctx context.Context, in *test.CheckReq1) (*test.CheckResp1, error) {
	l := logic.NewCheck1Logic(ctx, s.svcCtx)
	return l.Check1(in)
}