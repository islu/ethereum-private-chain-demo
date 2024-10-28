package chain

import (
	"math/big"

	eth_common "github.com/ethereum/go-ethereum/common"
)

type TransactionRepository interface {
}

type EthereumNodeClient interface {
	// Get the latest block number
	GetLatestBlockNumber() (uint64, error)
	// Get account balance
	GetBalance(account eth_common.Address) (*big.Int, error)
	// Send transaction (Simulate)
	SendTx_Simulate(amountToSend *big.Int, to eth_common.Address) error
}
