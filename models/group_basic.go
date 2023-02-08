package models

import "gorm.io/gorm"

// 群消息
type GroupBasic struct {
	gorm.Model
	Name    string `json:"name"`
	OwnerId uint   `json:"ownerId"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
	Type    int    `json:"type"`
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
