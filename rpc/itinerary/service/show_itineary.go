package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/rpc/itinerary/dal/db"
)

func (s *ItineraryService) ShowPartyItinerary(partyId int64) (*[]db.Itinerary, int64, error) {
	return db.ShowPartyItinerary(s.ctx, partyId)
}

func (s *ItineraryService) ShowPartyItineraryDraft(partyId int64) (*[]db.Itinerary, int64, error) {
	return db.ShowPartyItineraryDraft(s.ctx, partyId)
}

func (s *ItineraryService) GetItineraryInfo(itineraryId int64) (*db.Itinerary, error) {
	return db.GetItineraryById(s.ctx, itineraryId)
}

func (s *ItineraryService) GetMyItineraries(req *itinerary.GetMyItinerariesRequest) (*[]db.Itinerary, int64, error) {
	return db.GetItinerariesByUidAndPartyId(s.ctx, req.UserId, req.PartyId)
}
