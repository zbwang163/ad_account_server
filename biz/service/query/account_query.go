package query

import (
	"github.com/zbwang163/ad_base_overpass"
)

type LoginQuery struct {
	NameOrEmail string                        `protobuf:"bytes,1,opt,name=name_or_email,json=nameOrEmail,proto3" json:"name_or_email,omitempty"` //用户名或邮箱
	Password    string                        `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`                            //登录密码
	Captcha     string                        `protobuf:"bytes,3,opt,name=captcha,proto3" json:"captcha,omitempty"`                              //验证码
	Base        *ad_base_overpass.BaseRequest `protobuf:"bytes,15,opt,name=base,proto3" json:"base,omitempty"`
}
