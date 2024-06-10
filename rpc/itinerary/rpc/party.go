package rpc

import (
	"bocchi/config"
	"bocchi/kitex_gen/party"
	"bocchi/kitex_gen/party/partyhandler"
	"bocchi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var partyClient partyhandler.Client

func InitPartyRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	c, err := partyhandler.NewClient(
		constants.PartyServiceName,
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

	partyClient = c
}

func IsMemberExist(ctx context.Context, req *party.IsMemberExistRequest) (*party.IsMemberExistResponse, error) {
	resp, err := partyClient.IsMemberExist(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func IsMemberAdmin(ctx context.Context, req *party.IsMemberAdminRequest) (*party.IsMemberAdminResponse, error) {
	resp, err := partyClient.IsMemberAdmin(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
