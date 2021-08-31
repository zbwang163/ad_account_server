package main

import (
	"context"
	accountRpc "github.com/zbwang163/ad_info_account_rpc"
)

type server struct {
	accountRpc.UnimplementedAccountRouterServer
}

func (s *server) Login(ctx context.Context, in *accountRpc.LoginRequest) (*accountRpc.LoginResponse, error) {
	return &accountRpc.LoginResponse{
		Name:      "王志斌",
		AvatarUrl: "http://iloveyou.jpg",
		SessionId: "wasdchvbbjnk",
	}, nil
}

func (s *server) Register(ctx context.Context, in *accountRpc.RegisterRequest) (*accountRpc.LoginResponse, error) {
	panic("implement me")
}
