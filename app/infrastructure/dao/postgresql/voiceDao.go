package postgresql

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"nft-api/contract"
)

type VoiceDao struct{}

func NewVoiceDao() VoiceDao {
	return VoiceDao{}
}

func (dao VoiceDao) GetAll(instance *contract.Ethereum) *bind.CallOpts {
	opts := &bind.CallOpts{}
	_, err := instance.Get(opts)
	if err != nil {
		log.Fatalf("イーサリアムの取得に失敗しました。: %v", err)
	}
	return opts
}
