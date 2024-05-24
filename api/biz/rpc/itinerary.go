package rpc

import (
	"bocchi/config"
	"bocchi/kitex_gen/itinerary"
	"bocchi/kitex_gen/itinerary/itineraryhandler"
	"bocchi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitItineraryRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := itineraryhandler.NewClient(
		constants.ItineraryServiceName,
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

	itineraryClient = c
}

func ItineraryCreate(ctx context.Context, req *itinerary.CreateItineraryRequest) (*itinerary.CreateItineraryResponse, error) {
	//rpc client
	resp, err := itineraryClient.CreateItinerary(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ItineraryShow(ctx context.Context, req *itinerary.ShowPartyItineraryRequest) (*itinerary.ShowPartyItineraryResponse, error) {
	//rpc client
	resp, err := itineraryClient.ShowPartyItinerary(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ItinerarySequenceChange(ctx context.Context, req *itinerary.ChangeSequenceRequest) (*itinerary.ChangeSequenceResponse, error) {
	//rpc client
	resp, err := itineraryClient.ChangeSequence(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ItineraryMerge(ctx context.Context, req *itinerary.MergeItineraryRequest) (*itinerary.MergeItineraryResponse, error) {
	//rpc client
	resp, err := itineraryClient.MergeItinerary(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil
}
