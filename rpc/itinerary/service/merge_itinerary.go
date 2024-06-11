package service

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/kitex_gen/party"
	"bocchi/pkg/errno"
	"bocchi/rpc/itinerary/dal/db"
	"bocchi/rpc/itinerary/rpc"
)

func (s *ItineraryService) MergeItinerary(req *itinerary.MergeItineraryRequest) error {
	rpcResp, err := rpc.IsMemberAdmin(s.ctx, &party.IsMemberAdminRequest{
		PartyId: req.PartyId,
		UserId:  req.UserId,
	})
	if err != nil {
		return err
	}
	if rpcResp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	if rpcResp.IsExist == false {
		return errno.AuthorizationError
	}
	return db.ChangeItineraryStatus(s.ctx, req.PartyId, req.ItineraryId, 1)
}
