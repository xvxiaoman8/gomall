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

//package redis
//
//import (
//	"context"
//
//	"github.com/cloudwego/kitex/pkg/klog"
//	"github.com/redis/go-redis/extra/redisotel/v9"
//	"github.com/redis/go-redis/extra/redisprometheus/v9"
//	"github.com/redis/go-redis/v9"
//	"github.com/xvxiaoman8/gomall/app/product/conf"
//	"github.com/xvxiaoman8/gomall/common/mtl"
//)
//
//var RedisClient *redis.Client
//
//func Init() {
//	RedisClient = redis.NewClient(&redis.Options{
//		Addr:     conf.GetConf().Redis.Address,
//		Username: conf.GetConf().Redis.Username,
//		Password: conf.GetConf().Redis.Password,
//		DB:       conf.GetConf().Redis.DB,
//	})
//	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
//		panic(err)
//	}
//	if err := redisotel.InstrumentTracing(RedisClient); err != nil {
//		klog.Error("redis tracing collect error ", err)
//	}
//	if err := mtl.Registry.Register(redisprometheus.NewCollector("default", "product", RedisClient)); err != nil {
//		klog.Error("redis metric collect error ", err)
//	}
//	redisotel.InstrumentTracing(RedisClient) //nolint:errcheck
//}

package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/xvxiaoman8/gomall/app/product/conf"
)

var RedisClient *redis.Client

// 初始化函数
//func Init() {
//	// 创建Redis客户端
//	RedisClient = redis.NewClient(&redis.Options{
//		Addr:     conf.GetConf().Redis.Address,  // Redis地址
//		Username: conf.GetConf().Redis.Username, // Redis用户名
//		Password: conf.GetConf().Redis.Password, // Redis密码
//		DB:       conf.GetConf().Redis.DB,       // Redis数据库
//	})
//	// 检查Redis客户端是否连接成功
//	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
//		fmt.Println("客户端连接失败：", err.Error())
//	}
//	// 注册Redis追踪
//	if err := redisotel.InstrumentTracing(RedisClient); err != nil {
//		klog.Error("redis追踪失败 ", err)
//	}
//	// 注册Redis指标收集
//	if err := mtl.Registry.Register(redisprometheus.NewCollector("default", "product", RedisClient)); err != nil {
//		klog.Error("redis指标收集失败 ", err)
//	}
//	// 注册Redis追踪
//	redisotel.InstrumentTracing(RedisClient) //nolint:errcheck
//}

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
