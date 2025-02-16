package main

import (
	"github.com/v2pro/plz/countlog/output/lumberjack"
	"github.com/xvxiaoman8/gomall/app/auth/biz/middleware"
	"net"
	"path/filepath"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/xvxiaoman8/gomall/app/user/conf"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/auth/authservice"
	"go.uber.org/zap/zapcore"
)

func main() {
	opts := kitexInit()

	// 初始化 Casbin
	enforcer := initCasbin()

	// 创建中间件链

	opts = append(opts, server.WithMiddleware(
		middleware.ChainMiddleware(
			middleware.JWTAuthMiddleware,
			middleware.CasbinMiddleware(enforcer),
		),
	))

	svr := authservice.NewServer(new(AuthServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

func initCasbin() *casbin.Enforcer {
	modelPath := filepath.Join("conf", "model.conf")
	policyPath := filepath.Join("conf", "policy.csv")

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}
	return enforcer
}
