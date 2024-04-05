package service

import (
	"bocchi/kitex_gen/party"
	"bocchi/rpc/party/dal/db"
	"time"
)

func (s *PartyService) CreateParty(req party.CreatePartyRequest) error {
	st, _ := time.Parse("2006-01-02", req.StartTime)
	et, _ := time.Parse("2006-01-02", req.EndTime)
	partyModel := &db.Party{
		FounderId: req.FounderId,
		Title:     req.Title,
		Content:   req.Content,
		Type:      req.Type,
		Province:  req.Province,
		City:      req.City,
		StartTime: st,
		EndTime:   et,
	}
	return db.CreateParty(s.ctx, partyModel)
}

func (s *PartyService) SearchParty(req *party.SearchPartyRequest) (*[]db.Party, int64, error) {
	return db.GetPartyByMultiple(s.ctx, req)
}
