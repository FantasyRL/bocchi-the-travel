package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/rpc/itinerary/dal/db"
)

func (s *ItineraryService) ChangeSequence(req *itinerary.ChangeSequenceRequest) error {
	return db.ChangeSequence(s.ctx, req)
}
