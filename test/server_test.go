package test

import (
	"context"
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"github.com/zbwang163/ad_account_server/common/client"
	"testing"
)

func TestLogin(t *testing.T) {
	resp, err := client.AccountClient.Login(context.Background(), &accountRpc.LoginRequest{})
	if err != nil {
		t.Fatalf("err :%v", err)
	}
	t.Log(resp)
}
