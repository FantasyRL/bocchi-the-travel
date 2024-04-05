package db

import (
	"bocchi/kitex_gen/party"
	"bocchi/pkg/constants"
	"bocchi/pkg/errno"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Party struct {
	Id        int64
	FounderId int64
	Title     string
	Content   string
	Type      int64
	Province  string
	City      string
	StartTime time.Time
	EndTime   time.Time
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateParty(ctx context.Context, partyModel *Party) error {
	return DBParty.WithContext(ctx).Create(partyModel).Error
}

func GetPartyById(ctx context.Context, partyId int64) (*Party, error) {
	partyResp := new(Party)
	if err := DBParty.WithContext(ctx).Where("id = ?", partyId).First(partyResp).Error; err != nil {
		return nil, err
	}
	return partyResp, nil
}

func GetPartyByMultiple(ctx context.Context, req *party.SearchPartyRequest) (*[]Party, int64, error) {
	partysResp := new([]Party)
	var count int64
	if req.Content != nil {
		DBParty.Where("content LIKE ?", "%"+*req.Content+"%")
	}
	if req.PartyType != nil {
		DBParty.Where("type = ?", *req.PartyType)
	}
	if req.Province != nil {
		DBParty.Where("province = ?", *req.Province)
	}
	if req.Province != nil && req.City != nil {
		DBParty.Where("city = ?", *req.City)
	}
	if req.StartTimeDuration != nil {
		du := time.Now().Add(time.Hour * 24 * time.Duration(*req.StartTimeDuration))
		DBParty.Where("start_time > ?", du)
	}
	if req.SearchType != nil {
		switch *req.SearchType {
		case 0:
			//什么都不做
		case 1:
			//开始时间排序
			DBParty.Order("start_time DESC")
		case 2:
			//todo:热门(人多)(redis?)
		}
	}
	err := DBParty.Count(&count).
		Limit(constants.PageSize).Offset((int(req.PageNum) - 1) * constants.PageSize).
		Find(partysResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	return partysResp, count, nil
}

func GetFounderIdByPartyId(ctx context.Context, partyId int64) (int64, error) {
	partyResp := new(Party)
	if err := DBParty.WithContext(ctx).Where("party_id = ?", partyId).First(partyResp).Error; err != nil {
		return -1, err
	}
	return partyResp.FounderId, nil
}

type Member struct {
	Id        int64
	PartyId   int64
	MemberId  int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ApplyToParty(ctx context.Context, memberModel *Member) error {
	return DBMember.WithContext(ctx).Create(memberModel).Error
}

func GetMemberListByStatus(ctx context.Context, partyId int64, pageNum int64, status int64) (*[]Member, int64, error) {
	memberResp := new([]Member)
	var count int64
	err := DBMember.WithContext(ctx).Where("party_id = ? AND status = ?", partyId, status).
		Count(&count).Limit(constants.PageSize).Offset((int(pageNum) - 1) * constants.PageSize).
		Find(memberResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	return memberResp, count, nil
}

func ChangeMemberStatus(ctx context.Context, memberModel *Member) error {
	return DBMember.WithContext(ctx).Where("party_id = ? AND member_id = ?", memberModel.PartyId, memberModel.MemberId).Update("status", memberModel.Status).Error
}

func ISMemberExist(ctx context.Context, memberModel *Member) error {
	err := DBMember.WithContext(ctx).Where("party_id = ? AND member_id = ?", memberModel.PartyId, memberModel.MemberId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.MemberNotExistError
	}
	if err != nil {
		return err
	}
	return nil
}

func CheckMemberStatus(ctx context.Context, memberModel *Member) error {
	err := DBMember.WithContext(ctx).Where("party_id = ? AND member_id = ? AND status = ?", memberModel.PartyId, memberModel.MemberId, memberModel.Status).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return errno.MemberStatusDuplicateError
}
