package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/rpc/itinerary/dal/db"
)

func (s *ItineraryService) CreateItinerary(req *itinerary.CreateItineraryRequest) (*db.Itinerary, error) {
	itineraryModel := &db.Itinerary{
		Title:             req.Title,
		FounderId:         req.FounderId,
		PartyId:           req.PartyId,
		ActionType:        req.ActionType,
		Rectangle:         req.Rectangle,
		RouteJson:         req.RouteJson,
		Remark:            req.Remark,
		ScheduleStartTime: req.ScheduleStartTime,
		ScheduleEndTime:   req.ScheduleEndTime,
	}
	return db.CreateItinerary(s.ctx, itineraryModel)
}
