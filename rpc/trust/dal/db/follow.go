package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	ID         int64 `gorm:"primary_key"`
	Uid        int64 `gorm:"uid"`
	FollowedId int64 `gorm:"followed_id"`
	Status     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
}

func IsFollowExist(ctx context.Context, uid int64, followerId int64) (bool, error) {
	follow := new(Follow)
	err := DB.WithContext(ctx).Where("uid = ? AND followed_id = ? ", uid, followerId).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsFollowStatus(ctx context.Context, uid int64, followerId int64, status int64) (bool, error) {
	follow := new(Follow)
	err := DB.WithContext(ctx).Where("uid = ? AND followed_id = ? AND status = ? ", uid, followerId, status).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateFollow(ctx context.Context, uid int64, followerId int64, status int64) error {
	followModel := &Follow{
		Uid:        uid,
		FollowedId: followerId,
		Status:     status,
	}
	return DB.WithContext(ctx).Create(followModel).Error
}

func UpdateFollowStatus(ctx context.Context, uid int64, followerId int64, status int64) error {
	return DB.WithContext(ctx).Where("uid = ? AND followed_id = ? ",
		uid, followerId).Update("status", status).Error
}

// FollowerAndFriendList : followerList friendList count error
func FollowerAndFriendList(ctx context.Context, followedId int64) (*[]Follow, *[]Follow, int64, error) {
	followerList := new([]Follow)
	friendList := new([]Follow)
	var count int64
	if err := DB.WithContext(ctx).Where("followed_id = ? AND status = 1", followedId).Find(followerList).Count(&count).Error; err != nil {
		return nil, nil, 0, err
	}

	isFriendIdList := make([]int64, len(*followerList))
	for i, v := range *followerList {
		isFriendIdList[i] = v.Uid
	}
	if err := DB.WithContext(ctx).Where("uid = ? AND followed_id IN ? AND status = 1", followedId, isFriendIdList).Find(friendList).Error; err != nil {
		return nil, nil, 0, err
	}
	return followerList, friendList, count, nil
}

func FollowingList(ctx context.Context, uid int64) (*[]Follow, int64, error) {
	followedList := new([]Follow)
	var count int64
	if err := DB.WithContext(ctx).Where("uid = ? AND status = 1", uid).Count(&count).Find(followedList).Error; err != nil {
		return nil, 0, err
	}
	return followedList, count, nil
}

func IsFollow(ctx context.Context, uid int64, followedId int64) (bool, error) {
	follow := new(Follow)
	err := DB.WithContext(ctx).Where("uid = ? AND followed_id = ?", uid, followedId).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
