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

func CreateScore(ctx context.Context, uid int64, score float64) error {
	scoreModel := &Score{
		Uid:   uid,
		Score: score,
		Count: 1,
	}
	return DBScore.WithContext(ctx).Create(scoreModel).Error
}

func UpdateScore(ctx context.Context, uid int64, score float64) error {
	scoreModel := new(Score)
	if err := DBScore.WithContext(ctx).Where("uid = ?", uid).Find(scoreModel).Error; err != nil {
		return err
	}
	newScore := (scoreModel.Score + score) / float64(scoreModel.Count+1)
	scoreModel.Count++
	scoreModel.Score = newScore
	return DBScore.WithContext(ctx).Save(scoreModel).Error
}
