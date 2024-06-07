package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/kitex_gen/trust"
	"bocchi/kitex_gen/user"
	"bocchi/pkg/errno"
	"bocchi/rpc/trust/dal/db"
	"bocchi/rpc/trust/rpc"
)

func (s *FollowService) FollowerList(req *trust.FollowerListRequest) ([]*base.User, int64, error) {
	////redis
	//followerList, err := cache.GetFollower(s.ctx, uid)
	//if err != nil {
	//	return nil,nil, 0, err
	//}
	//e1, count, err := cache.GetFollowerCount(s.ctx, uid)
	//if err != nil {
	//	return nil, nil,0, err
	//}
	//
	//if len(followerList) != 0 && e1 {
	//	return followerList, nil,count, nil
	//}//todo:redis

	//mysql
	followerList, friendList, count, err := db.FollowerAndFriendList(s.ctx, req.UserId)
	if err != nil {
		return nil, 0, err
	}
	//if err = cache.SetFollowerList(s.ctx, uid, followerList); err != nil {
	//	return nil, nil,0, err
	//}
	//
	//if err = cache.SetFollowerCount(s.ctx, uid, count); err != nil {
	//	return nil, nil,0, err
	//}

	followerIdList := make([]int64, len(*followerList))
	for i, v := range *followerList {
		followerIdList[i] = v.Uid
	}
	friendIdList := make([]int64, len(*friendList))
	for i, v := range *friendList {
		friendIdList[i] = v.FollowedId
	}
	rpcResp, err := rpc.UserGetUserList(s.ctx, &user.GetUsersRequest{
		UserIdList: followerIdList,
	})
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	if err != nil {
		return nil, 0, err
	}
	for i, v := range rpcResp.UserList {
		for j := 0; j < len(*friendList); j++ {
			if v.Id == friendIdList[j] {
				rpcResp.UserList[i].IsFollow = true
				break
			}
		}
	}
	return rpcResp.UserList, count, nil
}

func FollowerCount(followedId int64) {

}
