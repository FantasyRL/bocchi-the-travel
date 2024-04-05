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

func InitPartyRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

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

func PartyCreate(ctx context.Context, req *party.CreatePartyRequest) (*party.CreatePartyResponse, error) {
	//rpc client
	resp, err := partyClient.CreateParty(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PartyJoin(ctx context.Context, req *party.JoinPartyRequest) (*party.JoinPartyResponse, error) {
	//rpc client
	resp, err := partyClient.JoinParty(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PartyApplyList(ctx context.Context, req *party.ApplyListRequest) (*party.ApplyListResponse, error) {
	//rpc client
	resp, err := partyClient.ApplyList(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PartyPermitJoin(ctx context.Context, req *party.PermitJoinRequest) (*party.PermitJoinResponse, error) {
	//rpc client
	resp, err := partyClient.PermitJoin(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PartyMembersList(ctx context.Context, req *party.GetPartyMembersRequest) (*party.GetPartyMembersResponse, error) {
	//rpc client
	resp, err := partyClient.GetPartyMembers(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PartySearch(ctx context.Context, req *party.SearchPartyRequest) (*party.SearchPartyResponse, error) {
	//rpc client
	resp, err := partyClient.SearchParty(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}
