package mysql

import (
	"fmt"
	"os"

	"github.com/xvxiaoman8/gomall/app/user/biz/model"
	"github.com/xvxiaoman8/gomall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 从文件中读取环境变量
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSOL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	// 设置数据库自动迁移
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
