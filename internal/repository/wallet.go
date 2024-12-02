package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidBalance = errors.New("insufficient balance")
	ErrWalletNotFound = errors.New("wallet with this id doesnt exist")
)

const (
	DEPOSIT  = "deposit"
	WITHDRAW = "withdraw"
)

type WalletRepo struct {
	db DBTX
}

func NewWalletRepo(db DBTX) *WalletRepo {
	return &WalletRepo{
		db: db,
	}
}

func (w *WalletRepo) CreateWallet(ctx context.Context, walletID uuid.UUID, amount int64) (Wallet, error) {

	query := "INSERT INTO wallets (wallet_id,balance) values ($1, $2) RETURNING wallet_id, balance,created_at"

	row := w.db.QueryRow(ctx, query, walletID, amount)
	var i Wallet
	err := row.Scan(
		&i.WalletID,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

func (w *WalletRepo) UpdateWalletBalance(ctx context.Context, walletID uuid.UUID, amount int64) (Wallet, error) {
	query := "UPDATE wallets SET balance = $2 WHERE wallet_id = $1 RETURNING wallet_id, balance,created_at"

	row := w.db.QueryRow(ctx, query,
		walletID,
		amount,
	)
	var i Wallet
	err := row.Scan(
		&i.WalletID,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

func (w *WalletRepo) GetBalance(ctx context.Context, walletID uuid.UUID) (int64, error) {
	query := "SELECT balance FROM wallets WHERE wallet_id = $1"

	var balance int64
	err := w.db.QueryRow(ctx, query, walletID).Scan(&balance)
	return balance, err
}
