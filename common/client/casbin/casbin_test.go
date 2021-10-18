package casbin

import (
	"github.com/zbwang163/ad_account_server/common/client/mysql"
	"github.com/zbwang163/ad_account_server/common/consts"
	"testing"
)

func TestName(t *testing.T) {
	mysql.InitMysql(consts.AccountPSM)
	InitCasbin()

	//a,err := Enforcer.AddRolesForUser("777777",[]string{"user_group_deny_2","user_group_deny_3"},"ad.info.account_server")
	//a,err:= Enforcer.rolepolicy("user_group_deny_1")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(a)
}
