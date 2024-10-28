package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/router"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase"
)

func initAppConfig() *usecase.ApplicationParams {
	env := mustGetenv("NODE_ENV")

	// Database
	dbHost := mustGetenv("DB_HOST")
	dbPort := mustGetenv("DB_PORT")
	dbName := mustGetenv("DB_NAME")
	dbUser := mustGetenv("DB_USER")
	dbPwd := mustGetenv("DB_PASS")
	dbSchemaName := "public"

	return &usecase.ApplicationParams{
		Environment:  env,
		DBHost:       dbHost,
		DBPort:       dbPort,
		DBName:       dbName,
		DBUser:       dbUser,
		DBPassword:   dbPwd,
		DBSchemaName: dbSchemaName,
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Fatal Error in connect_connector.go: %s environment variable not set.\n", k)
	}
	return v
}

func main() {
	// Init configuration
	conf := initAppConfig()

	// Create root context
	rootCtx := context.Background()

	// Dependency injection
	app, err := usecase.NewApplication(rootCtx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	// Run server
	runHttpServer(rootCtx, app)
}

// Setup Gin route and run Gin server
//
//	@title			Ethereum Private Chain RPC Server
//	@version		1.0
//	@description	Ethereum Private Chain RPC Server
//	@host			localhost:8080
//	@basePath		/api/v1
func runHttpServer(rootCtx context.Context, app *usecase.Application) {

	// Set to release mode to disable Gin logger
	gin.SetMode(gin.ReleaseMode)
	if common.Local == app.Params.Environment {
		gin.SetMode(gin.DebugMode)
	}

	// Create gin router
	r := router.SetupGinRoute(rootCtx, app)

	// Listen and serve
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
