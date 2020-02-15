package main

import (
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	rl "github.com/juju/ratelimit"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"

	userRpcConfig "micro-im/application/userserver/cmd/config"
	"micro-im/application/userserver/models"
	userpb "micro-im/application/userserver/protos"
	"micro-im/application/userserver/rpcserverimpl"
)

func main() {
	userRpcFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_rpc.json",
		Usage: "please use xxx -f config_rpc.json",
	}
	configFile := flag.String(userRpcFlag.Name, userRpcFlag.Value, userRpcFlag.Usage)
	flag.Parse()
	conf := new(userRpcConfig.RpcConfig)

	if err := config.LoadFile(*configFile); err != nil {
		log.Fatal(err)
	}
	if err := config.Scan(conf); err != nil {
		log.Fatal(err)
	}
	engineUser, err := xorm.NewEngine(conf.Engine.Name, conf.Engine.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	etcdRegisty := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Address
		});
	b := rl.NewBucketWithRate(float64(conf.Server.RateLimit), int64(conf.Server.RateLimit))
	service := micro.NewService(
		micro.Name(conf.Server.Name),
		micro.Registry(etcdRegisty),
		micro.Version(conf.Version),
		micro.Transport(grpc.NewTransport()),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(b, false)),
		micro.Flags(userRpcFlag),
	)
	service.Init()
	userModel := models.NewMembersModel(engineUser)
	userRpcServer := rpcserverimpl.NewUserRpcServer(userModel)
	if err := userpb.RegisterUserHandler(service.Server(), userRpcServer); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
