// Code generated by Kitex v0.9.1. DO NOT EDIT.

package itineraryhandler

import (
	itinerary "bocchi/kitex_gen/itinerary"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateItinerary(ctx context.Context, req *itinerary.CreateItineraryRequest, callOptions ...callopt.Option) (r *itinerary.CreateItineraryResponse, err error)
	GetItineraryInfo(ctx context.Context, req *itinerary.GetItineraryInfoRequest, callOptions ...callopt.Option) (r *itinerary.GetItineraryInfoResponse, err error)
	ShowPartyItinerary(ctx context.Context, req *itinerary.ShowPartyItineraryRequest, callOptions ...callopt.Option) (r *itinerary.ShowPartyItineraryResponse, err error)
	ChangeSequence(ctx context.Context, req *itinerary.ChangeSequenceRequest, callOptions ...callopt.Option) (r *itinerary.ChangeSequenceResponse, err error)
	MergeItinerary(ctx context.Context, req *itinerary.MergeItineraryRequest, callOptions ...callopt.Option) (r *itinerary.MergeItineraryResponse, err error)
	GetMyItineraries(ctx context.Context, req *itinerary.GetMyItinerariesRequest, callOptions ...callopt.Option) (r *itinerary.GetMyItinerariesResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kItineraryHandlerClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kItineraryHandlerClient struct {
	*kClient
}

func (p *kItineraryHandlerClient) CreateItinerary(ctx context.Context, req *itinerary.CreateItineraryRequest, callOptions ...callopt.Option) (r *itinerary.CreateItineraryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateItinerary(ctx, req)
}

func (p *kItineraryHandlerClient) GetItineraryInfo(ctx context.Context, req *itinerary.GetItineraryInfoRequest, callOptions ...callopt.Option) (r *itinerary.GetItineraryInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetItineraryInfo(ctx, req)
}

func (p *kItineraryHandlerClient) ShowPartyItinerary(ctx context.Context, req *itinerary.ShowPartyItineraryRequest, callOptions ...callopt.Option) (r *itinerary.ShowPartyItineraryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowPartyItinerary(ctx, req)
}

func (p *kItineraryHandlerClient) ChangeSequence(ctx context.Context, req *itinerary.ChangeSequenceRequest, callOptions ...callopt.Option) (r *itinerary.ChangeSequenceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeSequence(ctx, req)
}

func (p *kItineraryHandlerClient) MergeItinerary(ctx context.Context, req *itinerary.MergeItineraryRequest, callOptions ...callopt.Option) (r *itinerary.MergeItineraryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MergeItinerary(ctx, req)
}

func (p *kItineraryHandlerClient) GetMyItineraries(ctx context.Context, req *itinerary.GetMyItinerariesRequest, callOptions ...callopt.Option) (r *itinerary.GetMyItinerariesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMyItineraries(ctx, req)
}
