package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChildChain struct {
	con          *ethclient.Client
	contractAddr common.Address
}

func NewChildChain(contractAddr common.Address) *ChildChain {
	return &ChildChain{
		contractAddr: contractAddr,
	}
}

func (c *ChildChain) Dial(rawURL string) error {
	con, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}

	c.con = con
	return nil
}
