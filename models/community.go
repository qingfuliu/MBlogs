package models

import "time"

type CommunityDetail struct {
	CommunityName         string    `gorm:"column:community_name" json:"community_name" binding:"required"`
	CommunityIntroduction string    `gorm:"column:community_introduction" json:"community_introduction,omitempty"`
	Creator               string    `gorm:"column:creator" json:"-"`
	CreatedAt             time.Time `gorm:"column:create_date" json:"create_date,omitempty"`
}

func (c CommunityDetail) TableName() string {
	return "community"
}

type Community struct {
	CommunityId      int64 `gorm:"column:id"`
	*CommunityDetail `gorm:"embedded"`
}

func (c Community) TableName() string {
	return "community"
}

type BatchCommunities struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Order    string `json:"order" binding:"omitempty"`
}

func (c BatchCommunities) TableName() string {
	return "community"
}
