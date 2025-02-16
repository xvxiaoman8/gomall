package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/xvxiaoman8/gomall/app/product/biz/dal"
	"github.com/xvxiaoman8/gomall/app/product/conf"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"
)

func main() {
	_ = godotenv.Load() //加载环境变量
	// 初始化配置
	dal.Init()

	opts := kitexInit()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

// kitexInit 初始化kitex服务
func kitexInit() (opts []server.Option) {
	// address
	// 解析TCP地址
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	// 将地址添加到opts中
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	// 添加服务基本信息
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// klog
	// 初始化klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	// 初始化异步写入器
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	// 设置klog输出
	klog.SetOutput(asyncWriter)
	// 注册关闭钩子函数
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
