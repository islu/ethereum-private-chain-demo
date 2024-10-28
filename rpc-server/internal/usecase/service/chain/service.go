package chain

import "context"

type ChainService struct {
	txRepo     TransactionRepository
	nodeClient EthereumNodeClient
}

type ChainServiceParam struct {
	TxRepo     TransactionRepository
	NodeClient EthereumNodeClient
}

func NewChainService(_ context.Context, param ChainServiceParam) *ChainService {
	return &ChainService{
		txRepo:     param.TxRepo,
		nodeClient: param.NodeClient,
	}
}
