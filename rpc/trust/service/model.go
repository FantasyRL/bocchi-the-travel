package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/rpc/user/dal/db"
	"context"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{
		ctx: ctx,
	}
}

func BuildFollowedUsersResp(users []db.User) (usersResp []*base.User) {
	usersResp = make([]*base.User, 0, len(users))
	t := true
	for _, u := range users {
		usersResp = append(usersResp, &base.User{
			Id:      u.ID,
			Name:    u.UserName,
			Avatar:  u.Avatar,
			IsTrust: &t,
		})
	}
	return
}

func BuildFollowerUsersResp(uid int64, users []db.User, isFollowList []bool) (usersResp []*base.User) {
	usersResp = make([]*base.User, 0, len(users))
	for i, u := range users {
		usersResp = append(usersResp, &base.User{
			Id:      u.ID,
			Name:    u.UserName,
			Avatar:  u.Avatar,
			IsTrust: &isFollowList[i],
		})
	}
	return
}
