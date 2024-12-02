package service

import (
	"context"
	"errors"

	"github.com/Bakhram74/wallet-app-test-work/internal/entity"
	"github.com/Bakhram74/wallet-app-test-work/internal/repository"
	"github.com/google/uuid"
)

type WalletService struct {
	repo *repository.Repository
}

func NewWalletService(repo *repository.Repository) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (w *WalletService) OperationWithWallet(ctx context.Context, params entity.WalletReq) (repository.Wallet, error) {
	
	switch params.OperationType {
	case repository.DEPOSIT:
		return w.repo.UpdateWallet(ctx, params.WalletID, params.Amount)

	case repository.WITHDRAW:
		return w.repo.UpdateWallet(ctx, params.WalletID, -params.Amount)
		
	default:
		return repository.Wallet{}, errors.New("unexpected error")
	}
}

func (w *WalletService) GetWalletBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	return w.repo.GetBalance(ctx, id)
}
