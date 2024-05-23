package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/rpc/itinerary/dal/db"
)

func (s *ItineraryService) MergeItinerary(req *itinerary.MergeItineraryRequest) error {
	return db.ChangeItineraryStatus(s.ctx, req.PartyId, req.ItineraryId, 1)
}
