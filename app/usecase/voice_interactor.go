package usecase

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain"
	"nft-api/app/infrastructure"
)

type VoiceInteractor struct {
	repository infrastructure.VoiceRepository
}

func NewVoiceUsecase(voiceRepository infrastructure.VoiceRepository) VoiceUsecase {
	return VoiceInteractor{repository: voiceRepository}
}

func (interactor VoiceInteractor) Get() domain.Voice {
	return interactor.repository.Get()
}

func (interactor VoiceInteractor) Create(request request.VoiceCreateRequest) {
	interactor.repository.Create(request)
}
