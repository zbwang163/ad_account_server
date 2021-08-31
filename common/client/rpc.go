package client

import (
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"google.golang.org/grpc"
	"log"
)

var (
	AccountClient accountRpc.AccountRouterClient
)

func init() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	if AccountClient == nil {
		AccountClient = accountRpc.NewAccountRouterClient(conn)
	}
}
