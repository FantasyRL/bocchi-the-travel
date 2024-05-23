package pack

import (
	"bocchi/api/biz/model/api"
	base2 "bocchi/api/biz/model/base"
	"bocchi/kitex_gen/base"
	"bocchi/pkg/errno"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return ErrToResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToResp(e)
	}

	_e := errno.ServiceError.WithMessage(err.Error()) //未知错误
	return ErrToResp(_e)
}

func ErrToResp(err errno.ErrNo) *base.BaseResp {
	return &base.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func ConvertToAPIBaseResp(baseResp *base.BaseResp) *base2.BaseResp {
	return &base2.BaseResp{
		Code: baseResp.Code,
		Msg:  baseResp.Msg,
	}
}

func SendRPCFailResp(c *app.RequestContext, err error) {
	klog.Error("rpc failed:", err)
	c.JSON(consts.StatusOK, base2.BaseResp{
		Code: -1,
		Msg:  errno.ConvertErr(err).Error(),
	})
}

func ConvertToAPIUser(kitexUser *base.User) *api.User {
	return &api.User{
		ID:        kitexUser.Id,
		Name:      kitexUser.Name,
		Email:     kitexUser.Email,
		Avatar:    kitexUser.Avatar,
		Signature: kitexUser.Signature,
	}
}

func ConvertToAPIUsers(kitexUsers []*base.User) []*api.User {
	usersResp := make([]*api.User, len(kitexUsers))
	for i, kitexUser := range kitexUsers {
		usersResp[i] = ConvertToAPIUser(kitexUser)
	}
	return usersResp
}

func ConvertToAPIParty(kitexParty *base.Party) *api.Party {
	return &api.Party{
		ID:          kitexParty.Id,
		FounderID:   kitexParty.FounderId,
		Title:       kitexParty.Title,
		Content:     kitexParty.Content,
		Type:        kitexParty.Type,
		Province:    kitexParty.Province,
		City:        kitexParty.City,
		StartTime:   kitexParty.StartTime,
		EndTime:     kitexParty.EndTime,
		MemberCount: kitexParty.MemberCount,
		Status:      kitexParty.Status,
		Rectangle:   kitexParty.Rectangle,
	}
}

func ConvertToAPIParties(kitexParties []*base.Party) []*api.Party {
	partiesResp := make([]*api.Party, len(kitexParties))
	for i, kitexParty := range kitexParties {
		partiesResp[i] = ConvertToAPIParty(kitexParty)
	}
	return partiesResp
}

func ConvertToAPIItinerary(kitexItinerary *base.Itinerary) *api.Itinerary {
	return &api.Itinerary{
		ID:                kitexItinerary.Id,
		Title:             kitexItinerary.Title,
		FounderID:         kitexItinerary.FounderId,
		ActionType:        kitexItinerary.ActionType,
		Rectangle:         kitexItinerary.Rectangle,
		RouteJSON:         kitexItinerary.RouteJson,
		Remark:            kitexItinerary.Remark,
		Sequence:          kitexItinerary.Sequence,
		ScheduleStartTime: kitexItinerary.ScheduleStartTime,
		ScheduleEndTime:   kitexItinerary.ScheduleEndTime,
		PartyID:           kitexItinerary.PartyId,
		IsMerged:          kitexItinerary.IsMerged,
	}
}

func ConvertToAPIItineraries(kitexItineraries []*base.Itinerary) []*api.Itinerary {
	ItinerariesResp := make([]*api.Itinerary, len(kitexItineraries))
	for i, kitexItinerary := range kitexItineraries {
		ItinerariesResp[i] = ConvertToAPIItinerary(kitexItinerary)
	}
	return ItinerariesResp
}

func ToUserResp(_user interface{}) *api.User {
	//这里使用了一个及其抽象的断言
	p, _ := (_user).(*base.User)
	if p == nil {
		return nil
	}
	return &api.User{
		ID:        p.Id,
		Name:      p.Name,
		Email:     p.Email,
		Avatar:    p.Avatar,
		Signature: p.Signature,
	}
}
