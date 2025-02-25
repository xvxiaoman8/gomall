package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/xvxiaoman8/gomall/app/payment/biz/dal"
	"github.com/xvxiaoman8/gomall/app/payment/conf"
	"github.com/xvxiaoman8/gomall/common/mtl"
	"github.com/xvxiaoman8/gomall/common/serversuite"
	"github.com/xvxiaoman8/gomall/common/utils"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"strings"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovered:", r)
		}
	}()
	_ = godotenv.Load()
	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	mtl.InitTracing(serviceName)
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	dal.Init()

	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println("payment server started", err)

}
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

//func kitexInit() (opts []server.Option) {
//	// address
//	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
//	if err != nil {
//		panic(err)
//	}
//	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
//	if err != nil {
//		fmt.Println("cannot connect to consul", err)
//		klog.Fatal(err)
//	}
//	opts = append(opts, server.WithServiceAddr(addr))
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
