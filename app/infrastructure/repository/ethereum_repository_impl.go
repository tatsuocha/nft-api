package repository

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"nft-api/app/domain/model"
	"nft-api/app/domain/repository"
	"nft-api/app/infrastructure/dao/ethereum"
	"nft-api/contract"
)

const rpcUrl = "http://127.0.0.1:7545"

type EthereumRepositoryImpl struct {
	dao ethereum.EthereumDao
}

func NewEthereumRepository(ethereumDao ethereum.EthereumDao) repository.EthereumRepository {
	return EthereumRepositoryImpl{dao: ethereumDao}
}

func (repositoryImpl EthereumRepositoryImpl) Get(key string, contractAddress string) model.Ethereum {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("プライベートキーの変換に失敗しました。: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("パブリックキーへのキャストに失敗しました。: %v", err)
	}

	client := repositoryImpl.dao.Connect(rpcUrl)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("ナンスの取得に失敗しました。: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("ガス代の取得に失敗しました。: %v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("チェインIDの取得に失敗しました。: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("認証トランザクションの生成に失敗しました。: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	// 送金する場合のETHの量
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	// インスタンス化
	hexContractAddress := common.HexToAddress(contractAddress)
	instance, err := contract.NewEthereum(hexContractAddress, client)
	if err != nil {
		log.Fatalf("イーサリアムスマートコントラクトのインスタンス化に失敗しました。: %v", err)
	}

	// Ethereumの値を取得
	result := repositoryImpl.dao.Get(instance)

	return *model.Ethereum{}.Build(result)
}
