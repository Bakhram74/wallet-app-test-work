package entity

import "github.com/google/uuid"

type WalletReq struct {
	WalletID      uuid.UUID `json:"valletId" binding:"required"`
	OperationType string    `json:"operationType" binding:"required"`
	Amount        int64     `json:"amount" binding:"required"`
}
