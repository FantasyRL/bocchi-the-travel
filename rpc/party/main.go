package main

import (
	"bocchi/api/biz/rpc"
	"bocchi/config"
	party "bocchi/kitex_gen/party/partyhandler"
	"bocchi/pkg/constants"
	"bocchi/pkg/utils"
	"bocchi/pkg/utils/eslogrus"
	"bocchi/rpc/party/dal"
	"crypto/tls"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/netpoll"
	elastic "github.com/elastic/go-elasticsearch/v8"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	listenAddr string
	EsClient   *elastic.Client
)

func Init() {
	config.Init(constants.PartyServiceName)
	InitEs()
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))

	dal.Init()
	rpc.InitUserRPC()
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

	partyHandlerImpl := new(PartyHandlerImpl)
	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	svr := party.NewServer(partyHandlerImpl, // 指定 Registry 与服务基本信息
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

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(EsClient, config.ElasticSearch.Host, logrus.DebugLevel, constants.ElasticSearchIndexName)
	if err != nil {
		klog.Warn(err)
	}

	return hook
}

func InitEs() {
	esConn := fmt.Sprintf("http://%s", config.ElasticSearch.Addr)
	cfg := elastic.Config{
		Addresses: []string{esConn},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	client, err := elastic.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	EsClient = client
}
