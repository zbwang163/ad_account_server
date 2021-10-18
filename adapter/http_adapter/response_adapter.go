package http_adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/zbwang163/ad_account_server/biz/controller"
	"github.com/zbwang163/ad_account_server/biz/service/query"
	"github.com/zbwang163/ad_account_server/common/biz_error"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var account = controller.NewAccountController()

type Handler func(*gin.Context) (interface{}, error)

func HandlerFunc(f Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		dto, err := f(c)
		if err == nil {
			c.JSON(http.StatusOK, Response{0, "success", dto})
		} else if bizErr, ok := err.(*biz_error.BizError); ok {
			c.JSON(http.StatusOK, Response{bizErr.GetErrorCode(), bizErr.GetErrorType() + ":" + bizErr.GetBizErrorMessage(), nil})
		} else {
			temp := biz_error.NewInternalError(err)
			c.JSON(http.StatusOK, Response{temp.GetErrorCode(), temp.GetErrorType() + ":" + temp.GetBizErrorMessage(), nil})
		}
	}

}

func Login(c *gin.Context) (interface{}, error) {
	param := query.LoginQuery{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return account.Login(c, &param)
}

func GetCaptureImage(c *gin.Context) (interface{}, error) {
	return account.GetCaptureImage(c)
}

func Register(c *gin.Context) (interface{}, error) {
	param := query.RegisterQuery{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return account.Register(c, &param)
}

func SendEmailCapture(c *gin.Context) (interface{}, error) {
	param := query.SendEmailCapture{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return nil, account.SendEmailCapture(c, &param)
}

func GetUserInfo(c *gin.Context) (interface{}, error) {
	param := query.UserInfoQuery{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return account.GetUserInfo(c, &param)
}

func UpdateUserInfo(c *gin.Context) (interface{}, error) {
	param := query.UpdateUserInfoQuery{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return account.UpdateUserInfo(c, &param)
}

func AddPolicy(c *gin.Context) (interface{}, error) {
	param := query.PolicyQuery{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return nil, account.AddPolicyRule(c, &param)
}

func RemovePolicy(c *gin.Context) (interface{}, error) {
	param := query.PolicyQuery{}
	if err := c.ShouldBind(&param); err != nil {
		return nil, biz_error.NewParamError(err)
	}
	return nil, account.RemovePolicyRule(c, &param)
}
