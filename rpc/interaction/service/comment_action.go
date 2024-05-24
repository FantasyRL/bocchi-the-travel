package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/kitex_gen/interaction"
	"bocchi/pkg/errno"
	"bocchi/rpc/interaction/dal/db"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) CommentCreate(req *interaction.CommentCreateRequest, uid int64) (*db.Comment, error) {
	var eg errgroup.Group
	var err error
	//var exist = false
	comment := new(db.Comment)
	commentModel := &base.Comment{
		PoiId:   req.PoiId,
		Content: req.Content,
		User: &base.User{
			Id: uid,
		},
	}

	eg.Go(func() error {
		//若内容完全重复，则删除最早发的那个(其实是懒得再开一个接口了)
		comment, err = db.CreateComment(s.ctx, commentModel)
		return err
		//return cache.AddVideoComment(s.ctx, comment)
	})
	//eg.Go(func() error {
	//	exist, _, err = cache.GetVideoCommentCount(s.ctx, req.VideoId)
	//	if err != nil {
	//		return err
	//	}
	//	if exist {
	//		return cache.IncrVideoCommentCount(s.ctx, req.VideoId)
	//	}
	//	return nil
	//})

	if err = eg.Wait(); err != nil {
		return nil, err
	}
	//if !exist {
	//	count, err := db.GetCommentCount(s.ctx, req.PoiId)
	//	if err != nil {
	//		return nil, err
	//	}
	//	err = cache.SetVideoCommentCount(s.ctx, req.PoiId, count)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	return comment, nil
}

func (s *InteractionService) CommentDelete(req *interaction.CommentDeleteRequest, uid int64) error {
	var eg errgroup.Group
	var commentModel = &base.Comment{
		Id: req.Id,
		User: &base.User{
			Id: uid,
		},
	}
	exist, err := db.IsCommentExist(s.ctx, commentModel)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistError
	}
	eg.Go(func() error {

		_, err = db.DeleteComment(s.ctx, commentModel)
		return err
		//return cache.DelVideoComment(s.ctx, comment)
	})

	//eg.Go(func() error {
	//	exist, _, err := cache.GetVideoCommentCount(s.ctx, req.VideoId)
	//	if err != nil {
	//		return err
	//	}
	//	if exist {
	//		return cache.DecrVideoCommentCount(s.ctx, req.VideoId)
	//	}
	//	return nil
	//})
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
