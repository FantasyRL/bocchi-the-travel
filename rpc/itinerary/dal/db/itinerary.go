package db

import (
	"bocchi/kitex_gen/itinerary"
	"bocchi/pkg/errno"
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Itinerary struct {
	Id                int64
	Title             string
	FounderId         int64
	PartyId           int64
	ActionType        int64
	Rectangle         *string
	RouteJson         *string
	Remark            *string
	Sequence          int64
	ScheduleStartTime string
	ScheduleEndTime   string
	IsMerged          int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func CreateItinerary(ctx context.Context, itineraryModel *Itinerary) (*Itinerary, error) {
	itineraryResp := new(Itinerary)
	if err := DB.WithContext(ctx).Create(itineraryModel).First(itineraryResp).Error; err != nil {
		return nil, err
	}
	return itineraryResp, nil
}

func ShowPartyItinerary(ctx context.Context, partyId int64) (*[]Itinerary, int64, error) {
	itinerariesResp := new([]Itinerary)
	var count int64
	if err := DB.WithContext(ctx).Where("party_id = ? and is_merged = 1", partyId).Order("sequence ASC").Count(&count).Find(itinerariesResp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return itinerariesResp, count, nil
}

func ChangeSequence(ctx context.Context, req *itinerary.ChangeSequenceRequest) error {
	itineraries := make([]Itinerary, len(req.ItineraryIdList))
	for i := 0; i < len(req.ItineraryIdList); i++ {
		itineraries[i] = Itinerary{
			Id:       req.ItineraryIdList[i],
			Sequence: req.SequenceList[i],
		}
	}
	return DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sequence"}),
	}).Create(&itineraries).Error

}

func ChangeItineraryStatus(ctx context.Context, partyId int64, itineraryId int64, status int64) error {
	itineraryResp := new(Itinerary)
	if err := DB.WithContext(ctx).Where("id = ? AND party_id = ?", itineraryId, partyId).First(itineraryResp).Error; err != nil {
		return err
	}
	if itineraryResp.IsMerged == status {
		return errno.ParamError
	}
	return DB.WithContext(ctx).Model(itineraryResp).Update("is_merged", status).Error
}

func GetItineraryById(ctx context.Context, itineraryId int64) (*Itinerary, error) {
	itineraryResp := new(Itinerary)
	if err := DB.WithContext(ctx).Where("id = ?", itineraryId).First(itineraryResp).Error; err != nil {
		return nil, err
	}
	return itineraryResp, nil
}

func GetItinerariesByUidAndPartyId(ctx context.Context, userId int64, partyId int64) (*[]Itinerary, int64, error) {
	itinerariesResp := new([]Itinerary)
	var count int64
	if err := DB.WithContext(ctx).Where("founder_id = ? AND party_id = ?", userId, partyId).Order("is_merged ASC").Count(&count).Find(itinerariesResp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return itinerariesResp, count, nil
}

func DeleteItinerary(ctx context.Context, itineraryId int64) error {
	return DB.WithContext(ctx).Delete(&Itinerary{Id: itineraryId}).Error
}
