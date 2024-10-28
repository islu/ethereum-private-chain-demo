package chain

import (
	"context"
	"errors"
	"math/big"

	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
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

	err := s.nodeClient.SendTx_Simulate(big.NewInt(20000000000000000), eth_common.HexToAddress(accountAddress))
	if err != nil {
		err = errors.Join(errors.New("[ChainService][GetCoinFromFaucet] Send tx failed"), err)
		return common.NewError(common.ErrorCodeInternalProcess, err)
	}
	return nil
}
