package dal

import (
	"github.com/xvxiaoman8/gomall/app/frontend/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
