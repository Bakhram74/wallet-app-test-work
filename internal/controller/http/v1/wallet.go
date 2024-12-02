package v1

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Bakhram74/wallet-app-test-work/internal/entity"
	"github.com/Bakhram74/wallet-app-test-work/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Routes) walletOperation(ctx *gin.Context) {
	var reqBody entity.WalletReq

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	if reqBody.Amount <= 0 {
		err := errors.New("amount must be greater than 0")
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	reqBody.OperationType = strings.ToLower(reqBody.OperationType)

	if reqBody.OperationType != repository.DEPOSIT && reqBody.OperationType != repository.WITHDRAW {
		err := errors.New("operation type must be deposit or withdraw")
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	wallet, err := r.service.Wallets.OperationWithWallet(ctx, reqBody)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidBalance) {
			errorResponse(ctx, http.StatusBadRequest, err.Error())
			slog.Error(err.Error())
			return
		}
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, repository.ErrWalletNotFound) {
			errorResponse(ctx, http.StatusNotFound, err.Error())
			slog.Error(err.Error())
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, wallet)
}

func (r *Routes) getBalance(ctx *gin.Context) {

	walletID, err := uuid.Parse(ctx.Param("walletId"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	balance, err := r.service.Wallets.GetWalletBalance(ctx, walletID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errorResponse(ctx, http.StatusNotFound, err.Error())
			slog.Error(err.Error())
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]int64{"balance": balance})
}
