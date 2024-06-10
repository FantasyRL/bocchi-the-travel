package service

import (
	"bocchi/kitex_gen/trust"
	"bocchi/pkg/errno"
	"bocchi/rpc/trust/dal/cache"
	"bocchi/rpc/trust/dal/db"
)

func (s *FollowService) UnFollow(req *trust.FollowActionRequest) error {
	//redis
	e, err := cache.IsFollowedCacheExist(s.ctx, req.ObjectUid)
	if err != nil {
		return err
	}
	if e {
		e1, err := cache.IsUserFollowExist(s.ctx, req.UserId, req.ObjectUid)
		if err != nil {
			return err
		}
		if !e1 {
			return errno.FollowNotExistError
		}
		err = cache.DelFollower(s.ctx, req.UserId, req.ObjectUid)
		if err != nil {
			return err
		}
	}

	//mysql
	e1, err := db.IsFollowStatus(s.ctx, req.UserId, req.ObjectUid, 0)
	if err != nil {
		return err
	}
	if e1 {
		return errno.FollowNotExistError
	}

	e2, err := db.IsFollowExist(s.ctx, req.UserId, req.ObjectUid)
	if err != nil {
		return err
	}
	if !e2 {
		return errno.FollowNotExistError
	}
	return db.UpdateFollowStatus(s.ctx, req.UserId, req.ObjectUid, 0)
}
