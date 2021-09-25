package logs

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/zbwang163/ad_common/env"
	"os"
	"runtime"
)

var defaultLogger *logrus.Entry

func init() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{}) //log日志序列化为json
	log.Out = os.Stdout
	//file, err := os.OpenFile(fmt.Sprintf("/var/log/ad_platform_info/%v.log", time_utils.Time20060102_15(time.Now())),
	//	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	log.Out = file
	//} else {
	//	log.Info("Failed to log to file, using default stderr")
	//}

	ctxLog := log.WithFields(logrus.Fields{
		"ip":       env.GetPodIp(),
		"pod_name": env.GetPodName(),
	})
	defaultLogger = ctxLog
}

func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.WithField("log_id", ctx.Value("log_id")).Infof(format, args)
}

func CtxError(ctx context.Context, format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	defaultLogger.WithFields(logrus.Fields{
		"log_id": ctx.Value("log_id"), //log id
		"file":   file,                //报错的文件
		"line":   line,                // 报错行
		"func":   f.Name(),            //报错的函数名
	}).Errorf(format, args)
}
