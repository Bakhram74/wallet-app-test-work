package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type Queries interface {
	CreateWallet(ctx context.Context, walletID uuid.UUID, amount int64) (Wallet, error)
	UpdateWalletBalance(ctx context.Context, walletID uuid.UUID, amount int64) (Wallet, error)
	GetBalance(ctx context.Context, walletID uuid.UUID) (int64, error)
}

type Repository struct {
	connPool *pgxpool.Pool
	Queries
}

func New(connPool *pgxpool.Pool) *Repository {
	return &Repository{
		connPool: connPool,
		Queries:  NewWalletRepo(connPool),
	}
}
