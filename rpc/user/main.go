package main

import (
	"bocchi/config"
	user "bocchi/kitex_gen/user/userhandler"
	"bocchi/pkg/constants"
	"bocchi/pkg/utils"
	"bocchi/rpc/user/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/netpoll"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var listenAddr string

func Init() {
	config.Init(constants.UserServiceName)
	dal.Init()

}

func main() {
	Init()
	//注册到etcd
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		klog.Fatal(err)
	}

	//获取addr
	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.AddrList)-1 {
			klog.Fatal("not available addr")
		}
	}

	userHandlerImpl := new(UserHandlerImpl)
	userCli, err := NewUserClient(listenAddr)
	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	userHandlerImpl.userCli = userCli
	//然而不使用WithServiceAddr方法的话，server还是在监听8888
	//那Impl携带一个Client就没用了

	svr := user.NewServer(userHandlerImpl, // 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServiceAddr(serviceAddr),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: constants.UserServiceName,
			}),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}))

	err = svr.Run()

	if err != nil {
		klog.Error(err.Error())
	}
}
