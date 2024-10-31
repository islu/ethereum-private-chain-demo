package router

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/blockchain"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase"
)

// @description	Block response
type BlockResponse struct {
	Height uint64 `json:"height" example:"1030"`
} //	@name	BlockResponse

// @description	Balance response
type BalanceResponse struct {
	Wei     uint64   `json:"wei" example:"20000000000000000"`
	Balance float64 `json:"balance" example:"0.02"`
} //	@name	BalanceResponse

// 查詢最新區塊高度
//
//	@summary		查詢最新區塊高度
//	@description	查詢最新區塊高度
//	@description
//	@description	Error code list
//	@description	- 500: INTERNAL_PROCESS
//	@tags			chain
//	@accept			json
//	@produce		json
//	@router			/chain/blocks/height   [get]
//	@success		200	{object}	router.BlockResponse
//	@failure		400	{object}	router.ErrorMessage
//	@failure		500	{object}	router.ErrorMessage
func GetLatestBlockHeight(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := c.Request.Context()

		// Get latest block height
		height, commErr := app.ChainService.GetLatestBlockHeight(ctx)
		if commErr != nil {
			respondWithError(c, commErr)
			return
		}

		// Build response
		response := BlockResponse{
			Height: height,
		}

		respondWithJSON(c, http.StatusOK, response)
	}
}

// 查詢指定帳戶餘額
//
//	@summary		查詢指定帳戶餘額
//	@description	查詢指定帳戶餘額
//	@description
//	@description	Error code list
//	@description	- 400: PARAMETER_INVALID
//	@description	- 500: INTERNAL_PROCESS
//	@tags			chain
//	@accept			json
//	@produce		json
//	@router			/chain/balance/{address}   [get]
//	@param			address	path		string	true	"帳戶地址" example(0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045) default(0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045)
//	@success		200		{object}	router.BalanceResponse
//	@failure		400		{object}	router.ErrorMessage
//	@failure		500		{object}	router.ErrorMessage
func GetBalance(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := c.Request.Context()

		// Get account address
		address := c.Param(KeyAccountAddress)

		// TODO: check valid account address

		// Get balance
		balance, commErr := app.ChainService.GetBalance(ctx, address)
		if commErr != nil {
			respondWithError(c, commErr)
			return
		}

		// Build response
		coin, _ := common.FromWei(balance).Float64()

		response := BalanceResponse{
			Balance: coin,
			Wei:     balance.Uint64(),
		}

		respondWithJSON(c, http.StatusOK, response)
	}
}

// 取得 0.02 測試幣 (模擬發送交易)
//
//	@summary		取得 0.02 測試幣 (模擬發送交易)
//	@description	取得 0.02 測試幣 (模擬發送交易)
//	@description
//	@description	Error code list
//	@description	- 400: PARAMETER_INVALID
//	@description	- 500: INTERNAL_PROCESS
//	@tags			chain
//	@accept			json
//	@produce		json
//	@router			/chain/faucet/{address}   [post]
//	@param			address	path		string	true	"接收帳戶地址" example(0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045) default(0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045)
//	@success		200		{object}	router.SuccessMessage
//	@failure		400		{object}	router.ErrorMessage
//	@failure		500		{object}	router.ErrorMessage
func GetCoinFromFaucet(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := c.Request.Context()

		// Get account address
		address := c.Param(KeyAccountAddress)

		// TODO: check valid account address

		// Get coin from faucet
		commErr := app.ChainService.GetCoinFromFaucet(ctx, address)
		if commErr != nil {
			respondWithError(c, commErr)
			return
		}

		respondWithSuccess(c)
	}
}

// @description	Transaction request body
type TxBody struct {
	Size           int    `form:"size" binding:"required,min=10" example:"1030"`                // 回傳筆數
	AccountAddress string `form:"address" example:"0x8de0c53fc169ba09f111aa4170697e8cf42ccbbe"` // 要查詢的帳戶
} //	@name	TxBody

// @description	Transaction response
type TxResponse struct {
	Height uint64 `json:"height" example:"1"`
	From   string `json:"from" example:"0x8de0c53fc169ba09f111aa4170697e8cf42ccbbe"`
	To     string `json:"to" example:"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"`
	Wei    uint64 `json:"wei" example:"1000000000000000000"`
} // @name	TxResponse

var syncTxLock sync.Mutex

// 更新指定地址的交易資料進資料庫，再次觸發會等待上次同步完畢，避免重複觸發
//
//	@summary		更新指定地址的交易資料進資料庫，再次觸發會等待上次同步完畢，避免重複觸發
//	@description	更新指定地址的交易資料進資料庫，再次觸發會等待上次同步完畢，避免重複觸發
//	@description
//	@description	Error code list
//	@description	- 400: PARAMETER_INVALID
//	@description	- 500: INTERNAL_PROCESS
//	@tags			chain
//	@accept			json
//	@produce		json
//	@router			/chain/tx/{address}/sync   [post]
//	@param			address	path		string	false	"Account address"	example(0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe)	default(0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe)
//	@success		200		{object}	router.SuccessMessage
//	@failure		400		{object}	router.ErrorMessage
//	@failure		500		{object}	router.ErrorMessage
func SyncTransactionForTargetAddress(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		// ctx := c.Request.Context()

		// Get account address
		address := c.Param(KeyAccountAddress)

		if !syncTxLock.TryLock() {
			respondWithError(c, common.NewError(common.ErrorCodeProcessInProgress, errors.New("[ChainService][SyncTransactionForTargetAddress] Try lock failed, wait for previous sync finish")))
			return
		}

		go func() {
			// fmt.Println("enter go")

			defer syncTxLock.Unlock()

			commErr := app.ChainService.SyncTransactionForTargetAddress(context.Background(), address)
			if commErr != nil {
				log.Println("Sync transaction failed: ", commErr)
			}

			// 模擬假設同步時間
			time.Sleep(2 * time.Second)
		}()

		respondWithSuccess(c)
	}
}

// 查詢交易資料
//
//	@summary		查詢交易資料
//	@description	查詢交易資料
//	@description
//	@description	Error code list
//	@description	- 400: PARAMETER_INVALID
//	@description	- 500: INTERNAL_PROCESS
//	@tags			chain
//	@accept			json
//	@produce		json
//	@router			/chain/tx   [get]
//	@param			size	query		int		false	"Page size"			minimum(10)											example(10)	default(10)
//	@param			address	query		string	false	"Account address"	example(0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe)	default(0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe)
//	@success		200		{object}	[]router.TxResponse
//	@failure		400		{object}	router.ErrorMessage
//	@failure		500		{object}	router.ErrorMessage
func GetTransactions(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := c.Request.Context()

		// Get request body
		var body TxBody
		err := c.BindQuery(&body)
		if err != nil {
			err = errors.Join(errors.New("[GetTransactions] Get request body failed"), err)
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
			return
		}

		// Get transactions
		transactions, commErr := app.ChainService.GetTransactions(ctx, body.Size, body.AccountAddress)
		if commErr != nil {
			respondWithError(c, commErr)
			return
		}

		// Build response
		response := []TxResponse{}

		for _, tx := range transactions {
			tmp := toTxResponse(tx)
			response = append(response, tmp)
		}

		respondWithJSON(c, http.StatusOK, response)
	}
}

func toTxResponse(params blockchain.Transaction) TxResponse {
	return TxResponse{
		Height: params.BlockNumber,
		From:   params.From.Hex(),
		To:     params.To.Hex(),
		Wei:    params.TxValue,
	}
}
