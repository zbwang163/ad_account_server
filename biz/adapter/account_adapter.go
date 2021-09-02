package adapter

import (
	"context"
	"github.com/zbwang163/ad_account_server/biz/service/dto"
	"github.com/zbwang163/ad_account_server/biz/service/query"
	"github.com/zbwang163/ad_common/biz_error"

	"github.com/go-playground/validator"
)

type AccountAdapter interface {
	GetQueryAdapter() QueryAdapter
	GetDtoAdapter() DtoAdapter
	Login(context.Context, *query.LoginQuery) (*dto.UserDTO, *biz_error.BizError)
}

type AccountAdapterV1 struct {
	QueryAdapter QueryAdapter
	DtoAdapter   DtoAdapter
}

func NewAccountAdapter() AccountAdapter {
	return &AccountAdapterV1{
		QueryAdapter: newQueryAdapterV1(),
		DtoAdapter:   newDtoAdapterV1(),
	}
}

func (a AccountAdapterV1) Login(ctx context.Context, query *query.LoginQuery) (*dto.UserDTO, *biz_error.BizError) {
	err := validator.New().Struct(query)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return &dto.UserDTO{
		Name:      "王志斌",
		AvatarUrl: "https://wangzhibin.jpg",
		SessionId: "qisdfzcxvr",
	}, nil
}

func (a AccountAdapterV1) GetQueryAdapter() QueryAdapter {
	return a.QueryAdapter
}

func (a AccountAdapterV1) GetDtoAdapter() DtoAdapter {
	return a.DtoAdapter
}
