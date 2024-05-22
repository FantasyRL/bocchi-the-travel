package db

import (
	"gorm.io/gorm"
	"time"
)

type Itinerary struct {
	Id                int64
	Title             string
	FounderId         int64
	PartyId           int64
	ActionType        int64
	Rectangle         string
	RouteJson         string
	Remark            string
	Sequence          int64
	ScheduleStartTime string
	ScheduleEndTime   string
	IsMerged          int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
