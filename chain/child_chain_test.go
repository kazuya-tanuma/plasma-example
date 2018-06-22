package chain

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewChildChain(t *testing.T) {
	// contAddr := os.Getenv("CONTRACT_ADDRESS")
	contAddr := "0x0f42e3ae62e5565Bab369719433C06bC913e750b"
	childChain := NewChildChain(common.HexToAddress(contAddr))
	assert.Equal(t, childChain.contractAddr.Hex(), contAddr)

}
