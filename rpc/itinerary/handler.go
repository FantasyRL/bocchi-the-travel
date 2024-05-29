package main

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/pkg/pack"
	"bocchi/rpc/itinerary/service"
	"context"
)

// ItineraryHandlerImpl implements the last service interface defined in the IDL.
type ItineraryHandlerImpl struct{}

// CreateItinerary implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) CreateItinerary(ctx context.Context, req *itinerary.CreateItineraryRequest) (resp *itinerary.CreateItineraryResponse, err error) {
	resp = new(itinerary.CreateItineraryResponse)
	itineraryResp, err := service.NewItineraryService(ctx).CreateItinerary(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Itinerary = service.BuildItineraryResp(itineraryResp)
	return resp, nil
}

// ShowPartyItinerary implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) ShowPartyItinerary(ctx context.Context, req *itinerary.ShowPartyItineraryRequest) (resp *itinerary.ShowPartyItineraryResponse, err error) {
	resp = new(itinerary.ShowPartyItineraryResponse)
	itineraries, count, err := service.NewItineraryService(ctx).ShowPartyItinerary(req.PartyId)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.Itineraries = service.BuildItinerariesResp(itineraries)
	return resp, nil
}

// ChangeSequence implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) ChangeSequence(ctx context.Context, req *itinerary.ChangeSequenceRequest) (resp *itinerary.ChangeSequenceResponse, err error) {
	resp = new(itinerary.ChangeSequenceResponse)
	err = service.NewItineraryService(ctx).ChangeSequence(req)
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// MergeItinerary implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) MergeItinerary(ctx context.Context, req *itinerary.MergeItineraryRequest) (resp *itinerary.MergeItineraryResponse, err error) {
	resp = new(itinerary.MergeItineraryResponse)
	err = service.NewItineraryService(ctx).MergeItinerary(req)
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// GetItineraryInfo implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) GetItineraryInfo(ctx context.Context, req *itinerary.GetItineraryInfoRequest) (resp *itinerary.GetItineraryInfoResponse, err error) {
	resp = new(itinerary.GetItineraryInfoResponse)
	itineraryResp, err := service.NewItineraryService(ctx).GetItineraryInfo(req.ItineraryId)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Itinerary = service.BuildItineraryResp(itineraryResp)
	return resp, nil
}
