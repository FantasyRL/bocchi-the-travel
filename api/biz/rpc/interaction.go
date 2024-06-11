package rpc

import (
	"bocchi/config"
	"bocchi/kitex_gen/interaction"
	"bocchi/kitex_gen/interaction/interactionhandler"
	"bocchi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitInteractionRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := interactionhandler.NewClient(
		constants.InteractionServiceName,
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

	interactionClient = c
}

func CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest) (*interaction.CommentCreateResponse, error) {
	resp, err := interactionClient.CommentCreate(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest) (*interaction.CommentDeleteResponse, error) {
	resp, err := interactionClient.CommentDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *interaction.CommentListRequest) (*interaction.CommentListResponse, error) {
	resp, err := interactionClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
