package service

import (
	"bocchi/kitex_gen/user"
	"bocchi/rpc/user/dal/db"
)

func (s *UserService) GetMember(req *user.GetMemberRequest) (*[]db.User, error) {
	return db.QueryUserByIDList(req.MemberIdList)
}
