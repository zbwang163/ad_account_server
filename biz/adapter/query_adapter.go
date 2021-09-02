package adapter

import (
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"github.com/zbwang163/ad_account_server/biz/service/query"
)

type QueryAdapter interface {
	LoginRequestAdapter(*accountRpc.LoginRequest) *query.LoginQuery
}

type QueryAdapterV1 struct{}

func newQueryAdapterV1() QueryAdapter {
	return &QueryAdapterV1{}
}

func (QueryAdapterV1) LoginRequestAdapter(in *accountRpc.LoginRequest) *query.LoginQuery {
	if in == nil {
		return nil
	}
	return &query.LoginQuery{
		NameOrEmail: in.NameOrEmail,
		Password:    in.Password,
		Captcha:     in.Captcha,
	}
}
