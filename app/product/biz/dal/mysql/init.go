// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//package mysql
//
//import (
//	"fmt"
//	"os"
//
//	"github.com/xvxiaoman8/gomall/common/mtl"
//
//	"github.com/xvxiaoman8/gomall/app/product/conf"
//	"github.com/xvxiaoman8/gomall/app/product/model"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/plugin/opentelemetry/tracing"
//)
//
// 定义全局变量DB和err，分别用于存储GORM数据库连接实例和错误信息
//var (
//	DB  *gorm.DB
//	err error
//)
//
//// Init 函数用于初始化MySQL数据库连接，设置自动迁移，并在非线上环境下插入示例数据
//func Init() {
//	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
//	DB, err = gorm.Open(mysql.Open(dsn),
//		&gorm.Config{
//			PrepareStmt:            true,
//			SkipDefaultTransaction: true,
//		},
//	)
//	if err != nil {
//		panic(err)
//	}
//	if os.Getenv("GO_ENV") != "online" {
//		// 检查Product表是否存在，如果不存在则需要插入示例数据
//		needDemoData := !DB.Migrator().HasTable(&model.Product{})
//		// 自动迁移数据库表结构
//		DB.AutoMigrate( //nolint:errcheck
//			&model.Product{},
//			&model.Category{},
//		)
//		if needDemoData {
//			// 插入 示例数据 到Category表
//			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')")
//			// 插入示例数据到Product表
//			DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
//			// 插入示例数据到Product_Category关联表
//			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
//		}
//	}
//	// 使用OpenTelemetry插件进行链路追踪，不包含指标数据
//	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
//		panic(err)
//	}
//}

//package mysql
//
//import (
//	"fmt"
//	"os"
//
//	"github.com/xvxiaoman8/gomall/app/product/conf"
//	"github.com/xvxiaoman8/gomall/app/product/model"
//	"github.com/xvxiaoman8/gomall/common/mtl"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/plugin/opentelemetry/tracing"
//)
//
//var (
//	DB  *gorm.DB
//	err error
//)
//
//// Init 函数用于初始化MySQL数据库连接，设置自动迁移，并在非线上环境下插入示例数据
//func Init() (*gorm.DB, error) {
//	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
//	db, err := gorm.Open(mysql.Open(dsn),
//		&gorm.Config{
//			PrepareStmt:            true,
//			SkipDefaultTransaction: true,
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	if os.Getenv("GO_ENV") != "online" {
//		// 检查Product表是否存在，如果不存在则需要插入示例数据
//		needDemoData := !db.Migrator().HasTable(&model.Product{})
//		// 自动迁移数据库表结构
//		if err := db.AutoMigrate(&model.Product{}, &model.Category{}); err != nil {
//			return nil, err
//		}
//		if needDemoData {
//			// 插入示例数据到Category表
//			if result := db.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')"); result.Error != nil {
//				return nil, result.Error
//			}
//			// 插入示例数据到Product表
//			if result := db.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )"); result.Error != nil {
//				return nil, result.Error
//			}
//			// 插入示例数据到Product_Category关联表
//			if result := db.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )"); result.Error != nil {
//				return nil, result.Error
//			}
//		}
//	}
//	// 使用OpenTelemetry插件进行链路追踪，不包含指标数据
//	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
//		return nil, err
//	}
//	return db, nil
//}

package mysql

import (
	"fmt"
	model2 "github.com/xvxiaoman8/gomall/app/product/biz/model"
	"os"

	"github.com/xvxiaoman8/gomall/common/mtl"

	"github.com/xvxiaoman8/gomall/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	//获取数据库连接所需的信息
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//创建db对象，承载连接
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		fmt.Println("连接数据库失败：", err.Error())
	} //抛错

	//检查是否是线上环境
	if os.Getenv("GO_ENV") != "online" {
		//检查表是否存在
		needDemoData := !DB.Migrator().HasTable(&model2.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model2.Product{},
			&model2.Category{},
		) //自动迁移数据库表结构
		if needDemoData { //初始化的时候尝试插入的数据
			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')")
			DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
		}
	}
	// 使用OpenTelemetry插件进行链路追踪，不包含指标数据
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		fmt.Println("链路追踪失败：", err.Error())
	}
}
