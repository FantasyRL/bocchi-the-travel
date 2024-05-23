package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/rpc/itinerary/dal/db"
	"context"
)

type ItineraryService struct {
	ctx context.Context
}

func NewItineraryService(ctx context.Context) *ItineraryService {
	return &ItineraryService{
		ctx: ctx,
	}
}

func BuildItineraryResp(dbItinerary *db.Itinerary) *base.Itinerary {
	return &base.Itinerary{
		Id:                dbItinerary.Id,
		Title:             dbItinerary.Title,
		FounderId:         dbItinerary.FounderId,
		ActionType:        dbItinerary.ActionType,
		Rectangle:         dbItinerary.Rectangle,
		RouteJson:         dbItinerary.RouteJson,
		Remark:            dbItinerary.Remark,
		Sequence:          dbItinerary.Sequence,
		ScheduleStartTime: dbItinerary.ScheduleStartTime,
		ScheduleEndTime:   dbItinerary.ScheduleEndTime,
		PartyId:           dbItinerary.PartyId,
		IsMerged:          dbItinerary.IsMerged,
	}
}

func BuildItinerariesResp(dbItineraries *[]db.Itinerary) []*base.Itinerary {
	partiesResp := make([]*base.Itinerary, len(*dbItineraries))
	for i, dbItinerary := range *dbItineraries {
		partiesResp[i] = BuildItineraryResp(&dbItinerary)
	}
	return partiesResp
}
