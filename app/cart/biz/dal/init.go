package dal

import (
	"github.com/xvxiaoman8/gomall/app/cart/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
