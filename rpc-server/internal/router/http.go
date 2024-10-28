package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase"
)

// Setup gin route
func SetupGinRoute(rootCtx context.Context, app *usecase.Application) *gin.Engine {

	// Create gin router
	ginRouter := gin.New()

	// Set general middleware
	SetGeneralMiddlewares(rootCtx, ginRouter, app)

	// Register handlers
	r := RegisterHandlers(ginRouter, app)

	return r
}
