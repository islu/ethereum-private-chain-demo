// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres_sqlc

import (
	"time"
)

type BlockTx struct {
	Seqno       int32
	BlockNumber int64
	FromAddress string
	ToAddress   string
	TxNonce     int32
	TxHash      string
	TxValue     int64
	TxGas       int64
	TxGasPrice  int64
	TxTime      time.Time
	TxData      string
	CreateTime  time.Time
}
