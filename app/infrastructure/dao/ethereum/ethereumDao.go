package ethereum

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"nft-api/contract"
)

type EthereumDao struct{}

func NewEthereumDao() EthereumDao {
	return EthereumDao{}
}

func (dao EthereumDao) Connect(rpcUrl string) *ethclient.Client {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("イーサリアムとの接続に失敗しました。: %v", err)
	}
	return client
}

func (dao EthereumDao) Get(instance *contract.Ethereum) *bind.CallOpts {
	opts := &bind.CallOpts{}
	_, err := instance.Get(opts)
	if err != nil {
		log.Fatalf("イーサリアムの取得に失敗しました。: %v", err)
	}
	return opts
}
