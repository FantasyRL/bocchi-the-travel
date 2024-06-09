package service

import (
	"bocchi/kitex_gen/trust"
	"bocchi/kitex_gen/user"
	"bocchi/pkg/errno"
	"bocchi/pkg/utils/otp2fa"
	"bocchi/pkg/utils/pwd"
	"bocchi/rpc/user/dal/db"
	"bocchi/rpc/user/rpc"
)

func (s *UserService) Register(req *user.RegisterRequest) (*db.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := pwd.SetPassword(req.Password)
	userModel := &db.User{
		UserName: req.Username,
		Email:    req.Email,
		Password: PwdDigest,
	}

	return db.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginRequest) (*db.User, error) {
	userModel := &db.User{
		UserName: req.Username,
		Password: req.Password,
	}
	userResp, err := db.Login(s.ctx, userModel)
	if err != nil {
		return nil, err
	}
	if req.Otp == nil {
		return nil, errno.Verify2FAError
	}
	if userResp.Type2fa == 1 && !otp2fa.CheckTotp(*req.Otp, userResp.Otp) {
		return nil, errno.Verify2FAError
	}
	return userResp, nil
}

func (s *UserService) Info(req *user.InfoRequest) (*db.User, bool, error) {
	userModel := &db.User{
		ID: req.UserId,
	}
	userResp, err := db.QueryUserByID(s.ctx, userModel)
	if err != nil {
		return nil, false, err
	}
	rpcResp, err := rpc.IsTrust(s.ctx, &trust.IsTrustRequest{
		UserId:   req.MyId,
		TargetId: req.UserId,
	})
	if err != nil {
		return nil, false, err
	}
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, false, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	return userResp, rpcResp.IsTrust, nil
}

func (s *UserService) Signature(req *user.SignatureRequest) error {
	userModel := &db.User{
		ID:        req.UserId,
		Signature: req.Signature,
	}
	return db.PutSignature(s.ctx, userModel)
}
