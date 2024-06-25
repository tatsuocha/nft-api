package entity

import (
	"time"
)

type VoiceEntity struct {
	Id              uint `gorm:"primaryKey;autoIncrement"`
	UserId          int64
	Name            string
	Content         string
	CreatedDatetime time.Time `gorm:"autoCreateTime"`
}

func (voiceEntity VoiceEntity) BuildForInsert(userId int64, name string, content string) *VoiceEntity {
	return &VoiceEntity{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
