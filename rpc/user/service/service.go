package service

import (
	"bocchi/kitex_gen/user"
	aliyunoss "bocchi/pkg/utils/oss"
	"bocchi/rpc/user/dal/db"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type UserService struct {
	ctx context.Context
}

type AvatarService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func NewAvatarService(ctx context.Context) *AvatarService {
	bucket, err := aliyunoss.OSSBucketCreate()
	if err != nil {
		log.Fatal(err)
	}
	return &AvatarService{ctx: ctx, bucket: bucket}
}

func BuildUserResp(dbUser *db.User) *user.User {
	return &user.User{
		Id:        dbUser.ID,
		Name:      dbUser.UserName,
		Email:     dbUser.Email,
		Avatar:    dbUser.Avatar,
		Signature: dbUser.Signature,
	}
}

func BuildUsersResp(dbUsers *[]db.User) []*user.User {
	userResp := make([]*user.User, len(*dbUsers))
	for i, dbUser := range *dbUsers {
		userResp[i] = BuildUserResp(&dbUser)
	}
	return userResp
}
