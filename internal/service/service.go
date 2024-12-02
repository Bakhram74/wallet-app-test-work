package service

import (
	"context"

	"github.com/Bakhram74/wallet-app-test-work/internal/entity"
	"github.com/Bakhram74/wallet-app-test-work/internal/repository"
	"github.com/google/uuid"
)

type Wallets interface {
	GetWalletBalance(ctx context.Context, id uuid.UUID) (int64, error)
	OperationWithWallet(ctx context.Context, params entity.WalletReq) (repository.Wallet, error)
}

type Service struct {
	Wallets
}

func NewService(repo *repository.Repository) *Service {

	return &Service{
		Wallets: NewWalletService(repo),
	}

}
