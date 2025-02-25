package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/xvxiaoman8/gomall/app/checkout/biz/dal"
	"github.com/xvxiaoman8/gomall/app/checkout/conf"
	"github.com/xvxiaoman8/gomall/app/checkout/infra/rpc"
	"github.com/xvxiaoman8/gomall/common/mtl"
	"github.com/xvxiaoman8/gomall/common/serversuite"
	"github.com/xvxiaoman8/gomall/common/utils"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"strings"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()

	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	mtl.InitTracing(serviceName)
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	//mq.NewConnCh()
	//defer mq.ConnClose()
	dal.Init()
	rpc.InitClient()
	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

//
//func kitexInit() (opts []server.Option) {
//	// address
//	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
//	if err != nil {
//		panic(err)
//	}
//
//	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
//	if err != nil {
//		klog.Fatal(err)
//	}
//	opts = append(opts, server.WithServiceAddr(addr))
//
//	// service info
//	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
//		ServiceName: conf.GetConf().Kitex.Service,
//	}), server.WithRegistry(r))
//
//	// klog
//	logger := kitexlogrus.NewLogger()
//	klog.SetLogger(logger)
//	klog.SetLevel(conf.LogLevel())
//	asyncWriter := &zapcore.BufferedWriteSyncer{
//		WS: zapcore.AddSync(&lumberjack.Logger{
//			Filename:   conf.GetConf().Kitex.LogFileName,
//			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
//			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
//			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
//		}),
//		FlushInterval: time.Minute,
//	}
//	klog.SetOutput(asyncWriter)
//	server.RegisterShutdownHook(func() {
//		asyncWriter.Sync()
//	})
//	return
//}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}

	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName, RegistryAddr: conf.GetConf().Registry.RegistryAddress[0]}))
	return
}
