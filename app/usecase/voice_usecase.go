package usecase

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain"
)

type VoiceUsecase interface {
	Get() *domain.Voice
	Create(request request.VoiceCreateRequest)
}
