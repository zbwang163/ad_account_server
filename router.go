package main

//import (
//	"context"
//	accountRpc "github.com/zbwang163/ad_account_overpass"
//	"github.com/zbwang163/ad_account_server/biz/controller"
//	"log"
//)
//
//type server struct {
//	accountRpc.UnimplementedAccountServiceServer
//	AccountAdapter controller.AccountAdapter
//}
//
//func NewServer() *server {
//	return &server{
//		AccountAdapter: controller.NewAccountController(),
//	}
//}
//
//func (s *server) Login(ctx context.Context, req *accountRpc.LoginRequest) (*accountRpc.LoginResponse, error) {
//	query := s.AccountAdapter.GetQueryAdapter().LoginRequestAdapter(req)
//	log.Println("login called")
//	dto, bizError := s.AccountAdapter.Login(ctx, query)
//	log.Printf("login dto:%v", dto)
//	return s.AccountAdapter.GetDtoAdapter().LoginDataAdapter(dto, bizError), nil
//}
