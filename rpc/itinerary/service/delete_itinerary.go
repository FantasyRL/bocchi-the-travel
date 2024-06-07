package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/pkg/errno"
	"bocchi/rpc/itinerary/dal/db"
)

func (s *ItineraryService) DeleteItinerary(req *itinerary.DeleteItineraryRequest) error {
	itineraryResp, err := db.GetItineraryById(s.ctx, req.ItineraryId)
	if err != nil {
		return err
	}
	if itineraryResp.FounderId != req.UserId {
		return errno.NotFounderError
	}
	return db.DeleteItinerary(s.ctx, req.ItineraryId)
}
