package service

import (
	"bocchi/kitex_gen/party"
	"bocchi/pkg/errno"
	"bocchi/rpc/party/dal/db"
)

func (s *PartyService) JoinParty(req *party.JoinPartyRequest) error {
	founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	if err != nil {
		return err
	}
	if founderId == req.MemberId {
		return errno.FounderError
	}

	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.MemberId,
	}
	ok, err := db.ISMemberExist(s.ctx, memberModel)
	if !ok {
		return db.ApplyToParty(s.ctx, memberModel)
	}
	return err

}

func (s *PartyService) ApplyList(req *party.ApplyListRequest) (*[]db.Member, int64, error) {
	//founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	//if err != nil {
	//	return nil, 0, err
	//}
	//if founderId != req.UserId {
	//	return nil, 0, errno.NotFounderError
	//}
	memberModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.UserId,
	}
	is, err := db.ISMemberAdmin(s.ctx, memberModel)
	if err != nil {
		return nil, 0, err
	}
	if !is {
		return nil, 0, errno.NotAdminError
	}
	return db.GetMemberListByStatus(s.ctx, req.PartyId, req.PageNum, 0)
}

func (s *PartyService) PermitJoin(req *party.PermitJoinRequest) error {
	//founderId, err := db.GetFounderIdByPartyId(s.ctx, req.PartyId)
	//if err != nil {
	//	return err
	//}
	//if founderId != req.UserId {
	//	return errno.NotFounderError
	//}

	adminModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.UserId,
	}
	is, err := db.ISMemberAdmin(s.ctx, adminModel)
	if err != nil {
		return err
	}
	if !is {
		return errno.NotAdminError
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

func (s *PartyService) DeleteMember(req *party.DeleteMemberRequest) error {
	adminModel := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.UserId,
	}
	is, err := db.ISMemberAdmin(s.ctx, adminModel)
	if err != nil {
		return err
	}
	if !is {
		return errno.NotAdminError
	}
	adminModel2 := &db.Member{
		PartyId:  req.PartyId,
		MemberId: req.TargetId,
	}
	is2, err := db.ISMemberAdmin(s.ctx, adminModel2)
	if err != nil {
		return err
	}
	if !is2 {
		return errno.CantKickAdminError
	}

	return db.DeleteMember(s.ctx, req.TargetId)
}
