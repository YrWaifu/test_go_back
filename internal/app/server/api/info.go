package api

import (
	"encoding/json"
	"fmt"
	"github.com/YrWaifu/test_go_back/internal/domain/auth/utils"
	infoUsecase "github.com/YrWaifu/test_go_back/internal/usecase/info"
	pkgErrors "github.com/YrWaifu/test_go_back/pkg/errors"
	"net/http"
)

type InfoApi struct {
	d InfoDependency
}

func NewInfoApi(d InfoDependency) *InfoApi {
	return &InfoApi{d: d}
}

type InfoRequest struct {
}

type CoinHistory struct {
	Received []Received `json:"received"`
	Sent     []Sent     `json:"sent"`
}
type Received struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type Sent struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type Inventory struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type InfoResponse struct {
	Coins       int         `json:"coins"`
	Inventory   []Inventory `json:"inventory"`
	CoinHistory CoinHistory `json:"coinHistory"`
}

func (a *InfoApi) Info(w http.ResponseWriter, r *http.Request) {
	senderId, ok := utils.ExtractAuth(r.Context())
	if !ok {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "auth required"}, http.StatusUnauthorized)
		return
	}

	uResp, err := a.d.InfoUsecase.Info(r.Context(), infoUsecase.InfoRequest{UserId: senderId})
	if err != nil {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: fmt.Errorf("info error: %w", err).Error()}, http.StatusInternalServerError)
		return
	}

	resp := InfoResponse{
		Coins:       uResp.User.Balance,
		Inventory:   make([]Inventory, len(uResp.Purchases)),
		CoinHistory: CoinHistory{Received: make([]Received, len(uResp.Received)), Sent: make([]Sent, len(uResp.Sent))},
	}

	for i, v := range uResp.Purchases {
		resp.Inventory[i] = Inventory{
			Type:     v.MerchName,
			Quantity: v.Quantity,
		}
	}

	for i, v := range uResp.Received {
		resp.CoinHistory.Received[i] = Received{
			FromUser: v.SenderName,
			Amount:   v.Amount,
		}
	}

	for i, v := range uResp.Sent {
		resp.CoinHistory.Sent[i] = Sent{
			ToUser: v.ReceiverName,
			Amount: v.Amount,
		}
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "json encode"}, http.StatusInternalServerError)
		return
	}
}
