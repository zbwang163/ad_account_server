package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/zbwang163/ad_account_server/biz/service"
	"github.com/zbwang163/ad_account_server/biz/service/dto"
	"github.com/zbwang163/ad_account_server/biz/service/po"
	"github.com/zbwang163/ad_account_server/biz/service/query"
	"github.com/zbwang163/ad_account_server/common/biz_error"
	"github.com/zbwang163/ad_account_server/common/utils"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController() AccountController {
	return &AccountControllerImpl{
		AccountService: service.NewAccountService(),
	}
}

type AccountController interface {
	Register(context.Context, *query.RegisterQuery) (*dto.UserInfoDTO, error)
	Login(context.Context, *query.LoginQuery) (*dto.UserInfoDTO, error)
	GetCaptureImage(context.Context) (*dto.GetCaptureDTO, error)
	SendEmailCapture(context.Context, *query.SendEmailCapture) error
	GetCoreUserIdBySession(context.Context, *query.GetCoreUserIdFromSessionQuery) (*dto.UserInfoDTO, error)
	GetUserInfo(context.Context, *query.UserInfoQuery) (*dto.UserInfoDTO, error)                        //登录
	UpdateUserInfo(ctx context.Context, infoQuery *query.UpdateUserInfoQuery) (*dto.UserInfoDTO, error) //登录

	AddPolicyRule(ctx context.Context, policyQuery *query.PolicyQuery) error
	RemovePolicyRule(ctx context.Context, policyQuery *query.PolicyQuery) error

	GetAnswerList(ctx context.Context, pageQuery *query.PageQuery) (*dto.AnswerListDTO, error)
	AddAnswer(ctx context.Context, answerQuery *query.AddAnswerQuery) error
	AskQuestion(ctx context.Context, questionQuery *query.QuestionQuery) error
	ReplyQuestion(ctx context.Context, questionQuery *query.QuestionQuery) error
	ConfirmOrder(ctx context.Context, query *query.ConfirmOrderQuery) error
}

func (a AccountControllerImpl) GetAnswerList(ctx context.Context, pageQuery *query.PageQuery) (*dto.AnswerListDTO, error) {
	panic("implement me")
}

func (a AccountControllerImpl) AddAnswer(ctx context.Context, answerQuery *query.AddAnswerQuery) error {
	panic("implement me")
}

func (a AccountControllerImpl) AskQuestion(ctx context.Context, questionQuery *query.QuestionQuery) error {
	panic("implement me")
}

func (a AccountControllerImpl) ReplyQuestion(ctx context.Context, questionQuery *query.QuestionQuery) error {
	panic("implement me")
}

func (a AccountControllerImpl) ConfirmOrder(ctx context.Context, query *query.ConfirmOrderQuery) error {
	panic("implement me")
}

func (a AccountControllerImpl) RemovePolicyRule(ctx context.Context, policyQuery *query.PolicyQuery) error {
	err := validator.New().Struct(policyQuery)
	if err != nil {
		return biz_error.NewParamError(err)
	}

	a.AccountService.RemovePolicyRule(ctx, policyQuery.PType, policyQuery.Subject, policyQuery.Domain, policyQuery.Object, policyQuery.Action, policyQuery.Effect, policyQuery.Expiration)
	return nil
}

func (a AccountControllerImpl) AddPolicyRule(ctx context.Context, policyQuery *query.PolicyQuery) error {
	err := validator.New().Struct(policyQuery)
	if err != nil {
		return biz_error.NewParamError(err)
	}

	a.AccountService.AddPolicyRule(ctx, policyQuery.PType, policyQuery.Subject, policyQuery.Domain, policyQuery.Object, policyQuery.Action, policyQuery.Effect, policyQuery.Expiration)
	return nil
}

func (a AccountControllerImpl) UpdateUserInfo(ctx context.Context, infoQuery *query.UpdateUserInfoQuery) (*dto.UserInfoDTO, error) {
	infoQuery.Bind(ctx)
	err := validator.New().Struct(infoQuery)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}

	return a.AccountService.UpdateUserInfo(ctx, infoQuery)
}

func (a AccountControllerImpl) GetUserInfo(ctx context.Context, infoQuery *query.UserInfoQuery) (*dto.UserInfoDTO, error) {
	infoQuery.Bind(ctx)
	err := validator.New().Struct(infoQuery)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return a.AccountService.GetUserInfoByCoreUserId(ctx, infoQuery.CtxUserId)
}

func (a AccountControllerImpl) GetCoreUserIdBySession(ctx context.Context, sessionQuery *query.GetCoreUserIdFromSessionQuery) (*dto.UserInfoDTO, error) {
	err := validator.New().Struct(sessionQuery)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}

	coreUserId := utils.GetCoreUserIdFromSession(sessionQuery.SessionId)
	return &dto.UserInfoDTO{CoreUserId: coreUserId}, nil
}

func (a AccountControllerImpl) SendEmailCapture(ctx context.Context, query *query.SendEmailCapture) error {
	err := validator.New().Struct(query)
	if err != nil {
		return biz_error.NewParamError(err)
	}

	capture, err := a.AccountService.GenCapture(ctx, 4)
	if err != nil {
		return err
	}
	a.AccountService.AsyncSendEmailCapture(ctx, query.Email, capture)
	return nil
}

func (a AccountControllerImpl) Register(ctx context.Context, registerQuery *query.RegisterQuery) (*dto.UserInfoDTO, error) {
	err := validator.New().Struct(registerQuery)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}

	// 保存
	return a.AccountService.SaveUserInfo(ctx, registerQuery)
}

func (a AccountControllerImpl) Login(ctx context.Context, loginQuery *query.LoginQuery) (*dto.UserInfoDTO, error) {
	err := validator.New().Struct(loginQuery)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}

	// 校验capture是否过期、是否正确
	if ok, err := a.AccountService.ValidateCapture(ctx, loginQuery.Token, loginQuery.Captcha); !ok && err != nil {
		return nil, err
	}
	// 校验密码是否正确
	var user *po.AdUserInfoPo
	if user, err = a.AccountService.ValidateUser(ctx, loginQuery.Email, loginQuery.Password); user == nil && err != nil {
		return nil, err
	}
	// 生成session id
	sessionId, err := a.AccountService.GenSessionId(ctx, user)
	if err != nil {
		return nil, err
	}
	return &dto.UserInfoDTO{SessionId: sessionId}, nil
}

func (a AccountControllerImpl) GetCaptureImage(ctx context.Context) (*dto.GetCaptureDTO, error) {
	captureUrl, token, err := a.AccountService.GenCaptureImage(ctx, 4)
	if err != nil {
		return nil, err
	}
	return &dto.GetCaptureDTO{Token: token, CaptureUrl: captureUrl}, nil
}
