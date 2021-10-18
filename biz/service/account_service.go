package service

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/zbwang163/ad_account_server/biz/service/dto"
	"github.com/zbwang163/ad_account_server/biz/service/factory"
	"github.com/zbwang163/ad_account_server/biz/service/po"
	"github.com/zbwang163/ad_account_server/biz/service/query"
	"github.com/zbwang163/ad_account_server/common/biz_error"
	casbinAdapter "github.com/zbwang163/ad_account_server/common/client/casbin"
	"github.com/zbwang163/ad_account_server/common/client/minio"
	"github.com/zbwang163/ad_account_server/common/client/redis"
	"github.com/zbwang163/ad_account_server/common/client/smtp"
	"github.com/zbwang163/ad_account_server/common/consts"
	"github.com/zbwang163/ad_account_server/common/logs"
	"github.com/zbwang163/ad_account_server/common/utils"
	"gorm.io/gorm"
	"image/color"
	"image/jpeg"
	"lukechampine.com/frand"
	"time"
)

type AccountServiceImpl struct {
}

func NewAccountService() AccountService {
	return &AccountServiceImpl{}
}

type AccountService interface {
	// ValidateUser 验证用户的邮箱密码是否正确
	ValidateUser(ctx context.Context, email, password string) (po *po.AdUserInfoPo, err error)
	// ValidateCapture 校验验证码是否过期、是否相同
	ValidateCapture(ctx context.Context, token string, capture string) (ok bool, err error)
	// GenSessionId 生成session id
	GenSessionId(ctx context.Context, po *po.AdUserInfoPo) (sessionId string, err error)
	// GenCapture 生成随机len长度的验证码
	GenCapture(ctx context.Context, len int) (capture string, err error)
	// GenCaptureImage 生成验证图片，并获取url
	GenCaptureImage(ctx context.Context, len int) (url string, token string, err error)
	// SaveUserInfo 新增/更新用户信息
	SaveUserInfo(ctx context.Context, query *query.RegisterQuery) (dto *dto.UserInfoDTO, err error)
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo(ctx context.Context, query *query.UpdateUserInfoQuery) (dto *dto.UserInfoDTO, err error)
	// ExistsUserName 用户名是否存在
	ExistsUserName(ctx context.Context, userName string) (bool, error)
	// ExistsEmail 邮箱是否存在
	ExistsEmail(ctx context.Context, email string) (bool, error)
	// AsyncSendEmailCapture 发送邮件
	AsyncSendEmailCapture(ctx context.Context, email string, content string)
	// GetUserInfoByCoreUserId 获取用户信息
	GetUserInfoByCoreUserId(ctx context.Context, coreUserId int64) (dto *dto.UserInfoDTO, err error)

	AddPolicyRule(ctx context.Context, ptype string, subject string, domain string, object string, action string, effect string, expiration string) bool
	RemovePolicyRule(ctx context.Context, ptype string, subject string, domain string, object string, action string, effect string, expiration string) bool
	GetRulesByRole(ctx context.Context, role string)
}

func (svc AccountServiceImpl) GetRulesByRole(ctx context.Context, role string) {
	casbinAdapter.Enforcer.GetGroupingPolicy()
}

func (svc AccountServiceImpl) AddPolicyRule(ctx context.Context, ptype string, subject string, domain string, object string, action string, effect string, expiration string) bool {
	var ok bool
	var err error
	switch ptype {
	case "p":
		ok, err = casbinAdapter.Enforcer.AddPolicy(subject, domain, object, action, effect, expiration)
	case "g":
		ok, err = casbinAdapter.Enforcer.AddGroupingPolicy(subject, object, domain)
	case "g2":
		ok, err = casbinAdapter.Enforcer.AddNamedGroupingPolicy("g2", subject, object)
	}
	if err != nil {
		logs.CtxError(ctx, "AddPolicyRule err:%v", err)
	}
	return ok
}

func (svc AccountServiceImpl) RemovePolicyRule(ctx context.Context, ptype string, subject string, domain string, object string, action string, effect string, expiration string) bool {
	var ok bool
	var err error
	switch ptype {
	case "p":
		ok, err = casbinAdapter.Enforcer.RemovePolicy(subject, domain, object, action, effect, expiration)
	case "g":
		ok, err = casbinAdapter.Enforcer.RemoveGroupingPolicy(subject, object, domain)
	case "g2":
		ok, err = casbinAdapter.Enforcer.RemoveNamedGroupingPolicy("g2", subject, object)
	}
	if err != nil {
		logs.CtxError(ctx, "RemovePolicyRule err:%v", err)
	}
	return ok
}

func (svc AccountServiceImpl) UpdateUserInfo(ctx context.Context, query *query.UpdateUserInfoQuery) (dto *dto.UserInfoDTO, err error) {
	user := po.AdUserInfoPo{
		Model:      gorm.Model{},
		CoreUserId: query.CtxUserId,
		UserName:   query.UserName,
		AvatarUrl:  query.AvatarUrl,
		Slogan:     query.Slogan,
	}
	// 更新需要忽略空值，所以采用结构体做条件
	userPo, err := po.AdUserInfoPo{}.UpdateOrInsert(ctx, map[string]interface{}{"core_user_id": query.CtxUserId}, user)
	if err != nil {
		return nil, err
	}
	return factory.BuildUserInfoDTO(ctx, userPo), nil
}

func (svc AccountServiceImpl) GetUserInfoByCoreUserId(ctx context.Context, coreUserId int64) (dto *dto.UserInfoDTO, err error) {
	userPo, err := po.AdUserInfoPo{}.SelectOne(ctx, map[string]interface{}{"core_user_id": coreUserId})
	if err != nil {
		return nil, err
	}
	return factory.BuildUserInfoDTO(ctx, userPo), nil
}

func (svc AccountServiceImpl) AsyncSendEmailCapture(ctx context.Context, email string, content string) {
	smtp.AsyncSend(ctx, email, content)
	token := fmt.Sprintf("email-capture-%s", email)
	redis.Redis[consts.AccountPSM].Set(token, content, time.Minute*3)
}

func (svc AccountServiceImpl) ExistsEmail(ctx context.Context, email string) (bool, error) {
	userPo, err := po.AdUserInfoPo{}.SelectOne(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return false, err
	}
	if userPo != nil {
		return true, nil
	}
	return false, nil
}

func (svc AccountServiceImpl) ExistsUserName(ctx context.Context, userName string) (bool, error) {
	userPo, err := po.AdUserInfoPo{}.SelectOne(ctx, map[string]interface{}{"user_name": userName})
	if err != nil {
		return false, err
	}
	if userPo != nil {
		return true, nil
	}
	return false, nil
}

func (svc AccountServiceImpl) SaveUserInfo(ctx context.Context, query *query.RegisterQuery) (*dto.UserInfoDTO, error) {
	// 验证用户名
	if ok, _ := svc.ExistsUserName(ctx, query.UserName); ok {
		return nil, biz_error.NewParamError(errors.New("用户名已存在"))
	}

	// 验证邮箱
	if ok, _ := svc.ExistsEmail(ctx, query.Email); ok {
		return nil, biz_error.NewParamError(errors.New("邮箱已存在"))
	}

	// 验证邮箱验证码
	token := fmt.Sprintf("email-capture-%s", query.Email)
	if ok, err := svc.ValidateCapture(ctx, token, query.EmailCapture); !ok || err != nil {
		return nil, err
	}

	// 密码二次摘要
	salt := hex.EncodeToString(frand.New().Bytes(4))
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(query.Password))
	digest := salt + hex.EncodeToString(mac.Sum(nil))

	ctxUserId := ctx.Value(consts.CtxUserId).(int64)
	if ctxUserId == 0 {
		coreUserId, err := utils.GenGlobalUniqueId()
		if err != nil {
			return nil, err
		}
		ctxUserId = coreUserId
	}
	user := po.AdUserInfoPo{
		Model:      gorm.Model{},
		CoreUserId: ctxUserId,
		UserName:   query.UserName,
		Email:      query.Email,
		Password:   digest,
	}
	// 更新需要忽略空值，所以采用结构体做条件
	userPo, err := po.AdUserInfoPo{}.UpdateOrInsert(ctx, map[string]interface{}{"core_user_id": ctxUserId}, user)
	if err != nil {
		return nil, err
	}
	sessionId, _ := svc.GenSessionId(ctx, userPo)
	return &dto.UserInfoDTO{SessionId: sessionId}, nil
}

func (svc AccountServiceImpl) ValidateUser(ctx context.Context, email string, password string) (*po.AdUserInfoPo, error) {
	// 1. 取出记录
	userPo, err := po.AdUserInfoPo{}.SelectOne(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}
	dbPassword := userPo.Password
	if len(dbPassword) < 8 {
		return nil, biz_error.NewParamError(errors.New("密码错误"))
	}
	key := dbPassword[0:8]
	realPassword := dbPassword[8:]
	// 2. 密码摘要 hmacsha256
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(password))
	digest := hex.EncodeToString(mac.Sum(nil))
	// 3. 密码做比较
	if realPassword != digest {
		return nil, biz_error.NewParamError(errors.New("密码错误"))
	}
	return userPo, nil
}

func (svc AccountServiceImpl) ValidateCapture(ctx context.Context, token string, capture string) (ok bool, err error) {
	// 验证是否过期
	if redis.Redis[consts.AccountPSM].Exists(token).Val() != 1 {
		return false, biz_error.NewParamError(errors.New("验证码过期"))
	}
	// 验证是否正确
	realCapture := redis.Redis[consts.AccountPSM].Get(token).Val()
	if capture != realCapture {
		return false, biz_error.NewParamError(errors.New("验证码错误"))
	}
	return true, nil
}

func (svc AccountServiceImpl) GenSessionId(ctx context.Context, po *po.AdUserInfoPo) (sessionId string, err error) {
	if po == nil {
		return "", nil
	}
	// 随机生成4字节
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	// base64编码
	sessionId = base64.URLEncoding.EncodeToString(b)
	str, err := json.Marshal(po)
	// 存放redis，过期时间7天
	redis.Redis[consts.AccountPSM].Set(sessionId, string(str), time.Hour*24*7)
	return sessionId, nil
}

func (svc AccountServiceImpl) GenCapture(ctx context.Context, len int) (capture string, err error) {
	r := frand.New().Bytes(len)
	str := hex.EncodeToString(r)
	return str[0:len], nil
}

func (svc AccountServiceImpl) GenCaptureImage(ctx context.Context, len int) (url string, token string, err error) {
	c := captcha.New()
	c.SetFont("biz/service/resource/Verdana.ttf") // 字体选择
	c.SetSize(256, 128)                           // 尺寸
	c.SetDisturbance(captcha.MEDIUM)              //设置干扰度
	c.SetFrontColor(color.RGBA{B: 255, A: 255})   // 字体颜色
	img, capture := c.Create(len, captcha.NUM)    // 纯数字验证码
	buffer := bytes.NewBuffer([]byte{})
	err = jpeg.Encode(buffer, img, &jpeg.Options{Quality: 80})
	if err != nil {
		return "", "", err
	}

	err = minio.PutObject(ctx, consts.CaptureBucket, capture+".jpg", buffer)
	if err != nil {
		return "", "", err
	}
	url, err = minio.PreSignedObjectUrl(ctx, consts.CaptureBucket, capture+".jpg", time.Hour*1)
	if err != nil {
		return "", "", err
	}
	token = base64.URLEncoding.EncodeToString(frand.New().Bytes(8))
	redis.Redis[consts.AccountPSM].Set(token, capture, time.Minute*3) //验证码3分钟过期

	return url, token, nil
}
