package service

import (
	"bocchi/kitex_gen/party"
	"bocchi/rpc/party/dal/db"
)

func (s *PartyService) IsMemberExist(req *party.IsMemberExistRequest) (bool, error) {
	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.UserId,
	}
	return db.ISMemberExist(s.ctx, memberModel)
}

func (s *PartyService) IsMemberAdmin(req *party.IsMemberAdminRequest) (bool, error) {
	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.UserId,
	}
	return db.ISMemberAdmin(s.ctx, memberModel)
}
