package rpc

import (
	"bocchi/config"
	"bocchi/kitex_gen/user"
	"bocchi/kitex_gen/user/userhandler"
	"bocchi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := userhandler.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func UserRegister(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	//rpc client
	resp, err := userClient.Register(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func UserLogin(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserInfo(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserSwitch2FA(ctx context.Context, req *user.Switch2FARequest) (*user.Switch2FAResponse, error) {
	resp, err := userClient.Switch2FA(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserAvatar(ctx context.Context, req *user.AvatarRequest) (*user.AvatarResponse, error) {
	resp, err := userClient.Avatar(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
