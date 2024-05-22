package main

import (
	"bocchi/kitex_gen/itinerary"
	"context"
)

// ItineraryHandlerImpl implements the last service interface defined in the IDL.
type ItineraryHandlerImpl struct{}

// CreateItinerary implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) CreateItinerary(ctx context.Context, req *itinerary.CreateItineraryRequest) (resp *itinerary.CreateItineraryResponse, err error) {
	// TODO: Your code here...
	return
}

// ShowPartyItinerary implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) ShowPartyItinerary(ctx context.Context, req *itinerary.ShowPartyItineraryRequest) (resp *itinerary.ShowPartyItineraryResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeSequence implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) ChangeSequence(ctx context.Context, req *itinerary.ChangeSequenceRequest) (resp *itinerary.ChangeSequenceResponse, err error) {
	// TODO: Your code here...
	return
}

// MergeItinerary implements the ItineraryHandlerImpl interface.
func (s *ItineraryHandlerImpl) MergeItinerary(ctx context.Context, req *itinerary.MergeItineraryRequest) (resp *itinerary.MergeItineraryResponse, err error) {
	// TODO: Your code here...
	return
}
