package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/kitex_gen/party"
	"bocchi/pkg/errno"
	"bocchi/rpc/itinerary/dal/db"
	"bocchi/rpc/itinerary/rpc"
)

func (s *ItineraryService) CreateItinerary(req *itinerary.CreateItineraryRequest) (*db.Itinerary, error) {
	rpcResp, err := rpc.IsMemberExist(s.ctx, &party.IsMemberExistRequest{
		PartyId: req.PartyId,
		UserId:  req.FounderId,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	if rpcResp.IsExist == false {
		return nil, errno.AuthorizationError
	}
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
