package repository

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createWallet(t *testing.T) Wallet {
	id := uuid.New()

	wallet, err := repo.UpdateWallet(context.Background(), id, 0)

	require.NoError(t, err)
	require.NotEmpty(t, wallet)
	require.Equal(t, id, wallet.WalletID)
	require.Equal(t, wallet.Balance, int64(0))
	require.NotZero(t, wallet.CreatedAt)

	return wallet
}

func TestUpdateWallet(t *testing.T) {
	var mu sync.Mutex
	w := createWallet(t)
	// var err error
	// var wallet Wallet
	var num int64
	n := 1000
	errs := make(chan error)
	amount := make(chan int64)

	for i := 0; i < n; i++ {
		go func() {
			mu.Lock()
			wallet, err := repo.UpdateWallet(context.Background(), w.WalletID, 1)
			mu.Unlock()
			errs <- err
			amount <- wallet.Balance
		}()
	}

	go func() {
		<-amount
		mu.Lock()
		wallet, _ := repo.UpdateWallet(context.Background(), w.WalletID, -1)
		mu.Unlock()
		num = wallet.Balance
	}()

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
	require.Equal(t, 0, int(num))
}
