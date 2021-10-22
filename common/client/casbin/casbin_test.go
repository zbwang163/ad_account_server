package casbin

import (
	"github.com/zbwang163/ad_account_server/common/client/mysql"
	"github.com/zbwang163/ad_account_server/common/consts"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	mysql.InitMysql(consts.AccountPSM)
	InitCasbin()
	//Enforcer.Enforce(111,"ad.info.account_server")
	//a,err := Enforcer.AddRolesForUser("",[]string{"user_group_deny_2","user_group_deny_3"},"ad.info.account_server")
	Enforcer.AddPolicy("login_user", "ad_info_platform", "article_group", "GET", "allow", time.Now().Add(time.Hour*24*365).Format("2006-01-02 15:04:05"))
	Enforcer.AddPolicy("bob", "ad_info_platform", "/ad_info_platform/article/1", "GET", "deny", time.Now().Add(time.Hour*10).Format("2006-01-02 15:04:05"))
	Enforcer.AddRoleForUser("alice", "login_user", "ad_info_platform")
	Enforcer.AddRoleForUser("bob", "login_user", "ad_info_platform")
	Enforcer.AddNamedGroupingPolicy("g2", "/ad_info_platform/article/*", "article_group")

	//a,err:= Enforcer.rolepolicy("user_group_deny_1")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(a)
}

func TestUserPermission(t *testing.T) {
	mysql.InitMysql(consts.AccountPSM)
	InitCasbin()
	res, err := Enforcer.Enforce("alice", "ad_info_platform", "/ad_info_platform/article/1", "GET")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("1---", res)

	res1, err := Enforcer.Enforce("bob", "ad_info_platform", "/ad_info_platform/article/1", "GET")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("2---", res1)
}
