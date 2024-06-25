package repository

import (
	"nft-api/app/domain/model"
)

type VoiceRepository interface {
	Get() *[]model.Voice
	Create(name *string, content *string)
}
