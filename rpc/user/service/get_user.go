package service

import (
	"bocchi/kitex_gen/user"
	"bocchi/rpc/user/dal/db"
)

func (s *UserService) GetUser(req *user.GetUsersRequest) (*[]db.User, error) {
	return db.QueryUserByIDList(s.ctx, req.UserIdList)
}
