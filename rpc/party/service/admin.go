package service

import (
	"bocchi/kitex_gen/party"
	"bocchi/pkg/errno"
	"bocchi/rpc/party/dal/db"
)

func (s *PartyService) AddAdmin(req *party.AddAdminRequest) error {
	founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	if err != nil {
		return err
	}
	if founderId != req.UserId {
		return errno.AuthorizationError
	}
	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.TargetId,
		Status:   2,
	}
	return db.ChangeMemberStatus(s.ctx, memberModel)
}

func (s *PartyService) DeleteAdmin(req *party.DeleteAdminRequest) error {
	founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	if err != nil {
		return err
	}
	if founderId != req.UserId {
		return errno.AuthorizationError
	}
	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.TargetId,
		Status:   1,
	}
	return db.ChangeMemberStatus(s.ctx, memberModel)
}
