//package main
//
//import (
//	"github.com/cloudwego/kitex/pkg/klog"
//	"github.com/cloudwego/kitex/server"
//	"github.com/joho/godotenv"
//	"github.com/xvxiaoman8/gomall/app/product/biz/dal"
//	"github.com/xvxiaoman8/gomall/app/product/conf"
//	"github.com/xvxiaoman8/gomall/common/mtl"
//	"github.com/xvxiaoman8/gomall/common/serversuite"
//	"github.com/xvxiaoman8/gomall/common/utils"
//	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
//	"gopkg.in/natefinch/lumberjack.v2"
//	"net"
//	"strings"
//)
//
//var serviceName = conf.GetConf().Kitex.Service
//
//func main() {
//	_ = godotenv.Load() //加载环境变量
//	mtl.InitLog(&lumberjack.Logger{
//		Filename:   conf.GetConf().Kitex.LogFileName,
//		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
//		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
//		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
//	})
//	mtl.InitTracing(serviceName)
//	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
//
//	// 初始化配置
//	dal.Init()
//
//	opts := kitexInit()
//
//	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)
//
//	err := svr.Run()
//	if err != nil {
//		klog.Error(err.Error())
//	}
//}
//func kitexInit() (opts []server.Option) {
//	// address
//	address := conf.GetConf().Kitex.Address
//	if strings.HasPrefix(address, ":") {
//		localIp := utils.MustGetLocalIPv4()
//		address = localIp + address
//	}
//	addr, err := net.ResolveTCPAddr("tcp", address)
//	if err != nil {
//		panic(err)
//	}
//
//	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName, RegistryAddr: conf.GetConf().Registry.RegistryAddress[0]}))
//	return
//}

package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/xvxiaoman8/gomall/app/product/biz/dal"
	"github.com/xvxiaoman8/gomall/app/product/conf"
	"github.com/xvxiaoman8/gomall/common/mtl"
	"github.com/xvxiaoman8/gomall/common/serversuite"
	"github.com/xvxiaoman8/gomall/common/utils"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"strings"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load() //加载环境变量
	klog.Infof("环境变量加载完成")

	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	klog.Infof("日志初始化完成")

	mtl.InitTracing(serviceName)
	klog.Infof("追踪初始化完成")

	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	klog.Infof("指标初始化完成")

	// 初始化配置
	dal.Init()
	klog.Infof("数据库和Redis初始化完成")

	opts := kitexInit()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Errorf("服务启动失败: %v", err)
	}
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

//// kitexInit 初始化kitex服务
//func kitexInit() (opts []server.Option) {
//	// address
//	// 解析TCP地址
//	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
//	if err != nil {
//		panic(err)
//	}
//	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
//	if err != nil {
//		fmt.Println("cannot connect to consul", err)
//		klog.Fatal(err)
//	}
//	// 将地址添加到opts中
//	opts = append(opts, server.WithServiceAddr(addr))
//
//	// service info
//	// 添加服务基本信息
//	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
//		ServiceName: conf.GetConf().Kitex.Service,
//	}), server.WithRegistry(r))
//
//	// klog
//	// 初始化klog
//	logger := kitexlogrus.NewLogger()
//	klog.SetLogger(logger)
//	klog.SetLevel(conf.LogLevel())
//	// 初始化异步写入器
//	asyncWriter := &zapcore.BufferedWriteSyncer{
//		WS: zapcore.AddSync(&lumberjack.Logger{
//			Filename:   conf.GetConf().Kitex.LogFileName,
//			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
//			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
//			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
//		}),
//		FlushInterval: time.Minute,
//	}
//	// 设置klog输出
//	klog.SetOutput(asyncWriter)
//	// 注册关闭钩子函数
//	server.RegisterShutdownHook(func() {
//		asyncWriter.Sync()
//	})
//	return
//}
