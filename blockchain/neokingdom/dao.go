package neokingdom

import (
	"log"
	"math/big"

	"github.com/TelediskoDAO/oracle/blockchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type DAO struct {
	Wallet *blockchain.Wallet
	Tokens common.Address
}

func NewDAO(tokens common.Address) *DAO {
	return &DAO{
		Tokens: tokens,
	}
}

func (a *DAO) Connect(w *blockchain.Wallet) {
	a.Wallet = w
}

func (a *DAO) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	//tt, err := NewTelediskoTaler(a.Tokens, a.Wallet.GetClient())
	tt, err := NewTelediskoToken(a.Tokens, a.Wallet.GetClient())
	if err != nil {
		return nil, err
	}

	opts, err := a.Wallet.NewSigner()
	if err != nil {
		return nil, err
	}
	opts.NoSend = false

	tx, err := tt.Mint(opts, to, amount)

	log.Println(tx.Data())

	if err != nil {
		return nil, err
	}
	return tx, nil
}
