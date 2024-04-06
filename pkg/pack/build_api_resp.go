package pack

import (
	"bocchi/api/biz/model/api"
	base2 "bocchi/api/biz/model/base"
	"bocchi/kitex_gen/base"
	"bocchi/kitex_gen/party"
	"bocchi/kitex_gen/user"
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

func ConvertToAPIUser(kitexUser *user.User) *api.User {
	return &api.User{
		ID:        kitexUser.Id,
		Name:      kitexUser.Name,
		Email:     kitexUser.Email,
		Avatar:    kitexUser.Avatar,
		Signature: kitexUser.Signature,
	}
}

func ConvertToAPIUsers(kitexUsers []*user.User) []*api.User {
	usersResp := make([]*api.User, len(kitexUsers))
	for i, kitexUser := range kitexUsers {
		usersResp[i] = ConvertToAPIUser(kitexUser)
	}
	return usersResp
}

func ConvertToAPIParty(kitexParty *party.Party) *api.Party {
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
	}
}

func ConvertToAPIParties(kitexParties []*party.Party) []*api.Party {
	partiesResp := make([]*api.Party, len(kitexParties))
	for i, kitexParty := range kitexParties {
		partiesResp[i] = ConvertToAPIParty(kitexParty)
	}
	return partiesResp
}

func ToUserResp(_user interface{}) *api.User {
	//这里使用了一个及其抽象的断言
	p, _ := (_user).(*user.User)
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
