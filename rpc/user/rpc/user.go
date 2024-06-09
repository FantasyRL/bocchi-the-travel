package rpc

import (
	"bocchi/config"
	"bocchi/kitex_gen/trust"
	"bocchi/kitex_gen/trust/trusthandler"
	"bocchi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var trustClient trusthandler.Client

func InitTrustRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	c, err := trusthandler.NewClient(
		constants.TrustServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	trustClient = c
}

func IsTrust(ctx context.Context, req *trust.IsTrustRequest) (*trust.IsTrustResponse, error) {
	resp, err := trustClient.IsTrust(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
