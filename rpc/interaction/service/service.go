package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/rpc/interaction/dal/db"
	"context"
	"time"
)

type InteractionService struct {
	ctx context.Context
}

func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}

func BuildCommentResp(comment *db.Comment, commenter *base.User) *base.Comment {
	cr := comment.CreatedAt.Add(time.Hour * 8)
	return &base.Comment{
		Id:          comment.ID,
		PoiId:       comment.PoiID,
		User:        commenter,
		Content:     comment.Content,
		PublishTime: cr.Format("2006-01-02 15:01:04"),
	}
}

func BuildCommentsResp(comments *[]db.Comment, commenterList []*base.User) []*base.Comment {
	commentListResp := make([]*base.Comment, len(*comments))
	for i, v := range *comments {
		commentListResp[i] = BuildCommentResp(&v, commenterList[i])
	}
	return commentListResp
}
