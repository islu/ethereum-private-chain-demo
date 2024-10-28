package usecase

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/islu/ethereum_private_chain/rpc_server/internal/adapter/blockchain"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/adapter/repository/postgres"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/usecase/service/chain"
)

type Application struct {
	Params       ApplicationParams
	ChainService *chain.ChainService
}

type ApplicationParams struct {
	// Env
	Environment string

	// Database
	DBHost       string
	DBPort       string
	DBName       string
	DBUser       string
	DBPassword   string
	DBSchemaName string
}

func NewApplication(ctx context.Context, param *ApplicationParams) (*Application, error) {

	// Initialize database
	pgRepo, err := initDatabase(ctx, *param)
	if err != nil {
		return nil, err
	}

	// Initialize ethereum client
	ethereumClient := &blockchain.EthereumPrivateNodeClient{
		Env: param.Environment,
	}

	// New application
	app := &Application{
		Params: *param,
		ChainService: chain.NewChainService(ctx, chain.ChainServiceParam{
			TxRepo:     pgRepo,
			NodeClient: ethereumClient,
		}),
	}
	return app, nil
}

/*
	Database
*/

func initDatabase(ctx context.Context, cfg ApplicationParams) (*postgres.PostgresRepository, error) {

	conn, err := connect(cfg)
	if err != nil {
		return nil, err
	}
	pgRepo := postgres.NewPostgresRepository(ctx, conn)
	return pgRepo, nil
}

// Connect postgres
func connect(cfg ApplicationParams) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s search_path=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSchemaName)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}
