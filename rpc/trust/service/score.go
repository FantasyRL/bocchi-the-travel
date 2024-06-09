package service

import (
	"bocchi/kitex_gen/trust"
	"bocchi/rpc/trust/dal/db"
)

func (s *FollowService) MarkToOther(req *trust.MarkToOtherRequest) error {
	if err := db.CreateMark(s.ctx, req.UserId, req.TargetId, req.Score); err != nil {
		return err
	}
	is, err := db.IsScoreExist(s.ctx, req.UserId)
	if err != nil {
		return err
	}
	if !is {
		return db.CreateScore(s.ctx, req.UserId, req.Score)
	} else {
		return db.UpdateScore(s.ctx, req.UserId, req.Score)
	}
}

func (s *FollowService) GetUserScore(req *trust.GetUserScoreRequest) (float64, int64, error) {
	score, err := db.GetScoreByUId(s.ctx, req.UserId)
	if err != nil {
		return 0, 0, err
	}
	return score.Score, score.Count, nil
}
