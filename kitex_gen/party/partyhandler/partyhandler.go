// Code generated by Kitex v0.9.1. DO NOT EDIT.

package partyhandler

import (
	party "bocchi/kitex_gen/party"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"DeleteMember": kitex.NewMethodInfo(
		deleteMemberHandler,
		newPartyHandlerDeleteMemberArgs,
		newPartyHandlerDeleteMemberResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"IsMemberAdmin": kitex.NewMethodInfo(
		isMemberAdminHandler,
		newPartyHandlerIsMemberAdminArgs,
		newPartyHandlerIsMemberAdminResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"IsMemberExist": kitex.NewMethodInfo(
		isMemberExistHandler,
		newPartyHandlerIsMemberExistArgs,
		newPartyHandlerIsMemberExistResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AddAdmin": kitex.NewMethodInfo(
		addAdminHandler,
		newPartyHandlerAddAdminArgs,
		newPartyHandlerAddAdminResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteAdmin": kitex.NewMethodInfo(
		deleteAdminHandler,
		newPartyHandlerDeleteAdminArgs,
		newPartyHandlerDeleteAdminResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateParty": kitex.NewMethodInfo(
		createPartyHandler,
		newPartyHandlerCreatePartyArgs,
		newPartyHandlerCreatePartyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"JoinParty": kitex.NewMethodInfo(
		joinPartyHandler,
		newPartyHandlerJoinPartyArgs,
		newPartyHandlerJoinPartyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ApplyList": kitex.NewMethodInfo(
		applyListHandler,
		newPartyHandlerApplyListArgs,
		newPartyHandlerApplyListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"PermitJoin": kitex.NewMethodInfo(
		permitJoinHandler,
		newPartyHandlerPermitJoinArgs,
		newPartyHandlerPermitJoinResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPartyInfo": kitex.NewMethodInfo(
		getPartyInfoHandler,
		newPartyHandlerGetPartyInfoArgs,
		newPartyHandlerGetPartyInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPartyMembers": kitex.NewMethodInfo(
		getPartyMembersHandler,
		newPartyHandlerGetPartyMembersArgs,
		newPartyHandlerGetPartyMembersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SearchParty": kitex.NewMethodInfo(
		searchPartyHandler,
		newPartyHandlerSearchPartyArgs,
		newPartyHandlerSearchPartyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetMyParties": kitex.NewMethodInfo(
		getMyPartiesHandler,
		newPartyHandlerGetMyPartiesArgs,
		newPartyHandlerGetMyPartiesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ChangePartyStatus": kitex.NewMethodInfo(
		changePartyStatusHandler,
		newPartyHandlerChangePartyStatusArgs,
		newPartyHandlerChangePartyStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	partyHandlerServiceInfo                = NewServiceInfo()
	partyHandlerServiceInfoForClient       = NewServiceInfoForClient()
	partyHandlerServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return partyHandlerServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return partyHandlerServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return partyHandlerServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PartyHandler"
	handlerType := (*party.PartyHandler)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "party",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func deleteMemberHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerDeleteMemberArgs)
	realResult := result.(*party.PartyHandlerDeleteMemberResult)
	success, err := handler.(party.PartyHandler).DeleteMember(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerDeleteMemberArgs() interface{} {
	return party.NewPartyHandlerDeleteMemberArgs()
}

func newPartyHandlerDeleteMemberResult() interface{} {
	return party.NewPartyHandlerDeleteMemberResult()
}

func isMemberAdminHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerIsMemberAdminArgs)
	realResult := result.(*party.PartyHandlerIsMemberAdminResult)
	success, err := handler.(party.PartyHandler).IsMemberAdmin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerIsMemberAdminArgs() interface{} {
	return party.NewPartyHandlerIsMemberAdminArgs()
}

func newPartyHandlerIsMemberAdminResult() interface{} {
	return party.NewPartyHandlerIsMemberAdminResult()
}

func isMemberExistHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerIsMemberExistArgs)
	realResult := result.(*party.PartyHandlerIsMemberExistResult)
	success, err := handler.(party.PartyHandler).IsMemberExist(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerIsMemberExistArgs() interface{} {
	return party.NewPartyHandlerIsMemberExistArgs()
}

func newPartyHandlerIsMemberExistResult() interface{} {
	return party.NewPartyHandlerIsMemberExistResult()
}

func addAdminHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerAddAdminArgs)
	realResult := result.(*party.PartyHandlerAddAdminResult)
	success, err := handler.(party.PartyHandler).AddAdmin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerAddAdminArgs() interface{} {
	return party.NewPartyHandlerAddAdminArgs()
}

func newPartyHandlerAddAdminResult() interface{} {
	return party.NewPartyHandlerAddAdminResult()
}

func deleteAdminHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerDeleteAdminArgs)
	realResult := result.(*party.PartyHandlerDeleteAdminResult)
	success, err := handler.(party.PartyHandler).DeleteAdmin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerDeleteAdminArgs() interface{} {
	return party.NewPartyHandlerDeleteAdminArgs()
}

func newPartyHandlerDeleteAdminResult() interface{} {
	return party.NewPartyHandlerDeleteAdminResult()
}

func createPartyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerCreatePartyArgs)
	realResult := result.(*party.PartyHandlerCreatePartyResult)
	success, err := handler.(party.PartyHandler).CreateParty(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerCreatePartyArgs() interface{} {
	return party.NewPartyHandlerCreatePartyArgs()
}

func newPartyHandlerCreatePartyResult() interface{} {
	return party.NewPartyHandlerCreatePartyResult()
}

func joinPartyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerJoinPartyArgs)
	realResult := result.(*party.PartyHandlerJoinPartyResult)
	success, err := handler.(party.PartyHandler).JoinParty(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerJoinPartyArgs() interface{} {
	return party.NewPartyHandlerJoinPartyArgs()
}

func newPartyHandlerJoinPartyResult() interface{} {
	return party.NewPartyHandlerJoinPartyResult()
}

func applyListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerApplyListArgs)
	realResult := result.(*party.PartyHandlerApplyListResult)
	success, err := handler.(party.PartyHandler).ApplyList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerApplyListArgs() interface{} {
	return party.NewPartyHandlerApplyListArgs()
}

func newPartyHandlerApplyListResult() interface{} {
	return party.NewPartyHandlerApplyListResult()
}

func permitJoinHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerPermitJoinArgs)
	realResult := result.(*party.PartyHandlerPermitJoinResult)
	success, err := handler.(party.PartyHandler).PermitJoin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerPermitJoinArgs() interface{} {
	return party.NewPartyHandlerPermitJoinArgs()
}

func newPartyHandlerPermitJoinResult() interface{} {
	return party.NewPartyHandlerPermitJoinResult()
}

func getPartyInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerGetPartyInfoArgs)
	realResult := result.(*party.PartyHandlerGetPartyInfoResult)
	success, err := handler.(party.PartyHandler).GetPartyInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerGetPartyInfoArgs() interface{} {
	return party.NewPartyHandlerGetPartyInfoArgs()
}

func newPartyHandlerGetPartyInfoResult() interface{} {
	return party.NewPartyHandlerGetPartyInfoResult()
}

func getPartyMembersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerGetPartyMembersArgs)
	realResult := result.(*party.PartyHandlerGetPartyMembersResult)
	success, err := handler.(party.PartyHandler).GetPartyMembers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerGetPartyMembersArgs() interface{} {
	return party.NewPartyHandlerGetPartyMembersArgs()
}

func newPartyHandlerGetPartyMembersResult() interface{} {
	return party.NewPartyHandlerGetPartyMembersResult()
}

func searchPartyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerSearchPartyArgs)
	realResult := result.(*party.PartyHandlerSearchPartyResult)
	success, err := handler.(party.PartyHandler).SearchParty(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerSearchPartyArgs() interface{} {
	return party.NewPartyHandlerSearchPartyArgs()
}

func newPartyHandlerSearchPartyResult() interface{} {
	return party.NewPartyHandlerSearchPartyResult()
}

func getMyPartiesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerGetMyPartiesArgs)
	realResult := result.(*party.PartyHandlerGetMyPartiesResult)
	success, err := handler.(party.PartyHandler).GetMyParties(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerGetMyPartiesArgs() interface{} {
	return party.NewPartyHandlerGetMyPartiesArgs()
}

func newPartyHandlerGetMyPartiesResult() interface{} {
	return party.NewPartyHandlerGetMyPartiesResult()
}

func changePartyStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*party.PartyHandlerChangePartyStatusArgs)
	realResult := result.(*party.PartyHandlerChangePartyStatusResult)
	success, err := handler.(party.PartyHandler).ChangePartyStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPartyHandlerChangePartyStatusArgs() interface{} {
	return party.NewPartyHandlerChangePartyStatusArgs()
}

func newPartyHandlerChangePartyStatusResult() interface{} {
	return party.NewPartyHandlerChangePartyStatusResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DeleteMember(ctx context.Context, req *party.DeleteMemberRequest) (r *party.DeleteMemberResponse, err error) {
	var _args party.PartyHandlerDeleteMemberArgs
	_args.Req = req
	var _result party.PartyHandlerDeleteMemberResult
	if err = p.c.Call(ctx, "DeleteMember", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsMemberAdmin(ctx context.Context, req *party.IsMemberAdminRequest) (r *party.IsMemberAdminResponse, err error) {
	var _args party.PartyHandlerIsMemberAdminArgs
	_args.Req = req
	var _result party.PartyHandlerIsMemberAdminResult
	if err = p.c.Call(ctx, "IsMemberAdmin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsMemberExist(ctx context.Context, req *party.IsMemberExistRequest) (r *party.IsMemberExistResponse, err error) {
	var _args party.PartyHandlerIsMemberExistArgs
	_args.Req = req
	var _result party.PartyHandlerIsMemberExistResult
	if err = p.c.Call(ctx, "IsMemberExist", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddAdmin(ctx context.Context, req *party.AddAdminRequest) (r *party.AddAdminResponse, err error) {
	var _args party.PartyHandlerAddAdminArgs
	_args.Req = req
	var _result party.PartyHandlerAddAdminResult
	if err = p.c.Call(ctx, "AddAdmin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAdmin(ctx context.Context, req *party.DeleteAdminRequest) (r *party.DeleteAdminResponse, err error) {
	var _args party.PartyHandlerDeleteAdminArgs
	_args.Req = req
	var _result party.PartyHandlerDeleteAdminResult
	if err = p.c.Call(ctx, "DeleteAdmin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateParty(ctx context.Context, req *party.CreatePartyRequest) (r *party.CreatePartyResponse, err error) {
	var _args party.PartyHandlerCreatePartyArgs
	_args.Req = req
	var _result party.PartyHandlerCreatePartyResult
	if err = p.c.Call(ctx, "CreateParty", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) JoinParty(ctx context.Context, req *party.JoinPartyRequest) (r *party.JoinPartyResponse, err error) {
	var _args party.PartyHandlerJoinPartyArgs
	_args.Req = req
	var _result party.PartyHandlerJoinPartyResult
	if err = p.c.Call(ctx, "JoinParty", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ApplyList(ctx context.Context, req *party.ApplyListRequest) (r *party.ApplyListResponse, err error) {
	var _args party.PartyHandlerApplyListArgs
	_args.Req = req
	var _result party.PartyHandlerApplyListResult
	if err = p.c.Call(ctx, "ApplyList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PermitJoin(ctx context.Context, req *party.PermitJoinRequest) (r *party.PermitJoinResponse, err error) {
	var _args party.PartyHandlerPermitJoinArgs
	_args.Req = req
	var _result party.PartyHandlerPermitJoinResult
	if err = p.c.Call(ctx, "PermitJoin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPartyInfo(ctx context.Context, req *party.GetPartyInfoRequest) (r *party.GetPartyInfoResponse, err error) {
	var _args party.PartyHandlerGetPartyInfoArgs
	_args.Req = req
	var _result party.PartyHandlerGetPartyInfoResult
	if err = p.c.Call(ctx, "GetPartyInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPartyMembers(ctx context.Context, req *party.GetPartyMembersRequest) (r *party.GetPartyMembersResponse, err error) {
	var _args party.PartyHandlerGetPartyMembersArgs
	_args.Req = req
	var _result party.PartyHandlerGetPartyMembersResult
	if err = p.c.Call(ctx, "GetPartyMembers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchParty(ctx context.Context, req *party.SearchPartyRequest) (r *party.SearchPartyResponse, err error) {
	var _args party.PartyHandlerSearchPartyArgs
	_args.Req = req
	var _result party.PartyHandlerSearchPartyResult
	if err = p.c.Call(ctx, "SearchParty", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMyParties(ctx context.Context, req *party.GetMyPartiesRequest) (r *party.GetMyPartiesResponse, err error) {
	var _args party.PartyHandlerGetMyPartiesArgs
	_args.Req = req
	var _result party.PartyHandlerGetMyPartiesResult
	if err = p.c.Call(ctx, "GetMyParties", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangePartyStatus(ctx context.Context, req *party.ChangePartyStatusRequest) (r *party.ChangePartyStatusResponse, err error) {
	var _args party.PartyHandlerChangePartyStatusArgs
	_args.Req = req
	var _result party.PartyHandlerChangePartyStatusResult
	if err = p.c.Call(ctx, "ChangePartyStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
