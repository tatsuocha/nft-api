package infrastructure

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain"
	"nft-api/app/domain/entity"
)

type VoiceRepositoryImpl struct{}

func NewVoiceRepository() VoiceRepository {
	return VoiceRepositoryImpl{}
}

func (v VoiceRepositoryImpl) Get() domain.Voice {
	// MEMO: DBから取得した定の仮実装
	voiceEntity := entity.VoiceEntity{
		Id:      1,
		Name:    "sample voice name",
		Content: "sample voice",
	}

	return domain.Voice{}.Build(voiceEntity.Id, voiceEntity.Name, voiceEntity.Content)
}

func (v VoiceRepositoryImpl) Create(request request.VoiceCreateRequest) {
	// MEMO: DBに登録する処理
}
