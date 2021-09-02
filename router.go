package main

import (
	"context"
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"github.com/zbwang163/ad_account_server/biz/adapter"
)

type server struct {
	accountRpc.UnimplementedAccountServiceServer
	AccountAdapter adapter.AccountAdapter
}

func NewServer() *server {
	return &server{
		AccountAdapter: adapter.NewAccountAdapter(),
	}
}

func (s *server) Login(ctx context.Context, in *accountRpc.LoginRequest) (*accountRpc.LoginResponse, error) {
	query := s.AccountAdapter.GetQueryAdapter().LoginRequestAdapter(in)
	dto, bizError := s.AccountAdapter.Login(ctx, query)
	return s.AccountAdapter.GetDtoAdapter().LoginDataAdapter(dto, bizError), nil
}
