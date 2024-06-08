package main

import (
	trust "/home/fanr/projects/SOSD/bocchi-the-travel/kitex_gen/trust"
	"bocchi/kitex_gen/trust"
	"bocchi/pkg/errno"
	"bocchi/pkg/pack"
	"bocchi/rpc/trust/service"
	"context"
)

// FollowHandlerImpl implements the last service interface defined in the IDL.
type TrustHandlerImpl struct{}

// TrustAction implements the TrustHandlerImpl interface.
func (s *TrustHandlerImpl) TrustAction(ctx context.Context, req *trust.FollowActionRequest) (resp *trust.FollowActionResponse, err error) {
	resp = new(trust.FollowActionResponse)
	if req.ObjectUid == req.UserId {
		resp.Base = pack.BuildBaseResp(errno.FollowMyselfError)
		return resp, nil
	}
	switch req.ActionType {
	case 1:
		err = service.NewFollowService(ctx).Follow(req)

	case 0:
		err = service.NewFollowService(ctx).UnFollow(req)
	default:
		err = errno.ParamError
	}
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// FollowerList implements the TrustHandlerImpl interface.
func (s *TrustHandlerImpl) FollowerList(ctx context.Context, req *trust.FollowerListRequest) (resp *trust.FollowerListResponse, err error) {
	resp = new(trust.FollowerListResponse)
	followerResp, count, err := service.NewFollowService(ctx).FollowerList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.FollowerList = followerResp
	return resp, nil
}

// FollowingList implements the TrustHandlerImpl interface.
func (s *TrustHandlerImpl) FollowingList(ctx context.Context, req *trust.FollowingListRequest) (resp *trust.FollowingListResponse, err error) {
	resp = new(trust.FollowingListResponse)
	followingResp, count, err := service.NewFollowService(ctx).FollowingList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.FollowingList = followingResp
	return resp, nil
}

// TrustEachList implements the TrustHandlerImpl interface.
func (s *TrustHandlerImpl) TrustEachList(ctx context.Context, req *trust.FriendListRequest) (resp *trust.FriendListResponse, err error) {
	resp = new(trust.FriendListResponse)
	friendResp, count, err := service.NewFollowService(ctx).FriendList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.FriendList = friendResp
	return resp, nil
}

// MarkToOther implements the TrustHandlerImpl interface.
func (s *TrustHandlerImpl) MarkToOther(ctx context.Context, req *trust.MarkToOtherRequest) (resp *trust.MarkToOtherResponse, err error) {
	// TODO: Your code here...
	return
}
