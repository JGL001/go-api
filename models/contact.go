package models

import (
	"ginChat/utils"

	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint   `json:"formId"`   // 谁的关系信息
	TargetId int64  `json:"targetId"` // 对应的谁
	ChatType string `json:"chatType"` // 聊天类型 群聊 私聊 广播
	Desc     string `json:"desc"`     // 描述
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and chat_type=1", userId).Find(&contacts)

	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}
