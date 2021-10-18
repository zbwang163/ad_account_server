package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	_ "github.com/go-sql-driver/mysql"
	"github.com/weblazy/snowflake"
	"github.com/zbwang163/ad_account_server/common/client/redis"
	"github.com/zbwang163/ad_account_server/common/consts"
	"github.com/zbwang163/ad_base_overpass"
	"github.com/zbwang163/ad_common/biz_error"
	"net"
	"time"
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

func GetCoreUserIdFromSession(sessionId string) int64 {
	if sessionId == "" {
		return 0
	}
	str := redis.Redis[consts.AccountPSM].Get(sessionId).Val()
	var s struct {
		CoreUserId int64 `json:"core_user_id"`
	}
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		return 0
	}
	return s.CoreUserId
}

// GetLocalIp 获取本机的io
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GenerateLogId 生成log id
func GenerateLogId() string {
	t := time.Now().Format("200601021504")
	bytes, _ := uuid.GenerateRandomBytes(11)
	return fmt.Sprintf("%v%x", t, bytes)
}

// GetCtxLogId 从context中获取log id
func GetCtxLogId(c *gin.Context) string {
	if logId, ok := c.Value(consts.LogId).(string); ok {
		return logId
	}
	return ""
}

func GenGlobalUniqueId() (int64, error) {
	worker, err := snowflake.NewWorker(1)
	if err != nil {
		return 0, err
	}
	return worker.GetId(), nil
}

func IfUserHavePrivilege(ctx context.Context, coreUserId int64, url string, method string) bool {
	return false
}
