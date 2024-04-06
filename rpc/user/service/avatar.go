package service

import (
	"bocchi/config"
	"bocchi/kitex_gen/user"
	"bocchi/rpc/user/dal/db"
	"bytes"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

func (s *AvatarService) UploadAvatar(req *user.AvatarRequest) error {
	avatarReader := bytes.NewReader(req.AvatarFile)
	err := s.bucket.PutObject(config.OSS.MainDirectory+"/"+strconv.FormatInt(req.UserId, 10)+".jpg", avatarReader)
	if err != nil {
		klog.Errorf("upload file error:%v\n", err)
	}
	return err
}

func (s *AvatarService) PutAvatar(id int64, avatarUrl string) (*db.User, error) {
	userModel := &db.User{
		ID:     id,
		Avatar: avatarUrl,
	}
	return db.PutAvatar(s.ctx, userModel)
}
