package adapter

import (
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"github.com/zbwang163/ad_account_server/biz/service/dto"
	"github.com/zbwang163/ad_account_server/common/utils"
	"github.com/zbwang163/ad_common/biz_error"
)

type DtoAdapter interface {
	LoginDataAdapter(*dto.UserDTO, *biz_error.BizError) *accountRpc.LoginResponse
}

type DtoAdapterV1 struct{}

func newDtoAdapterV1() DtoAdapter {
	return &DtoAdapterV1{}
}

func (d DtoAdapterV1) LoginDataAdapter(userDTO *dto.UserDTO, bizError *biz_error.BizError) *accountRpc.LoginResponse {
	if userDTO == nil {
		return nil
	}
	data := &accountRpc.LoginResponse_LoginData{
		Name:      userDTO.Name,
		AvatarUrl: userDTO.AvatarUrl,
		SessionId: userDTO.SessionId,
	}
	// 固定写法
	resp := &accountRpc.LoginResponse{}
	resp.Data = data
	resp.Base = utils.GetBaseResp(bizError)
	return resp
}
