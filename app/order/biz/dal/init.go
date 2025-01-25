package dal

import (
	"github.com/xvxiaoman8/gomall/app/order/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
