package dal

import (
	"github.com/xvxiaoman8/gomall/app/product/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
