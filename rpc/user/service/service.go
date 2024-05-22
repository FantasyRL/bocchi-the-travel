package service

import (
	"bocchi/kitex_gen/base"
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

func BuildUserResp(dbUser *db.User) *base.User {
	return &base.User{
		Id:        dbUser.ID,
		Name:      dbUser.UserName,
		Email:     dbUser.Email,
		Avatar:    dbUser.Avatar,
		Signature: dbUser.Signature,
	}
}

func BuildUsersResp(dbUsers *[]db.User) []*base.User {
	userResp := make([]*base.User, len(*dbUsers))
	for i, dbUser := range *dbUsers {
		userResp[i] = BuildUserResp(&dbUser)
	}
	return userResp
}
