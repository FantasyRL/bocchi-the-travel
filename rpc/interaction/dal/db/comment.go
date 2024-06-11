package db

import (
	"bocchi/kitex_gen/base"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

//go:generate msgp -tests=false -o=comment_msgp.go -io=false
type Comment struct {
	ID        int64          `msg:"i"`
	PoiID     int64          `msg:"v"`
	Uid       int64          `msg:"u"`
	Content   string         `msg:"c"`
	CreatedAt time.Time      `msg:"pu"`
	UpdatedAt time.Time      `msg:"-"`             //ignore
	DeletedAt gorm.DeletedAt `sql:"index" msg:"-"` //ignore
}

func IsCommentExist(ctx context.Context, commentModel *base.Comment) (bool, error) {
	var comment = &Comment{
		ID:    commentModel.Id,
		PoiID: commentModel.PoiId,
		Uid:   commentModel.User.Id,
	}
	err := DBComment.WithContext(ctx).Take(comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func CreateComment(ctx context.Context, commentModel *base.Comment) (*Comment, error) {
	var comment = &Comment{
		PoiID:   commentModel.PoiId,
		Uid:     commentModel.User.Id,
		Content: commentModel.Content,
	}
	if err := DBComment.WithContext(ctx).Create(comment).Error; err != nil {
		return &Comment{}, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, commentModel *base.Comment) (*Comment, error) {
	var comment = &Comment{
		ID:    commentModel.Id,
		PoiID: commentModel.PoiId,
		Uid:   commentModel.User.Id,
	}
	if err := DBComment.WithContext(ctx).Take(comment).Delete(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentCount(ctx context.Context, poiId int64) (count int64, err error) {
	if err = DBComment.WithContext(ctx).Where("poi_id = ?", poiId).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

func GetCommentsByPoiID(ctx context.Context, poiId int64) (*[]Comment, int64, error) {
	comments := new([]Comment)
	var count int64
	if err := DBComment.WithContext(ctx).Where("poi_id = ? AND deleted_at IS NULL", poiId).Count(&count).Find(comments).Error; err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}
