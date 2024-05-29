package service

import "bocchi/rpc/itinerary/dal/db"

func (s *ItineraryService) ShowPartyItinerary(partyId int64) (*[]db.Itinerary, int64, error) {
	return db.ShowPartyItinerary(s.ctx, partyId)
}

func (s *ItineraryService) GetItineraryInfo(itineraryId int64) (*db.Itinerary, error) {
	return db.GetItineraryById(s.ctx, itineraryId)
}
