package main

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

	_ "github.com/joho/godotenv/autoload"
)

// Env variables
var (
	RPCAddress   = os.Getenv("RPC_SERVER_ADDRESS")
	KeystorePath = os.Getenv("KEYSTORE_PATH")
	KeystorePass = os.Getenv("KEYSTORE_PASS")
)

func main() {

	client, err := ethclient.Dial(RPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	from := "0x8de0c53fc169ba09f111aa4170697e8cf42ccbbe"
	to := "0x6290a833deb0975a76bce27fe6fec6c1fb982aef"

	fmt.Println()
	fmt.Println("[Before]")
	getBalance(client, to)

	blockNumber := sendTx_Simulate(client, common.HexToAddress(to))

	fmt.Println()
	fmt.Println("[After]")
	getBalance(client, to)

	getTransactionCount(client, common.HexToAddress(from))

	// i := big.NewInt(991)
	i := blockNumber
	block, err := client.BlockByNumber(context.Background(), i)
	if err != nil {
		log.Fatalf("Failed to retrieve block: %v", err)
	}

	fmt.Println()
	fmt.Println("Block Nonce: ", block.Nonce())
	fmt.Println("Block Number: ", block.NumberU64())
	fmt.Println("Block Hash: ", block.Hash().Hex())
	fmt.Println("Block Time: ", time.Unix(int64(block.Time()), 0))
	fmt.Println("Block Difficulty: ", block.Difficulty().Uint64())
	fmt.Println("Tx Hash", block.TxHash().Hex())
	fmt.Println("Parent Hash: ", block.ParentHash().Hex())
	fmt.Println("Receipt Hash: ", block.ReceiptHash().Hex())
	fmt.Println()

	for _, tx := range block.Transactions() {

		fmt.Println("Tx Nonce: ", tx.Nonce())
		if from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx); err == nil {
			fmt.Println("Tx From: ", from.Hex())
		}
		// if from, err := types.Sender(types.NewLondonSigner(tx.ChainId()), tx); err == nil {
		// 	fmt.Println("Tx From: ", from.Hex())
		// }
		fmt.Println("Tx Hash: ", tx.Hash().Hex())
		fmt.Println("Tx To: ", tx.To().Hex())
		fmt.Println("Tx Data: ", string(tx.Data()))
		fmt.Println("Tx Value: ", tx.Value())
		fmt.Println("Tx Gas: ", tx.Gas())
		fmt.Println("Tx Gas Price: ", tx.GasPrice())
		fmt.Println("Tx Type: ", tx.Type())
		fmt.Println("Tx Time: ", time.Unix(int64(block.Time()), 0))
	}

}

// Get transaction count
func getTransactionCount(client *ethclient.Client, account common.Address) uint64 {

	nonce, err := client.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalln("Get nonce failed", err)
		return 0
	}
	fmt.Println("Account nonce: ", nonce)

	return nonce
}

// Get the latest block number
func getLatestBlockNumber(client *ethclient.Client) uint64 {
	// Get the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get block number: %v", err)
	}
	fmt.Printf("Latest block number: %d\n", blockNumber)

	return blockNumber
}

// Get account balance
func getBalance(client *ethclient.Client, accountAddress string) {

	// Get the latest block number
	blockNumber := getLatestBlockNumber(client)

	// Get account balance
	account := common.HexToAddress(accountAddress)
	balance, err := client.BalanceAt(context.Background(), account, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	fmt.Println("Account balance: ", balance)
}

// Send transaction (Simulate)
func sendTx_Simulate(client *ethclient.Client, to common.Address) *big.Int {

	// Read keystore
	keyjson, err := os.ReadFile(KeystorePath)
	if err != nil {
		log.Fatalf("Read keystore failed: %v", err)
	}

	// Unlock account
	key, err := keystore.DecryptKey(keyjson, KeystorePass)
	if err != nil {
		log.Fatalf("Unlock keystore failed: %v", err)
	}

	// Get unlock address
	address := key.Address
	fmt.Println("Unlock address: ", address)
	privateKey := key.PrivateKey

	amountToSend := big.NewInt(1000000000000000) // 0.001 eth in wei

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
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
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
	fmt.Printf("tx is confirmed: %v. Block number: %v\n", receipt.Status, receipt.BlockNumber)

	return receipt.BlockNumber
}
