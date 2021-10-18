package http_adapter

import (
	"github.com/gin-gonic/gin"
	casbinAdapter "github.com/zbwang163/ad_account_server/common/client/casbin"
	"github.com/zbwang163/ad_account_server/common/consts"
	"github.com/zbwang163/ad_account_server/common/convert"
	"github.com/zbwang163/ad_account_server/common/logs"
	"github.com/zbwang163/ad_account_server/common/utils"
	"net/http"
)

func UserInfoMiddleware(c *gin.Context) {
	deviceIdStr := c.GetHeader("device_id")
	deviceId, _ := convert.StringToInt64(deviceIdStr)
	c.Set(consts.CtxDeviceId, deviceId)

	sessionId, _ := c.Cookie("session_id")
	c.Set(consts.CtxUserId, utils.GetCoreUserIdFromSession(sessionId))
	c.Next()

	ip, _ := c.RemoteIP()
	if ip != nil {
		c.Set(consts.CtxIp, ip.String())
	}
}

func LogIdMiddleware(c *gin.Context) {
	logId := utils.GenerateLogId()
	c.Set(consts.LogId, logId)
	c.Next()
}

func ResponseMiddleware(c *gin.Context) {
	c.Header(consts.LogId, utils.GetCtxLogId(c))
	c.Next()
}

func UserPrivilegeManagement(c *gin.Context) {
	ctxUserId := c.Value(consts.CtxUserId).(int64)
	url := c.Request.RequestURI
	method := c.Request.Method
	res, err := casbinAdapter.Enforcer.Enforce(convert.Int64ToString(ctxUserId), consts.AccountPSM, url, method)
	if err != nil {
		logs.CtxError(c, "UserPrivilegeManagement err:%v", err)
	}
	if !res {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "message": "未授权"})
		c.Abort()
	}
	c.Next()
}
