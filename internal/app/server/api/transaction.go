package api

import (
	"encoding/json"
	"github.com/YrWaifu/test_go_back/internal/domain/auth/utils"
	transactionUsecase "github.com/YrWaifu/test_go_back/internal/usecase/transaction"
	pkgErrors "github.com/YrWaifu/test_go_back/pkg/errors"
	"net/http"
)

type TransactionAPI struct {
	d TransactionDependency
}

func NewTransactionAPI(d TransactionDependency) *TransactionAPI {
	return &TransactionAPI{d: d}
}

type SendCoinRequest struct {
	ReceiverName string `json:"toUser"`
	Amount       int    `json:"amount"`
}

func (a TransactionAPI) SendCoin(w http.ResponseWriter, r *http.Request) {
	var req SendCoinRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "invalid request body"}, http.StatusBadRequest)
		return
	} else if req.ReceiverName == "" {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "receiver name is required"}, http.StatusBadRequest)
		return
	} else if req.Amount <= 0 {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "amount must be > 0"}, http.StatusBadRequest)
		return
	}

	senderId, ok := utils.ExtractAuth(r.Context())
	if !ok {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "auth required"}, http.StatusUnauthorized)
		return
	}

	_, err = a.d.TransactionUsecase.Transfer(r.Context(), transactionUsecase.TransferRequest{
		SenderID:     senderId,
		ReceiverName: req.ReceiverName,
		Amount:       req.Amount,
	})
	if err != nil {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "transfer error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
