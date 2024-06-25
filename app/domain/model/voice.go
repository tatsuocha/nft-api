package model

import (
	"nft-api/app/domain/entity"
	"time"
)

type Voice struct {
	Id              uint
	UserId          int64
	Name            string
	Content         string
	CreatedDatetime time.Time
}

func (voice Voice) Build(voiceEntity *entity.VoiceEntity) Voice {
	return Voice{
		Id:              voiceEntity.Id,
		UserId:          voiceEntity.UserId,
		Name:            voiceEntity.Name,
		Content:         voiceEntity.Content,
		CreatedDatetime: voiceEntity.CreatedDatetime,
	}
}
