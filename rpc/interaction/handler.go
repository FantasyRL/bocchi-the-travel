package main

import (
	"bocchi/kitex_gen/interaction"
	"bocchi/pkg/pack"
	"bocchi/rpc/interaction/service"
	"context"
)

// InteractionHandlerImpl implements the last service interface defined in the IDL.
type InteractionHandlerImpl struct{}

// CommentCreate implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest) (resp *interaction.CommentCreateResponse, err error) {
	resp = new(interaction.CommentCreateResponse)
	_, err = service.NewInteractionService(ctx).CommentCreate(req, req.UserId)

	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// CommentDelete implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest) (resp *interaction.CommentDeleteResponse, err error) {
	resp = new(interaction.CommentDeleteResponse)
	err = service.NewInteractionService(ctx).CommentDelete(req, req.UserId)
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// CommentList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	resp = new(interaction.CommentListResponse)
	commentResp, commenterList, count, err := service.NewInteractionService(ctx).CommentList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.CommentCount = &count
	resp.CommentList = service.BuildCommentsResp(commentResp, commenterList)
	return resp, nil
}
