package repository

import (
	"nft-api/app/domain/model"
)

type EthereumRepository interface {
	Get(key string, contractAddress string) model.Ethereum
}
