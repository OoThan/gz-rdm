// Code generated by goctl. DO NOT EDIT!
// Source: add.proto

package server

import (
	"context"

	"rapid-development-microservices/rpc/add/add"
	"rapid-development-microservices/rpc/add/internal/logic"
	"rapid-development-microservices/rpc/add/internal/svc"
)

type AdderServer struct {
	svcCtx *svc.ServiceContext
	add.UnimplementedAdderServer
}

func NewAdderServer(svcCtx *svc.ServiceContext) *AdderServer {
	return &AdderServer{
		svcCtx: svcCtx,
	}
}

func (s *AdderServer) Add(ctx context.Context, in *add.AddReq) (*add.AddResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}
