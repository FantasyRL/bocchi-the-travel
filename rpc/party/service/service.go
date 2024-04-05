package service

import "context"

type PartyService struct {
	ctx context.Context
}

func NewPartyService(ctx context.Context) *PartyService {
	return &PartyService{
		ctx: ctx,
	}
}
