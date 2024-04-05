package main

import (
	party "bocchi/kitex_gen/party"
	"context"
)

// PartyHandlerImpl implements the last service interface defined in the IDL.
type PartyHandlerImpl struct{}

// CreateParty implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) CreateParty(ctx context.Context, req *party.CreatePartyRequest) (resp *party.CreatePartyResponse, err error) {
	// TODO: Your code here...
	return
}

// JoinParty implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) JoinParty(ctx context.Context, req *party.JoinPartyRequest) (resp *party.JoinPartyResponse, err error) {
	// TODO: Your code here...
	return
}

// ApplyList implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) ApplyList(ctx context.Context, req *party.ApplyListRequest) (resp *party.ApplyListResponse, err error) {
	// TODO: Your code here...
	return
}

// PermitJoin implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) PermitJoin(ctx context.Context, req *party.PermitJoinRequest) (resp *party.PermitJoinResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPartyMembers implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) GetPartyMembers(ctx context.Context, req *party.GetPartyMembersRequest) (resp *party.GetPartyMembersResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchParty implements the PartyHandlerImpl interface.
func (s *PartyHandlerImpl) SearchParty(ctx context.Context, req *party.SearchPartyRequest) (resp *party.SearchPartyResponse, err error) {
	// TODO: Your code here...
	return
}
