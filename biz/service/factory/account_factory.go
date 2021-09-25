package factory

import (
	"context"
	"github.com/zbwang163/ad_account_server/biz/service/dto"
	"github.com/zbwang163/ad_account_server/biz/service/po"
)

func BuildUserInfoDTO(ctx context.Context, po *po.AdUserInfoPo) *dto.UserInfoDTO {
	if po == nil {
		return nil
	}
	return &dto.UserInfoDTO{
		CoreUserId: po.CoreUserId,
		UserName:   po.UserName,
		AvatarUrl:  po.AvatarUrl,
		Slogan:     po.Slogan,
	}
}
