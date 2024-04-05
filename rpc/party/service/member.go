package service

import (
	"bocchi/kitex_gen/party"
	"bocchi/pkg/errno"
	"bocchi/rpc/party/dal/db"
)

func (s *PartyService) JoinParty(req *party.JoinPartyRequest) error {
	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.MemberId,
	}
	if err := db.ISMemberExist(s.ctx, memberModel); err != nil {
		return err
	}
	return db.ApplyToParty(s.ctx, memberModel)
}

func (s *PartyService) ApplyList(req *party.ApplyListRequest) (*[]db.Member, int64, error) {
	founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	if err != nil {
		return nil, 0, err
	}
	if founderId != req.UserId {
		return nil, 0, errno.NotFounderError
	}
	return db.GetMemberListByStatus(s.ctx, req.PartyId, req.PageNum, 0)
}

func (s *PartyService) PermitJoin(req *party.PermitJoinRequest) error {
	founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	if err != nil {
		return err
	}
	if founderId != req.UserId {
		return errno.NotFounderError
	}

	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.MemberId,
		Status:   1,
	}
	if err = db.CheckMemberStatus(s.ctx, memberModel); err != nil {
		return err
	}
	return db.ChangeMemberStatus(s.ctx, memberModel)
}

func (s *PartyService) GetPartyMembers(req *party.GetPartyMembersRequest) (*[]db.Member, int64, error) {
	return db.GetMemberListByStatus(s.ctx, req.PartyId, req.PageNum, 1)
}
