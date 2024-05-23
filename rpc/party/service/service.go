package service

import (
	"bocchi/kitex_gen/base"
	"bocchi/rpc/party/dal/db"
	"context"
)

type PartyService struct {
	ctx context.Context
}

func NewPartyService(ctx context.Context) *PartyService {
	return &PartyService{
		ctx: ctx,
	}
}

func BuildPartyResp(dbParty *db.Party) *base.Party {
	_, memberCount, _ := db.GetMemberListByStatus(context.TODO(), dbParty.Id, 0, 1)
	return &base.Party{
		Id:          dbParty.Id,
		FounderId:   dbParty.FounderId,
		Title:       dbParty.Title,
		Content:     dbParty.Content,
		Type:        dbParty.Type,
		Province:    dbParty.Province,
		City:        dbParty.City,
		StartTime:   dbParty.StartTime.Format("2006-01-02"),
		EndTime:     dbParty.EndTime.Format("2006-01-02"),
		MemberCount: memberCount,
		Status:      dbParty.Status,
		Rectangle:   dbParty.Rectangle,
	}
}

func BuildPartiesResp(dbParties *[]db.Party) []*base.Party {
	partiesResp := make([]*base.Party, len(*dbParties))
	for i, dbParty := range *dbParties {
		partiesResp[i] = BuildPartyResp(&dbParty)
	}
	return partiesResp
}
