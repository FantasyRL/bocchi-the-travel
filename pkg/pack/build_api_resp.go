package pack

import (
	"bocchi/api/biz/model/api"
	base2 "bocchi/api/biz/model/base"
	"bocchi/kitex_gen/base"
	"bocchi/kitex_gen/user"
	"bocchi/pkg/errno"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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

func ToUserResp(_user interface{}) *api.User {
	//这里使用了一个及其抽象的断言
	p, _ := (_user).(*user.User)
	return &api.User{
		ID:        p.Id,
		Name:      p.Name,
		Email:     p.Email,
		Avatar:    p.Avatar,
		Signature: p.Signature,
	}
}
