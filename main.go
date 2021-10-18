package main

import (
	"github.com/gin-gonic/gin"
	casbinAdapter "github.com/zbwang163/ad_account_server/common/client/casbin"
	"github.com/zbwang163/ad_account_server/common/client/minio"
	"github.com/zbwang163/ad_account_server/common/client/mysql"
	"github.com/zbwang163/ad_account_server/common/client/redis"
	"github.com/zbwang163/ad_account_server/common/consts"
)

func main() {
	//lis, err := net.Listen("tcp", ":50001")
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//accountRpc.RegisterAccountServiceServer(s, NewServer())
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}

	r := gin.Default()
	InitClients()
	Register(r)
	r.Run() //8080端口
}

func InitClients() {
	minio.InitMinIO()
	mysql.InitMysql(consts.AccountPSM)
	redis.InitRedis(consts.AccountPSM)
	casbinAdapter.InitCasbin()
}
