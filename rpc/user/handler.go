package main

import (
	user "bibi/kitex_gen/user"
	"context"
)

// UserHandlerImpl implements the last service interface defined in the IDL.
type UserHandlerImpl struct{}

// Register implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Info implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	// TODO: Your code here...
	return
}

// Avatar implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Avatar(ctx context.Context, req *user.AvatarRequest) (resp *user.AvatarResponse, err error) {
	// TODO: Your code here...
	return
}

// Switch2FA implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Switch2FA(ctx context.Context, req *user.Switch2FARequest) (resp *user.Switch2FAResponse, err error) {
	// TODO: Your code here...
	return
}
