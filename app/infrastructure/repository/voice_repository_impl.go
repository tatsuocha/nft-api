package repository

import (
	"nft-api/app/adapter/request"
	"nft-api/app/domain/entity"
	"nft-api/app/domain/model"
	"nft-api/app/domain/repository"
	"nft-api/app/infrastructure/dao/postgresql"
)

type VoiceRepositoryImpl struct {
	dao postgresql.VoiceDao
}

func NewVoiceRepository(voiceDao postgresql.VoiceDao) repository.VoiceRepository {
	return VoiceRepositoryImpl{dao: voiceDao}
}

func (repositoryImpl VoiceRepositoryImpl) Get() model.Voice {
	// MEMO: DBから取得した定の仮実装
	voiceEntity := entity.VoiceEntity{
		Id:      1,
		Name:    "sample voice name",
		Content: "sample voice",
	}

	return model.Voice{}.Build(voiceEntity.Id, voiceEntity.Name, voiceEntity.Content)
}

func (repositoryImpl VoiceRepositoryImpl) Create(request request.VoiceCreateRequest) {
	// MEMO: DBに登録する処理
}
