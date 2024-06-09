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
	Type      string
	Province  string
	City      string
	StartTime time.Time
	EndTime   time.Time
	Status    int64
	Rectangle string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateParty(ctx context.Context, partyModel *Party) (*Party, error) {
	//partyResp := new(Party)
	if err := DBParty.WithContext(ctx).Create(partyModel).Error; err != nil {
		return nil, err
	}

	memberModel := &Member{
		PartyId:  partyModel.Id,
		MemberId: partyModel.FounderId,
		Status:   2,
	}
	if err := DBMember.WithContext(ctx).Create(memberModel).Error; err != nil {
		return nil, err
	}
	return partyModel, nil
}

func GetPartyById(ctx context.Context, partyId int64) (*Party, error) {
	partyResp := new(Party)
	if err := DBParty.WithContext(ctx).Where("id = ?", partyId).First(partyResp).Error; err != nil {
		return nil, err
	}
	return partyResp, nil
}

func GetPartyByMultiple(ctx context.Context, req *party.SearchPartyRequest) (*[]Party, int64, error) {
	partiesResp := new([]Party)
	var count int64
	dbq := DBParty
	if req.Content != nil && *req.Content != "nothing" {
		dbq = dbq.WithContext(ctx).Where("content LIKE ? OR title LIKE ?", "%"+*req.Content+"%", "%"+*req.Content+"%")
	}
	if req.PartyType != nil && *req.PartyType != "nothing" {
		dbq = dbq.WithContext(ctx).Where("type = ?", *req.PartyType)
	}
	if req.Province != nil && *req.Province != "nothing" {
		dbq = dbq.WithContext(ctx).Where("province = ?", *req.Province)
	}
	if req.Province != nil && req.City != nil && *req.Province != "nothing" && *req.City != "nothing" {
		dbq = dbq.WithContext(ctx).Where("city = ?", *req.City)
	}
	if req.StartTimeDuration != nil && *req.StartTimeDuration != -1 {
		du := time.Now().Add(time.Hour * 24 * time.Duration(*req.StartTimeDuration))
		dbq = dbq.WithContext(ctx).Where("start_time > ?", du)
	} else {
		dbq = dbq.WithContext(ctx).Where("status = 0")
	}
	if req.SearchType != nil && *req.SearchType != -1 {
		switch *req.SearchType {
		case 0:
			//什么都不做
		case 1:
			//开始时间排序
			dbq = dbq.Order("start_time DESC")
		case 2:
			//todo:热门(人多)(redis?)
		}
	}
	err := dbq.WithContext(ctx).Count(&count).
		Limit(constants.PageSize).Offset((int(req.PageNum) - 1) * constants.PageSize).
		Find(partiesResp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	return partiesResp, count, nil
}

func GetFounderIdByPartyId(ctx context.Context, partyId int64) (int64, error) {
	partyResp := new(Party)
	if err := DBParty.WithContext(ctx).Where("id = ?", partyId).First(partyResp).Error; err != nil {
		return -1, err
	}
	return partyResp.FounderId, nil
}

func GetPartiesById(ctx context.Context, memberId int64, pageNum int64) (*[]Party, int64, error) {
	partyIdResp := new([]Member)
	var count int64
	err := DBMember.WithContext(ctx).Where("member_id = ?", memberId).Order("status DESC").
		Count(&count).Limit(constants.PageSize).Offset((int(pageNum) - 1) * constants.PageSize).
		Find(partyIdResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	partyIdList := make([]int64, count)
	for i, member := range *partyIdResp {
		partyIdList[i] = member.PartyId
	}
	partiesResp := make([]Party, 0)
	partiesResp1 := new([]Party)
	partiesResp2 := new([]Party)
	err = DBParty.WithContext(ctx).Where("id IN ? and status = 0", partyIdList).Order("start_time ASC").
		Find(partiesResp1).Error
	err = DBParty.WithContext(ctx).Where("id IN ? and status = 1", partyIdList).Order("end_time DESC").
		Find(partiesResp2).Error
	for _, p := range *partiesResp1 {
		partiesResp = append(partiesResp, p)
	}
	for _, p := range *partiesResp2 {
		partiesResp = append(partiesResp, p)
	}

	return &partiesResp, count, nil
}

func DeleteParty(ctx context.Context, partyId int64) error {
	return DBParty.WithContext(ctx).Delete(&Party{Id: partyId}).Error
}
func FinishParty(ctx context.Context, partyId int64) error {
	partyMid := new(Party)
	if err := DBParty.WithContext(ctx).Where("id = ?", partyId).First(partyMid).Error; err != nil {
		return err
	}
	if partyMid.Status != 0 {
		return errno.ParamError
	}

	return DBParty.WithContext(ctx).Where("id = ?", partyId).Update("status", 1).Error
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

func GetPartyMembers(ctx context.Context, partyId int64, pageNum int64) (*[]Member, int64, error) {
	memberResp := new([]Member)
	actionType := make([]int64, 2)
	actionType[0] = 1
	actionType[1] = 2
	var count int64
	err := DBMember.WithContext(ctx).Where("party_id = ? AND status IN ?", partyId, actionType).
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

func ISMemberExist(ctx context.Context, memberModel *Member) (bool, error) {
	memberResp := new(Member)
	err := DBMember.WithContext(ctx).Where("party_id = ? AND member_id = ?", memberModel.PartyId, memberModel.MemberId).First(memberResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errno.MemberNotExistError
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func ISMemberAdmin(ctx context.Context, memberModel *Member) (bool, error) {
	memberResp := new(Member)
	err := DBMember.WithContext(ctx).Where("party_id = ? AND member_id = ? AND status = 2", memberModel.PartyId, memberModel.MemberId).First(memberResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errno.MemberNotExistError
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckMemberStatus(ctx context.Context, memberModel *Member) error {
	memberResp := new(Member)
	err := DBMember.WithContext(ctx).Where("party_id = ? AND member_id = ? AND status = ?", memberModel.PartyId, memberModel.MemberId, memberModel.Status).First(memberResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return errno.MemberStatusDuplicateError
}

func DeleteMember(ctx context.Context, memberId int64) error {
	return DBMember.WithContext(ctx).Delete(&Member{Id: memberId}).Error
}
