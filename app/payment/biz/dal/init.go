package dal

import (
	"github.com/xvxiaoman8/gomall/app/payment/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
