package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zbwang163/ad_account_server/common/client/mysql"
	"github.com/zbwang163/ad_account_server/common/consts"
	"time"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	a, err := gormadapter.NewAdapterByDBWithCustomTable(mysql.Db[consts.AccountPSM], &ADCasbinRule{}, "ad_role_policy")
	if err != nil {
		panic(fmt.Sprintf("cahBin init gorm adapter err:%v", err))
	}
	Enforcer, _ = casbin.NewEnforcer("config/rbac_with_pattern_model.conf", a)

	Enforcer.AddFunction("expMatch", ExpFunc)
	Enforcer.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch2)
	err = Enforcer.LoadPolicy()
	if err != nil {
		panic(fmt.Sprintf("cahBin load policy err:%v", err))
	}
}

func ExpFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	return bool(ExpMatch(name1)), nil
}

func ExpMatch(key1 string) bool {
	if key1 == "-1" {
		return true
	}
	now := time.Now().Format("2006-01-02 15:04")
	return now <= key1
}
