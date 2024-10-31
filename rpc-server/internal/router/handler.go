package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/islu/ethereum_private_chain/rpc_server/docs"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RegisterHandlers(ginRouter *gin.Engine, app *usecase.Application) *gin.Engine {

	/*
		Handlers
	*/

	// Add swagger-ui
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Set docs info for swagger in local
	if common.Local == app.Params.Environment {
		// docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Host = "localhost:8080"
	}

	// Add ping
	ginRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Add health-check
	ginRouter.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "is alive"})
	})

	// Mount all handlers under /api path
	api := ginRouter.Group("/api")
	v1 := api.Group("/v1")

	// Add chain namespace
	chainGroup := v1.Group("/chain")
	{
		// 查詢最新區塊高度
		chainGroup.GET("/blocks/height", GetLatestBlockHeight(app))
		// 查詢指定帳戶餘額
		chainGroup.GET("/balance/:"+KeyAccountAddress, GetBalance(app))
		// 取得 0.02 測試幣 (模擬發送交易)
		chainGroup.POST("/faucet/:"+KeyAccountAddress, GetCoinFromFaucet(app))

		// 更新指定地址的交易資料進資料庫
		chainGroup.POST("/tx/:"+KeyAccountAddress+"/sync", SyncTransactionForTargetAddress(app))
		// 查詢交易資料
		chainGroup.GET("/tx", GetTransactions(app))
	}

	return ginRouter
}
