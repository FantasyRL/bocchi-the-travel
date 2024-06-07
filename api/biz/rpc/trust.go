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

func InitTrustRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

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

func FollowAction(ctx context.Context, req *trust.FollowActionRequest) (*trust.FollowActionResponse, error) {
	resp, err := trustClient.TrustAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FollowerList(ctx context.Context, req *trust.FollowerListRequest) (*trust.FollowerListResponse, error) {
	resp, err := trustClient.FollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FollowingList(ctx context.Context, req *trust.FollowingListRequest) (*trust.FollowingListResponse, error) {
	resp, err := trustClient.FollowingList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FriendList(ctx context.Context, req *trust.FriendListRequest) (*trust.FriendListResponse, error) {
	resp, err := trustClient.TrustEachList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
