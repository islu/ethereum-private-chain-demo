package chain

import (
	"context"
	"errors"
	"math/big"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/blockchain"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
	"github.com/jackc/pgx/v5"
)

// Get the latest block number
func (s *ChainService) GetLatestBlockHeight(ctx context.Context) (uint64, common.Error) {

	blockNumber, err := s.nodeClient.GetLatestBlockNumber()
	if err != nil {
		err = errors.Join(errors.New("[ChainService][GetLatestBlockNumber] Get latest block number failed"), err)
		return 0, common.NewError(common.ErrorCodeInternalProcess, err)
	}

	return blockNumber, nil
}

// Get account balance
func (s *ChainService) GetBalance(ctx context.Context, accountAddress string) (*big.Int, common.Error) {

	balance, err := s.nodeClient.GetBalance(eth_common.HexToAddress(accountAddress))
	if err != nil {
		err = errors.Join(errors.New("[ChainService][GetBalance] Get balance failed"), err)
		return nil, common.NewError(common.ErrorCodeInternalProcess, err)
	}

	return balance, nil
}

// Get 0.02 coin from faucet
func (s *ChainService) GetCoinFromFaucet(ctx context.Context, accountAddress string) common.Error {

	amountToSend := big.NewInt(20000000000000000) // 0.02

	err := s.nodeClient.SendTx_Simulate(amountToSend, eth_common.HexToAddress(accountAddress))
	if err != nil {
		err = errors.Join(errors.New("[ChainService][GetCoinFromFaucet] Send tx failed"), err)
		return common.NewError(common.ErrorCodeInternalProcess, err)
	}
	return nil
}

var latestUpdatedBlockNumber = map[string]uint64{}

// 更新指定地址的交易資料進資料庫
func (s *ChainService) SyncTransactionForTargetAddress(ctx context.Context, accountAddress string) common.Error {

	var curBlockNumber uint64
	if number, exist := latestUpdatedBlockNumber[accountAddress]; exist {

		curBlockNumber = number

	} else {

		maxBlockNumber, err := s.txRepo.GetMaxBlockNumberByFromAddress(ctx, accountAddress)

		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			err = errors.Join(errors.New("[ChainService][SyncTransactionForTargetAddress] Get max block number failed"), err)
			return common.NewError(common.ErrorCodeInternalProcess, err)
		}

		curBlockNumber = uint64(maxBlockNumber)
	}

	// fmt.Println("curBlockNumber: ", curBlockNumber)

	latestBlockNumber, err := s.nodeClient.GetLatestBlockNumber()
	if err != nil {
		err = errors.Join(errors.New("[ChainService][SyncTransactionForTargetAddress] Get latest block number failed"), err)
		return common.NewError(common.ErrorCodeInternalProcess, err)
	}

	for i := curBlockNumber + 1; i <= latestBlockNumber; i++ {

		transactions, err := s.nodeClient.GetBlockTransactionsByNumber(int64(i))
		if err != nil {
			err = errors.Join(errors.New("[ChainService][SyncTransactionForTargetAddress] Get block transactions failed"), err)
			return common.NewError(common.ErrorCodeInternalProcess, err)
		}

		for _, transaction := range transactions {
			// 篩選出指定地址的交易
			if transaction.From == eth_common.HexToAddress(accountAddress) || transaction.To == eth_common.HexToAddress(accountAddress) {
				_, err := s.txRepo.CreateBlockTx(ctx, transaction)
				if err != nil {
					err = errors.Join(errors.New("[ChainService][SyncTransactionForTargetAddress] Create block transaction failed"), err)
					return common.NewError(common.ErrorCodeInternalProcess, err)
				}

			}
		}

		latestUpdatedBlockNumber[accountAddress] = i
	}

	// fmt.Println("latestUpdatedBlockNumber: ", latestUpdatedBlockNumber)

	return nil
}

// 查詢交易資料
func (s *ChainService) GetTransactions(ctx context.Context, size int, accountAddress string) ([]blockchain.Transaction, common.Error) {

	if accountAddress == "" {

		transactions, err := s.txRepo.ListBlockTx(ctx, size)
		if err != nil {
			err = errors.Join(errors.New("[ChainService][GetTransactions] Get transactions failed"), err)
			return nil, common.NewError(common.ErrorCodeInternalProcess, err)
		}

		return transactions, nil
	}

	transactions, err := s.txRepo.ListBlockTxByFromAddress(ctx, size, accountAddress)
	if err != nil {
		err = errors.Join(errors.New("[ChainService][GetTransactions] Get transactions by from address failed"), err)
		return nil, common.NewError(common.ErrorCodeInternalProcess, err)
	}

	return transactions, nil
}
