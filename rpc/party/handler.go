package main

import (
	party "bocchi/kitex_gen/party"
	"bocchi/kitex_gen/user"
	"bocchi/pkg/pack"
	"bocchi/rpc/party/rpc"
	"bocchi/rpc/party/service"
	"context"
)

// PartyHandlerImpl implements the last service interface defined in the IDL.
type PartyHandlerImpl struct{}

// CreateParty implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) CreateParty(ctx context.Context, req *party.CreatePartyRequest) (resp *party.CreatePartyResponse, err error) {
	resp = new(party.CreatePartyResponse)
	partyResp, err := service.NewPartyService(ctx).CreateParty(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Party = service.BuildPartyResp(partyResp)
	return resp, nil
}

// JoinParty implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) JoinParty(ctx context.Context, req *party.JoinPartyRequest) (resp *party.JoinPartyResponse, err error) {
	resp = new(party.JoinPartyResponse)
	err = service.NewPartyService(ctx).JoinParty(req)
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// ApplyList implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) ApplyList(ctx context.Context, req *party.ApplyListRequest) (resp *party.ApplyListResponse, err error) {
	resp = new(party.ApplyListResponse)
	members, count, err := service.NewPartyService(ctx).ApplyList(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	memberIdList := make([]int64, len(*members))
	for _, member := range *members {
		memberIdList = append(memberIdList, member.MemberId)
	}
	rpcResp, err := rpc.UserGetMembers(ctx, &user.GetUsersRequest{
		UserIdList: memberIdList,
	})
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.ApplicantCount = &count
	resp.ApplicantList = rpcResp.UserList

	return resp, nil
}

// PermitJoin implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) PermitJoin(ctx context.Context, req *party.PermitJoinRequest) (resp *party.PermitJoinResponse, err error) {
	resp = new(party.PermitJoinResponse)
	err = service.NewPartyService(ctx).PermitJoin(req)
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// GetPartyMembers implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) GetPartyMembers(ctx context.Context, req *party.GetPartyMembersRequest) (resp *party.GetPartyMembersResponse, err error) {
	resp = new(party.GetPartyMembersResponse)

	members, count, err := service.NewPartyService(ctx).GetPartyMembers(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	memberIdList := make([]int64, len(*members))
	for _, member := range *members {
		memberIdList = append(memberIdList, member.MemberId)
	}
	rpcResp, err := rpc.UserGetMembers(ctx, &user.GetUsersRequest{
		UserIdList: memberIdList,
	})
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.MemberCount = &count
	resp.MemberList = rpcResp.UserList
	return
}

// SearchParty implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) SearchParty(ctx context.Context, req *party.SearchPartyRequest) (resp *party.SearchPartyResponse, err error) {
	resp = new(party.SearchPartyResponse)

	parties, count, err := service.NewPartyService(ctx).SearchParty(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.PartyCount = &count
	resp.PartyList = service.BuildPartiesResp(parties)
	return resp, nil
}

// GetPartyInfo implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) GetPartyInfo(ctx context.Context, req *party.GetPartyInfoRequest) (resp *party.GetPartyInfoResponse, err error) {
	resp = new(party.GetPartyInfoResponse)
	partyResp, err := service.NewPartyService(ctx).GetPartyInfo(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Party = service.BuildPartyResp(partyResp)
	return resp, nil
}
