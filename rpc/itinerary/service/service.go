package service

import "context"

type ItineraryService struct {
	ctx context.Context
}

func NewItineraryService(ctx context.Context) *ItineraryService {
	return &ItineraryService{
		ctx: ctx,
	}
}
