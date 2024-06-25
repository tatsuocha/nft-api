package usecase

import (
	"nft-api/app/domain/model"
)

type VoiceUseCase interface {
	Get() *[]model.Voice
	Create(name string, content string)
}
