package usecase

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain/model"
)

type VoiceUseCase interface {
	Get() model.Voice
	Create(request request.VoiceCreateRequest)
}
