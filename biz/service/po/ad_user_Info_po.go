package po

import (
	"context"
	"github.com/zbwang163/ad_account_server/common/biz_error"
	"github.com/zbwang163/ad_account_server/common/client/mysql"
	"github.com/zbwang163/ad_account_server/common/consts"
	"github.com/zbwang163/ad_account_server/common/logs"
	"gorm.io/gorm"
)

type AdUserInfoPo struct {
	gorm.Model
	CoreUserId int64  `gorm:"column:core_user_id;comment:'用户uid'" json:"core_user_id"`
	UserName   string `gorm:"column:user_name;comment:'昵称'" json:"user_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	AvatarUrl  string `gorm:"column:avatar_url;comment:'头像url'" json:"avatar_url"`
	Slogan     string `gorm:"column:slogan;comment:'个人描述'" json:"slogan"`
	//Career      string `gorm:"column:career;comment:'用户职业'"`
	//Company     string `gorm:"column:company;comment:'用户公司'"`
	//Industry    string `gorm:"column:industry;comment:'用户行业'"`
	//Role        int64  `gorm:"column:role;comment:'用户角色，0:普通用户；1:创意号；2:创作达人；'"`
	//Extra       string `gorm:"column:extra;comment:'额外字段'"`
}

func (po AdUserInfoPo) TableName() string {
	return "ad_user_info"
}

func (po AdUserInfoPo) Db() *gorm.DB {
	return mysql.Db[consts.AccountPSM]
}

func (po AdUserInfoPo) SelectOne(ctx context.Context, where map[string]interface{}) (*AdUserInfoPo, error) {
	var result AdUserInfoPo
	err := po.Db().Where(where).First(&result).Error
	if err != nil {
		logs.CtxError(ctx, biz_error.NewMysqlError(err).Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return &result, nil
}

func (po AdUserInfoPo) UpdateOrInsert(ctx context.Context, where map[string]interface{}, options AdUserInfoPo) (*AdUserInfoPo, error) {
	var result AdUserInfoPo
	err := po.Db().Where(where).Assign(options).FirstOrCreate(&result).Error
	if err != nil {
		logs.CtxError(ctx, biz_error.NewMysqlError(err).Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return &result, nil
}
