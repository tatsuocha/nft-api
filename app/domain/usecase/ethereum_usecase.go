package usecase

import (
	"nft-api/app/domain/model"
)

type EthereumUseCase interface {
	Get() model.Ethereum
}
