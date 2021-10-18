package query

import (
	"context"
	"github.com/zbwang163/ad_account_server/common/consts"
)

type LoginQuery struct {
	Email    string `json:"email" validate:"required,email" binding:"required,email"`                  //邮箱
	Password string `json:"password" validate:"required,min=6,max=20" binding:"required,min=6,max=20"` //密码
	Captcha  string `json:"captcha" validate:"required,len=4"  binding:"required,len=4"`               // 验证码
	Token    string `json:"token" validate:"required"  binding:"required"`                             //token
}

type RegisterQuery struct {
	UserName     string `json:"user_name" binding:"required,min=1,max=15"`      //用户名
	Password     string `json:"password" binding:"required,min=6,max=20"`       //密码
	Password2    string `json:"password_2" binding:"required,eqfield=Password"` //二次输入的密码
	Email        string `json:"email" binding:"required,email"`                 //邮箱
	EmailCapture string `json:"email_capture" binding:"required,len=4"`         //邮箱验证码
}

type SendEmailCapture struct {
	Email string `json:"email" binding:"required,email"`
}

type GetCoreUserIdFromSessionQuery struct {
	SessionId string `json:"session_id" binding:"required"`
}

type UserInfoQuery struct {
	CtxUserId int64 `validate:"required,gt=0" binding:"-"`
}

func (q *UserInfoQuery) Bind(ctx context.Context) {
	ctxUserId := ctx.Value(consts.CtxUserId).(int64)
	q.CtxUserId = ctxUserId
}

type UpdateUserInfoQuery struct {
	UserName  string `json:"user_name" binding:"omitempty,min=1,max=15"` //用户名
	AvatarUrl string `json:"avatar_url" binding:"omitempty,url"`
	Slogan    string `json:"slogan" binding:"omitempty,max=100"`

	CtxUserId int64 `validate:"required,gt=0" binding:"-"`
}

func (q *UpdateUserInfoQuery) Bind(ctx context.Context) {
	ctxUserId := ctx.Value(consts.CtxUserId).(int64)
	q.CtxUserId = ctxUserId
}

type PolicyQuery struct {
	PType      string `json:"ptype" binding:"required"`
	Subject    string `json:"subject" binding:"required"`
	Domain     string `json:"domain"`
	Object     string `json:"object"  binding:"required"`
	Action     string `json:"action"`
	Effect     string `json:"effect"`
	Expiration string `json:"expiration"`
}

type GetRulesByRoleQuery struct {
	Role       string `json:"role"`
	CoreUserId string `json:"core_user_id"`
}

type QuestionQuery struct {
	Word   string   `json:"word"`
	Images []string `json:"images"`
	ToUid  string   `json:"to_uid"`

	CtxUserId int64 `validate:"required,gt=0" binding:"-"`
}

type PageQuery struct {
	Page  int64 `json:"page" binding:"min=1,max=100"`
	Limit int64 `json:"limit" binding:"min=1,max=100"`
}

type AddAnswerQuery struct {
	RealPrize          int64  `json:"real_prize"`
	ProductDescription string `json:"product_description"`

	CtxUserId int64 `validate:"required,gt=0" binding:"-"`
}

type ConfirmOrderQuery struct {
	OrderId string `json:"order_id"`

	CtxUserId int64 `validate:"required,gt=0" binding:"-"`
}
