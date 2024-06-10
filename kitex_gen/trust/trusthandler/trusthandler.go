// Code generated by Kitex v0.9.1. DO NOT EDIT.

package trusthandler

import (
	trust "bocchi/kitex_gen/trust"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"IsTrust": kitex.NewMethodInfo(
		isTrustHandler,
		newTrustHandlerIsTrustArgs,
		newTrustHandlerIsTrustResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"TrustAction": kitex.NewMethodInfo(
		trustActionHandler,
		newTrustHandlerTrustActionArgs,
		newTrustHandlerTrustActionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"FollowerList": kitex.NewMethodInfo(
		followerListHandler,
		newTrustHandlerFollowerListArgs,
		newTrustHandlerFollowerListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"FollowingList": kitex.NewMethodInfo(
		followingListHandler,
		newTrustHandlerFollowingListArgs,
		newTrustHandlerFollowingListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"MarkToOther": kitex.NewMethodInfo(
		markToOtherHandler,
		newTrustHandlerMarkToOtherArgs,
		newTrustHandlerMarkToOtherResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"TrustEachList": kitex.NewMethodInfo(
		trustEachListHandler,
		newTrustHandlerTrustEachListArgs,
		newTrustHandlerTrustEachListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUserScore": kitex.NewMethodInfo(
		getUserScoreHandler,
		newTrustHandlerGetUserScoreArgs,
		newTrustHandlerGetUserScoreResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	trustHandlerServiceInfo                = NewServiceInfo()
	trustHandlerServiceInfoForClient       = NewServiceInfoForClient()
	trustHandlerServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return trustHandlerServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return trustHandlerServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return trustHandlerServiceInfoForClient
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
	serviceName := "TrustHandler"
	handlerType := (*trust.TrustHandler)(nil)
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
		"PackageName": "trust",
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

func isTrustHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerIsTrustArgs)
	realResult := result.(*trust.TrustHandlerIsTrustResult)
	success, err := handler.(trust.TrustHandler).IsTrust(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerIsTrustArgs() interface{} {
	return trust.NewTrustHandlerIsTrustArgs()
}

func newTrustHandlerIsTrustResult() interface{} {
	return trust.NewTrustHandlerIsTrustResult()
}

func trustActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerTrustActionArgs)
	realResult := result.(*trust.TrustHandlerTrustActionResult)
	success, err := handler.(trust.TrustHandler).TrustAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerTrustActionArgs() interface{} {
	return trust.NewTrustHandlerTrustActionArgs()
}

func newTrustHandlerTrustActionResult() interface{} {
	return trust.NewTrustHandlerTrustActionResult()
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerFollowerListArgs)
	realResult := result.(*trust.TrustHandlerFollowerListResult)
	success, err := handler.(trust.TrustHandler).FollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerFollowerListArgs() interface{} {
	return trust.NewTrustHandlerFollowerListArgs()
}

func newTrustHandlerFollowerListResult() interface{} {
	return trust.NewTrustHandlerFollowerListResult()
}

func followingListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerFollowingListArgs)
	realResult := result.(*trust.TrustHandlerFollowingListResult)
	success, err := handler.(trust.TrustHandler).FollowingList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerFollowingListArgs() interface{} {
	return trust.NewTrustHandlerFollowingListArgs()
}

func newTrustHandlerFollowingListResult() interface{} {
	return trust.NewTrustHandlerFollowingListResult()
}

func markToOtherHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerMarkToOtherArgs)
	realResult := result.(*trust.TrustHandlerMarkToOtherResult)
	success, err := handler.(trust.TrustHandler).MarkToOther(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerMarkToOtherArgs() interface{} {
	return trust.NewTrustHandlerMarkToOtherArgs()
}

func newTrustHandlerMarkToOtherResult() interface{} {
	return trust.NewTrustHandlerMarkToOtherResult()
}

func trustEachListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerTrustEachListArgs)
	realResult := result.(*trust.TrustHandlerTrustEachListResult)
	success, err := handler.(trust.TrustHandler).TrustEachList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerTrustEachListArgs() interface{} {
	return trust.NewTrustHandlerTrustEachListArgs()
}

func newTrustHandlerTrustEachListResult() interface{} {
	return trust.NewTrustHandlerTrustEachListResult()
}

func getUserScoreHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*trust.TrustHandlerGetUserScoreArgs)
	realResult := result.(*trust.TrustHandlerGetUserScoreResult)
	success, err := handler.(trust.TrustHandler).GetUserScore(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTrustHandlerGetUserScoreArgs() interface{} {
	return trust.NewTrustHandlerGetUserScoreArgs()
}

func newTrustHandlerGetUserScoreResult() interface{} {
	return trust.NewTrustHandlerGetUserScoreResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) IsTrust(ctx context.Context, req *trust.IsTrustRequest) (r *trust.IsTrustResponse, err error) {
	var _args trust.TrustHandlerIsTrustArgs
	_args.Req = req
	var _result trust.TrustHandlerIsTrustResult
	if err = p.c.Call(ctx, "IsTrust", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TrustAction(ctx context.Context, req *trust.FollowActionRequest) (r *trust.FollowActionResponse, err error) {
	var _args trust.TrustHandlerTrustActionArgs
	_args.Req = req
	var _result trust.TrustHandlerTrustActionResult
	if err = p.c.Call(ctx, "TrustAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, req *trust.FollowerListRequest) (r *trust.FollowerListResponse, err error) {
	var _args trust.TrustHandlerFollowerListArgs
	_args.Req = req
	var _result trust.TrustHandlerFollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowingList(ctx context.Context, req *trust.FollowingListRequest) (r *trust.FollowingListResponse, err error) {
	var _args trust.TrustHandlerFollowingListArgs
	_args.Req = req
	var _result trust.TrustHandlerFollowingListResult
	if err = p.c.Call(ctx, "FollowingList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkToOther(ctx context.Context, req *trust.MarkToOtherRequest) (r *trust.MarkToOtherResponse, err error) {
	var _args trust.TrustHandlerMarkToOtherArgs
	_args.Req = req
	var _result trust.TrustHandlerMarkToOtherResult
	if err = p.c.Call(ctx, "MarkToOther", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TrustEachList(ctx context.Context, req *trust.FriendListRequest) (r *trust.FriendListResponse, err error) {
	var _args trust.TrustHandlerTrustEachListArgs
	_args.Req = req
	var _result trust.TrustHandlerTrustEachListResult
	if err = p.c.Call(ctx, "TrustEachList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserScore(ctx context.Context, req *trust.GetUserScoreRequest) (r *trust.GetUserScoreResponse, err error) {
	var _args trust.TrustHandlerGetUserScoreArgs
	_args.Req = req
	var _result trust.TrustHandlerGetUserScoreResult
	if err = p.c.Call(ctx, "GetUserScore", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
