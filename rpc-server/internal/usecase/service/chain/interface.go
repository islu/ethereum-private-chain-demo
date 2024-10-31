package chain

import (
	"context"
	"math/big"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/blockchain"
)

type TransactionRepository interface {
	// Create block transaction
	CreateBlockTx(ctx context.Context, params blockchain.Transaction) (*blockchain.Transaction, error)
	// Get block transaction by tx_hash
	GetBlockTxByTxHash(ctx context.Context, txHash string) (*blockchain.Transaction, error)
	// Get max block number by from_address
	GetMaxBlockNumberByFromAddress(ctx context.Context, from string) (int64, error)
	// List block transaction
	ListBlockTx(ctx context.Context, size int) ([]blockchain.Transaction, error)
	// List block transaction by from_address
	ListBlockTxByFromAddress(ctx context.Context, size int, from string) ([]blockchain.Transaction, error)
}

type EthereumNodeClient interface {
	// Get the latest block number
	GetLatestBlockNumber() (uint64, error)
	// Get account balance
	GetBalance(account eth_common.Address) (*big.Int, error)
	// Send transaction (Simulate)
	SendTx_Simulate(amountToSend *big.Int, to eth_common.Address) error
	// 取得指定區塊的交易
	GetBlockTransactionsByNumber(blockNumber int64) ([]blockchain.Transaction, error)
	// Get transaction count
	GetTransactionCount(account eth_common.Address, blockNumber int64) (uint64, error)
}
