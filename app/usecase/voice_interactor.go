package usecase

import (
	"nft-api/app/domain/model"
	"nft-api/app/domain/repository"
	"nft-api/app/domain/usecase"
)

type VoiceInteractor struct {
	repository repository.VoiceRepository
}

func NewVoiceUseCase(voiceRepository repository.VoiceRepository) usecase.VoiceUseCase {
	return VoiceInteractor{repository: voiceRepository}
}

func (interactor VoiceInteractor) Get() *[]model.Voice {
	return interactor.repository.Get()
}

func (interactor VoiceInteractor) Create(name *string, content *string) {
	interactor.repository.Create(name, content)
}
