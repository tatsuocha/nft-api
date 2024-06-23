package model

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Ethereum struct {
	Context     context.Context
	From        common.Address
	BlockNumber *big.Int
	BlockHash   common.Hash
	Pending     bool
}

func (ethereum Ethereum) Build(callOpts *bind.CallOpts) *Ethereum {
	return &Ethereum{
		Context:     callOpts.Context,
		From:        callOpts.From,
		BlockNumber: callOpts.BlockNumber,
		BlockHash:   callOpts.BlockHash,
		Pending:     callOpts.Pending,
	}
}
