package repository

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	WalletID  uuid.UUID `json:"wallet_id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
