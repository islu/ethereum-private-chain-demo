package blockchain

import (
	"time"

	eth_common "github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	BlockNumber uint64
	From        eth_common.Address // Sender address
	To          eth_common.Address // Receiver address
	TxNonce     uint64
	TxHash      string
	TxValue     uint64 // amount
	TxGas       uint64
	TxGasPrice  uint64
	TxData      []byte
	Timestamp   time.Time
}
