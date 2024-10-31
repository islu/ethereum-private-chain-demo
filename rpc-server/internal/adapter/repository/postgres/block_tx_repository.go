package postgres

import (
	"context"
	"time"

	eth_common "github.com/ethereum/go-ethereum/common"
	psqlc "github.com/islu/ethereum_private_chain/rpc_server/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/blockchain"
)

// Create block transaction
func (r *PostgresRepository) CreateBlockTx(ctx context.Context, params blockchain.Transaction) (*blockchain.Transaction, error) {

	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	q := psqlc.New(r.connPool)
	qtx := q.WithTx(tx)

	blockTx, err := qtx.CreateBlockTx(ctx, psqlc.CreateBlockTxParams{
		BlockNumber: int64(params.BlockNumber),
		FromAddress: params.From.Hex(),
		ToAddress:   params.To.Hex(),
		TxNonce:     int32(params.TxNonce),
		TxHash:      params.TxHash,
		TxValue:     int64(params.TxValue),
		TxGas:       int64(params.TxGas),
		TxGasPrice:  int64(params.TxGasPrice),
		TxTime:      params.Timestamp,
		TxData:      string(params.TxData),
		CreateTime:  time.Now(),
	})
	if err != nil {
		return nil, err
	}

	result := toTransaction(blockTx)

	return &result, tx.Commit(ctx)
}

// Get block transaction by tx_hash
func (r *PostgresRepository) GetBlockTxByTxHash(ctx context.Context, txHash string) (*blockchain.Transaction, error) {

	q := psqlc.New(r.connPool)

	blockTx, err := q.GetBlockTxByTxHash(ctx, txHash)
	if err != nil {
		return nil, err
	}

	result := toTransaction(blockTx)

	return &result, nil
}

// Get max block number by from_address
func (r *PostgresRepository) GetMaxBlockNumberByFromAddress(ctx context.Context, from string) (int64, error) {

	q := psqlc.New(r.connPool)

	maxBlockNumber, err := q.GetMaxBlockNumberByFromAddress(ctx, from)
	if err != nil {
		return 0, err
	}

	if maxBlockNumber == nil {
		return 0, nil
	}

	return maxBlockNumber.(int64), nil
}

// List block transaction
func (r *PostgresRepository) ListBlockTx(ctx context.Context, size int) ([]blockchain.Transaction, error) {

	q := psqlc.New(r.connPool)

	blockTxs, err := q.ListBlockTx(ctx, int32(size))
	if err != nil {
		return nil, err
	}

	result := []blockchain.Transaction{}
	for _, blockTx := range blockTxs {
		tmp := toTransaction(blockTx)
		result = append(result, tmp)
	}

	return result, nil
}

// List block transaction by from_address
func (r *PostgresRepository) ListBlockTxByFromAddress(ctx context.Context, size int, from string) ([]blockchain.Transaction, error) {

	q := psqlc.New(r.connPool)

	blockTxs, err := q.ListBlockTxByFromAddress(ctx, psqlc.ListBlockTxByFromAddressParams{
		Limit:       int32(size),
		FromAddress: from,
	})
	if err != nil {
		return nil, err
	}

	result := []blockchain.Transaction{}
	for _, blockTx := range blockTxs {
		tmp := toTransaction(blockTx)
		result = append(result, tmp)
	}

	return result, nil
}

func toTransaction(params psqlc.BlockTx) blockchain.Transaction {
	return blockchain.Transaction{
		BlockNumber: uint64(params.BlockNumber),
		From:        eth_common.HexToAddress(params.FromAddress),
		To:          eth_common.HexToAddress(params.ToAddress),
		TxNonce:     uint64(params.TxNonce),
		TxHash:      params.TxHash,
		TxValue:     uint64(params.TxValue),
		TxGas:       uint64(params.TxGas),
		TxGasPrice:  uint64(params.TxGasPrice),
		TxData:      []byte(params.TxData),
		Timestamp:   params.TxTime,
	}
}
