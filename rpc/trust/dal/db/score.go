package db

import (
	"context"
	"errors"
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

func IsScoreExist(ctx context.Context, uid int64) (bool, error) {
	scoreModel := new(Score)
	err := DBScore.WithContext(ctx).Where("uid = ?", uid).First(scoreModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateScore(ctx context.Context, uid int64, score float64) error {
	scoreModel := new(Score)
	if err := DBScore.WithContext(ctx).Where("uid = ?", uid).First(scoreModel).Error; err != nil {
		return err
	}
	newScore := scoreModel.Score*float64(scoreModel.Count/(scoreModel.Count+1)) + score/float64(scoreModel.Count+1)
	scoreModel.Count++
	scoreModel.Score = newScore
	return DBScore.WithContext(ctx).Save(scoreModel).Error
}

func GetScoreByUId(ctx context.Context, uid int64) (*Score, error) {
	scoreResp := new(Score)
	if err := DBScore.WithContext(ctx).Where("uid = ?", uid).Find(scoreResp).Error; err != nil {
		return nil, err
	}
	return scoreResp, nil
}
