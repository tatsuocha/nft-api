package infrastructure

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain"
)

type VoiceRepository interface {
	Get() *domain.Voice
	Create(request request.VoiceCreateRequest)
}
