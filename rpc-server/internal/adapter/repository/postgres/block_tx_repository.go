package postgres

import (
	"context"
	"time"

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
		FromAddress: params.From,
		ToAddress:   params.To,
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

// List block transaction
func (r *PostgresRepository) ListBlockTx(ctx context.Context) ([]blockchain.Transaction, error) {

	q := psqlc.New(r.connPool)

	blockTxs, err := q.ListBlockTx(ctx, 1000) // Default page size: 1000
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
		From:        params.FromAddress,
		To:          params.ToAddress,
		TxNonce:     uint64(params.TxNonce),
		TxHash:      params.TxHash,
		TxValue:     uint64(params.TxValue),
		TxGas:       uint64(params.TxGas),
		TxGasPrice:  uint64(params.TxGasPrice),
		TxData:      []byte(params.TxData),
		Timestamp:   params.TxTime,
	}
}
