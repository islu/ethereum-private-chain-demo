package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/blockchain"
)

type EthereumPrivateNodeClient struct {
	Env          string
	RpcURL       string
	KeystorePath string
	KeystorePass string
}

// Get the latest block number
func (c *EthereumPrivateNodeClient) GetLatestBlockNumber() (uint64, error) {

	client, err := c.connect()
	if err != nil {
		// log.Printf("Failed to connect to the Ethereum client: %v\n", err)
		return 0, err
	}
	defer client.Close()

	// Get the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("Failed to get block number: %v\n", err)
		return 0, err
	}
	fmt.Printf("Latest block number: %d\n", blockNumber)

	return blockNumber, nil
}

// 取得指定區塊的交易
func (c *EthereumPrivateNodeClient) GetBlockTransactionsByNumber(blockNumber int64) ([]blockchain.Transaction, error) {

	client, err := c.connect()
	if err != nil {
		// log.Printf("Failed to connect to the Ethereum client: %v\n", err)
		return nil, err
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Println("Failed to get block: ", err)
		return nil, err
	}

	transactions := []blockchain.Transaction{}

	for _, tx := range block.Transactions() {

		from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
		if err != nil {
			log.Println("Failed to get tx sender: ", err)
		}

		// Build tx
		tmpTx := blockchain.Transaction{
			BlockNumber: block.NumberU64(),
			From:        from.Hex(),
			To:          tx.To().Hex(),
			TxNonce:     tx.Nonce(),
			TxHash:      tx.Hash().Hex(),
			TxValue:     tx.Value().Uint64(),
			TxGas:       tx.Gas(),
			TxGasPrice:  tx.GasPrice().Uint64(),
			TxData:      tx.Data(),
			Timestamp:   time.Unix(int64(block.Time()), 0),
		}

		transactions = append(transactions, tmpTx)
	}

	return transactions, nil
}

// Get account balance
func (c *EthereumPrivateNodeClient) GetBalance(account common.Address) (*big.Int, error) {

	client, err := c.connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Get the latest block number
	blockNumber, err := c.GetLatestBlockNumber()
	if err != nil {
		return nil, err
	}

	// Get account balance
	balance, err := client.BalanceAt(context.Background(), account, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Println("Get balance failed: ", err)
		return nil, err
	}
	// fmt.Println("Account balance: ", balance)

	return balance, nil
}

// Send transaction (Simulate)
func (c *EthereumPrivateNodeClient) SendTx_Simulate(amountToSend *big.Int, to common.Address) error {

	client, err := c.connect()
	if err != nil {
		return err
	}
	defer client.Close()

	// Read keystore
	keyjson, err := os.ReadFile(c.KeystorePath)
	if err != nil {
		log.Println("Read keystore failed: ", err)
		return err
	}

	// Unlock account
	key, err := keystore.DecryptKey(keyjson, c.KeystorePass)
	if err != nil {
		log.Println("Unlock keystore failed: ", err)
		return err
	}

	// Get unlock account
	address := key.Address
	privateKey := key.PrivateKey
	// amountToSend := big.NewInt(1000000000000000) // 0.001 eth in wei

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got nonce: %d\n", nonce)

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got gas price: %d\n", gasPrice)

	// Estimate gas
	estimateGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  address,
		To:    nil,
		Value: amountToSend,
		Data:  nil,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Estimated gas: %d\n", estimateGas)

	// Create transaction
	tx := types.NewTransaction(
		nonce,
		to,
		amountToSend,
		estimateGas,
		gasPrice,
		[]byte{},
	)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("Failed to get chain ID: ", err)
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println("Failed to sign transaction: ", err)
		return err
	}

	// Broadcast transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("Failed to send transaction:", err)
		return err
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	// Wait until transaction is confirmed
	var receipt *types.Receipt
	for {
		receipt, err = client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			fmt.Println("tx is not confirmed yet")
			time.Sleep(5 * time.Second)
		}
		if receipt != nil {
			break
		}
	}
	// Status = 1 if transaction succeeded
	log.Printf("tx is confirmed: %v. Block number: %v\n", receipt.Status, receipt.BlockNumber)

	return nil
}

// Connect to an Ethereum node
func (c *EthereumPrivateNodeClient) connect() (*ethclient.Client, error) {

	client, err := ethclient.Dial(c.RpcURL)

	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v\n", err)
		return nil, err
	}
	return client, nil
}
