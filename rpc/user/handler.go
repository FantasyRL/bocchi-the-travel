package main

import (
	"bocchi/config"
	"bocchi/kitex_gen/user"
	"bocchi/kitex_gen/user/userhandler"
	"bocchi/pkg/constants"
	"bocchi/pkg/errno"
	"bocchi/pkg/pack"
	"bocchi/rpc/user/dal/db"
	"bocchi/rpc/user/service"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"golang.org/x/sync/errgroup"
)

// UserHandlerImpl implements the last service interface defined in the IDL.
type UserHandlerImpl struct {
	userCli userhandler.Client
}

func NewUserClient(addr string) (userhandler.Client, error) {
	return userhandler.NewClient(constants.UserServiceName, client.WithHostPorts(addr))
}

// Register implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	//不开辟空间就会死
	resp = new(user.RegisterResponse)

	userResp, err := service.NewUserService(ctx).Register(req)

	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}

	resp.UserId = &userResp.ID

	return
}

// Login implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	userResp, err := service.NewUserService(ctx).Login(req)

	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.User = service.BuildUserResp(userResp)
	return
}

// Info implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = new(user.InfoResponse)

	userResp, err := service.NewUserService(ctx).Info(req)

	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.User = service.BuildUserResp(userResp)
	return
}

// Avatar implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Avatar(ctx context.Context, req *user.AvatarRequest) (resp *user.AvatarResponse, err error) {
	resp = new(user.AvatarResponse)

	//开启并发
	var eg errgroup.Group
	//上传至OSS
	eg.Go(func() error { //同时Add(1)
		err = service.NewAvatarService(ctx).UploadAvatar(req)
		if err != nil {
			return errno.UploadFileError
		}
		return nil
	})
	//上传url至数据库
	UserResp := new(db.User)
	eg.Go(func() error { //Add(1)
		avatarUrl := fmt.Sprintf("%s/%s/%d", config.OSS.EndPoint, config.OSS.MainDirectory, req.UserId)
		UserResp, err = service.NewAvatarService(ctx).PutAvatar(req.UserId, avatarUrl)
		if err != nil {
			return err
		}
		return nil
	})
	//Wait实现了错误处理与sync，仅返回第一个发生的错误
	if err = eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.User = service.BuildUserResp(UserResp)
	return
}

// Switch2FA implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Switch2FA(ctx context.Context, req *user.Switch2FARequest) (resp *user.Switch2FAResponse, err error) {
	resp = new(user.Switch2FAResponse)

	switch req.ActionType {
	case 0, 1:
		err = service.NewUserService(ctx).Switch2faType(req)
		resp.Base = pack.BuildBaseResp(err)
	default:
		resp.Base = pack.BuildBaseResp(errno.ParamError)
	}
	return resp, nil
}

// Signature implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Signature(ctx context.Context, req *user.SignatureRequest) (resp *user.SignatureResponse, err error) {
	resp = new(user.SignatureResponse)

	err = service.NewUserService(ctx).Signature(req)
	resp.Base = pack.BuildBaseResp(err)

	return resp, nil
}
