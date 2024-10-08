//MVC--Model

package db

import (
	"bocchi/pkg/errno"
	"bocchi/pkg/utils/pwd"
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64
	UserName  string
	Email     string
	Password  string
	Avatar    string
	Signature string
	Otp       string
	Type2fa   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func Register(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	//WithContext(ctx)是将一个context.Context对象和数据库连接绑定，以实现在数据库操作中使用context.Context上下文传递。
	if err := DB.WithContext(ctx).Where("user_name = ? OR email = ?", userModel.UserName, userModel.Email).First(&userResp).Error; err == nil {
		return nil, errno.ExistUserError
	}

	if err := DB.WithContext(ctx).Create(userModel).Error; err != nil {
		return nil, err
	}
	return userModel, nil
}

func Login(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := DB.WithContext(ctx).Where("user_name = ?", userModel.UserName).
		First(&userResp).Error; err != nil {
		return nil, errno.NotExistUserError
	}

	if pwd.CheckPassword(userResp.Password, userModel.Password) == false {
		return nil, errno.PwdError
	}

	return userResp, nil
}

func Update2FAType(ctx context.Context, type2fa int64, uid int64) error {
	return DB.WithContext(ctx).Where("id = ?", uid).Update("type2fa", type2fa).Error
}

func Update2FA(ctx context.Context, totp string, uid int64) error {
	return DB.WithContext(ctx).Where("id = ?", uid).Update("otp", totp).Error
}

func PutAvatar(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := DB.WithContext(ctx).Where("id = ?", userModel.ID).Update("avatar", userModel.Avatar).Error; err != nil {
		return nil, err
	}
	if err := DB.WithContext(ctx).Where("id = ?", userModel.ID).First(userResp).Error; err != nil {
		return nil, err
	}
	return userResp, nil
}

func QueryUserByID(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := DB.WithContext(ctx).Where("id = ?", userModel.ID).First(&userResp).Error; err != nil {
		return nil, err
	}
	return userResp, nil
}

func QueryUserByIDList(ctx context.Context, uidList []int64) (*[]User, error) {
	dbResp := new([]User)
	if err := DB.WithContext(ctx).Where("id IN ?", uidList).Find(dbResp).Error; err != nil {
		return nil, err
	}
	userResp := make([]User, 0)
	for _, id := range uidList {
		for _, v := range *dbResp {
			if v.ID == id {
				userResp = append(userResp, v)
			}
		}
	}
	return &userResp, nil
}

func PutSignature(ctx context.Context, userModel *User) error {
	return DB.WithContext(ctx).Where("id = ?", userModel.ID).Updates(User{Signature: userModel.Signature}).Error
}
