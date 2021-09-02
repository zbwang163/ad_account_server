package main

import (
	"context"
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"github.com/zbwang163/ad_account_server/biz/adapter"
	"log"
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

func (s *server) Login(ctx context.Context, req *accountRpc.LoginRequest) (*accountRpc.LoginResponse, error) {
	query := s.AccountAdapter.GetQueryAdapter().LoginRequestAdapter(req)
	log.Println("login called")
	dto, bizError := s.AccountAdapter.Login(ctx, query)
	log.Printf("login dto:%v", dto)
	return s.AccountAdapter.GetDtoAdapter().LoginDataAdapter(dto, bizError), nil
}
