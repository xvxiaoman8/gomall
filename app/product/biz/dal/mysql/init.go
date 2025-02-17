package mysql

import (
	"fmt"
	"github.com/xvxiaoman8/gomall/app/product/biz/model"
	"github.com/xvxiaoman8/gomall/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err := DB.AutoMigrate(model.Category{})
	err2 := DB.AutoMigrate(model.Product{})
	//TODO 这里需要处理err和err2 汉字输出
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
}
