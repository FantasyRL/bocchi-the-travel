package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/kitex_gen/trust"
	"bocchi/kitex_gen/user"
	"bocchi/pkg/errno"
	"bocchi/rpc/trust/dal/db"
	"bocchi/rpc/trust/rpc"
)

func (s *FollowService) FriendList(req *trust.FriendListRequest) ([]*base.User, int64, error) {
	_, friendList, _, err := db.FollowerAndFriendList(s.ctx, req.UserId)
	if err != nil {
		return nil, 0, err
	}
	count := int64(len(*friendList))
	friendIdList := make([]int64, len(*friendList))
	for i, v := range *friendList {
		friendIdList[i] = v.FollowedId
	}

	rpcResp, err := rpc.UserGetUserList(s.ctx, &user.GetUsersRequest{
		UserIdList: friendIdList,
	})
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	if err != nil {
		return nil, 0, err
	}
	t := true
	for i := range rpcResp.UserList {
		rpcResp.UserList[i].IsTrust = &t
	}
	return rpcResp.UserList, count, nil
}

func FriendCount(uid int64) {

}
