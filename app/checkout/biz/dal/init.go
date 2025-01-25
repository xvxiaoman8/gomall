package dal

import (
	"github.com/xvxiaoman8/gomall/app/checkout/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
