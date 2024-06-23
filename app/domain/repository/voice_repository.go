package repository

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain/model"
)

type VoiceRepository interface {
	Get() model.Voice
	Create(request request.VoiceCreateRequest)
}
