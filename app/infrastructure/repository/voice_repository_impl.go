package repository

import (
	"log"
	"nft-api/app/domain/entity"
	"nft-api/app/domain/model"
	"nft-api/app/domain/repository"
	"nft-api/app/infrastructure/dao/postgresql"
)

type VoiceRepositoryImpl struct {
	voiceDao  postgresql.VoiceDao
	commonDao postgresql.CommonDao
}

func NewVoiceRepository(voiceDao postgresql.VoiceDao, commonDao postgresql.CommonDao) repository.VoiceRepository {
	return VoiceRepositoryImpl{voiceDao: voiceDao, commonDao: commonDao}
}

func (repositoryImpl VoiceRepositoryImpl) Get() *[]model.Voice {
	db := repositoryImpl.commonDao.Connect()

	var voiceEntities []entity.VoiceEntity
	result := db.Find(&voiceEntities)
	if result.Error != nil {
		log.Fatalf("voiceの取得に失敗しました。 %v", result.Error)
	}

	return toVoices(voiceEntities)
}

func (repositoryImpl VoiceRepositoryImpl) Create(name string, content string) {
	db := repositoryImpl.commonDao.Connect()

	result := db.Create(entity.VoiceEntity{}.BuildForInsert(1, name, content))
	if result.Error != nil {
		log.Fatalf("voiceの登録に失敗しました。 %v", result.Error)
	}
}

func toVoices(voiceEntities []entity.VoiceEntity) *[]model.Voice {
	var voices []model.Voice
	for _, voiceEntity := range voiceEntities {
		voice := model.Voice{}.Build(int(voiceEntity.Id), voiceEntity.Name, voiceEntity.Content)
		voices = append(voices, voice)
	}

	return &voices
}
