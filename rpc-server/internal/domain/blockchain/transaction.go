package blockchain

import "time"

type Transaction struct {
	BlockNumber uint64
	From        string // Sender address
	To          string // Receiver address
	TxNonce     uint64
	TxHash      string
	TxValue     uint64 // amount
	TxGas       uint64
	TxGasPrice  uint64
	TxData      []byte
	Timestamp   time.Time
}
