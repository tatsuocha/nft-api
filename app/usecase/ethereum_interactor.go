package usecase

import (
	"nft-api/app/domain/model"
	"nft-api/app/domain/repository"
	"nft-api/app/domain/usecase"
)

const key = `22aabb811efca4e6f4748bd18a46b502fa85549df9fa07da649c0a148d7d5530`
const contractAddress = "0x8Ad287f900D8Dc07583Da5AE1aE427001479767D"

type EthereumInteractor struct {
	repository repository.EthereumRepository
}

func NewEthereumUseCase(ethereumRepository repository.EthereumRepository) usecase.EthereumUseCase {
	return EthereumInteractor{repository: ethereumRepository}
}

func (interactor EthereumInteractor) Get() model.Ethereum {
	return interactor.repository.Get(key, contractAddress)
}
