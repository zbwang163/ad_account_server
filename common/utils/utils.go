package utils

import (
	"github.com/zbwang163/ad_base_overpass"
	"github.com/zbwang163/ad_common/biz_error"
)

func GetBaseResp(bizError *biz_error.BizError) *ad_base_overpass.BaseResponse {
	baseResp := &ad_base_overpass.BaseResponse{
		Extra: make(map[string]string, 2),
	}
	if bizError == nil || bizError.Code == 0 {
		baseResp.StatusCode = 0
		baseResp.StatusMessage = "success"
	} else {
		baseResp.StatusCode = int32(bizError.Code)
		baseResp.StatusMessage = bizError.Message
		baseResp.Extra["error_msg"] = bizError.ErrorMsg
	}
	return baseResp
}
