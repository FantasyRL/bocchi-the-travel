package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Mark struct {
	ID        int64 `gorm:"primary_key"`
	Uid       int64 `gorm:"uid"`
	TargetId  int64
	Score     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

type Score struct {
	ID        int64 `gorm:"primary_key"`
	Uid       int64 `gorm:"uid"`
	Score     float64
	Count     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func CreateMark(ctx context.Context, uid int64, targetId int64, score float64) error {
	markModel := &Mark{
		Uid:      uid,
		TargetId: targetId,
		Score:    score,
	}
	return DBMark.WithContext(ctx).Create(markModel).Error
}
