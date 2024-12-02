package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpdateWallet(ctx context.Context, walletID uuid.UUID, amount int64) (Wallet, error) {
	// ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	var wallet Wallet
	var balance int64

	err := r.execTx(ctx, func(q Queries) error {
		var err error
		err = r.connPool.QueryRow(ctx, "SELECT balance FROM wallets WHERE wallet_id = $1 FOR UPDATE", walletID).Scan(&balance)

		if err != nil {
			if err == pgx.ErrNoRows {
				if amount < 0 {
					return ErrWalletNotFound
				}
				wallet, err = q.CreateWallet(ctx, walletID, amount)
				if err != nil {
					return err
				}
			}
			return err
		}

		newBalance := balance + amount
		if newBalance < 0 {
			return ErrInvalidBalance
		}
		wallet, err = q.UpdateWalletBalance(ctx, walletID, newBalance)
		if err != nil {
			return err
		}

		return nil
	})
	return wallet, err
}
