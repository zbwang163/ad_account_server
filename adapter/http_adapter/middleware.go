package http_adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/zbwang163/ad_account_server/common/consts"
	"github.com/zbwang163/ad_account_server/common/convert"
	"github.com/zbwang163/ad_account_server/common/utils"
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
