// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package checker

import (
	"context"

	"rapid-development-microservices/rpc/check/check"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckReq  = check.CheckReq
	CheckResp = check.CheckResp

	Checker interface {
		Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error)
	}

	defaultChecker struct {
		cli zrpc.Client
	}
)

func NewChecker(cli zrpc.Client) Checker {
	return &defaultChecker{
		cli: cli,
	}
}

func (m *defaultChecker) Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error) {
	client := check.NewCheckerClient(m.cli.Conn())
	return client.Check(ctx, in, opts...)
}
