package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase"
)

// @description	Block response
type BlockResponse struct {
	Height uint64 `json:"height"`
} //	@name	BlockResponse

// @description	Balance response
type BalanceResponse struct {
	Wei     int64   `json:"wei"`
	Balance float64 `json:"balance"`
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
//	@param			address	path		string	true	"帳戶地址"
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
			Wei:     balance.Int64(),
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
//	@param			address	path		string	true	"接收帳戶地址"
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
