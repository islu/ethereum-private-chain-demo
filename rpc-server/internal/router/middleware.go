package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase"
)

// SetGeneralMiddlewares add general-purpose middlewares
func SetGeneralMiddlewares(ctx context.Context, ginRouter *gin.Engine, app *usecase.Application) {

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	ginRouter.Use(gin.Recovery())

	// Add CORS middleware
	ginRouter.Use(CORSMiddleware())

	ginRouter.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/ping", "/healthz", "/health"},
	}))
}
