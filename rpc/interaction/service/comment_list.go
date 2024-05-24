package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/kitex_gen/interaction"
	"bocchi/kitex_gen/user"
	"bocchi/pkg/errno"
	"bocchi/rpc/interaction/dal/db"
	"bocchi/rpc/interaction/rpc"
	"errors"
	"gorm.io/gorm"
)

func (s *InteractionService) CommentList(req *interaction.CommentListRequest) (*[]db.Comment, []*base.User, int64, error) {
	//commentCache, err := cache.GetVideoComments(s.ctx, req.PoiId)
	//if err != nil {
	//	return nil, 0, err
	//}
	//exist, countCache, err := cache.GetVideoCommentCount(s.ctx, req.VideoId)
	//if err != nil {
	//	return nil, 0, err
	//}
	//if exist && len(commentCache) != 0 {
	//	return commentCache, countCache, nil
	//}
	//if !exist && len(commentCache) != 0 {
	//	count, err := db.GetCommentCount(s.ctx, req.PoiId)
	//	if err != nil {
	//		return nil, 0, err
	//	}
	//	return commentCache, count, nil
	//}
	comments, count, err := db.GetCommentsByPoiID(s.ctx, req.PoiId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, nil
	}
	if err != nil {
		return nil, nil, 0, err
	}
	//设置缓存
	//if err := cache.SetVideoComments(s.ctx, comments, req.PoiId); err != nil {
	//	return nil, 0, err
	//}
	commenterIdList := make([]int64, len(*comments))
	var commenterList []*base.User
	for i, v := range *comments {
		commenterIdList[i] = v.Uid
	}
	rpcResp, err := rpc.UserGetCommenters(s.ctx, &user.GetUsersRequest{
		UserIdList: commenterIdList,
	})
	if err != nil {
		return nil, nil, 0, err
	}
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, nil, 0, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	commenterList = rpcResp.UserList
	return comments, commenterList, count, nil
}
