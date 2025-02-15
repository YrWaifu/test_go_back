package api

import (
	"errors"
	"github.com/YrWaifu/test_go_back/internal/domain/auth/utils"
	merchDomain "github.com/YrWaifu/test_go_back/internal/domain/merch"
	purchaseUsecase "github.com/YrWaifu/test_go_back/internal/usecase/purchase"
	pkgErrors "github.com/YrWaifu/test_go_back/pkg/errors"
	"github.com/go-chi/chi"
	"net/http"
)

type PurchaseAPI struct {
	d PurchaseDependency
}

func NewPurchaseAPI(d PurchaseDependency) *PurchaseAPI {
	return &PurchaseAPI{d: d}
}

func (a *PurchaseAPI) PurchaseMerch(w http.ResponseWriter, r *http.Request) {
	merchName := chi.URLParam(r, "merchName")
	if merchName == "" {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "merch name is required"}, http.StatusBadRequest)
		return
	}

	id, ok := utils.ExtractAuth(r.Context())
	if !ok {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "auth required"}, http.StatusUnauthorized)
		return
	}

	_, err := a.d.PurchaseUsecase.BuyMerch(r.Context(), purchaseUsecase.BuyMerchRequest{
		UserID:    id,
		MerchName: merchName,
	})
	if err != nil {
		if errors.Is(err, merchDomain.ErrMerchNotFound) {
			pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "merch not found"}, http.StatusNotFound)
			return
		}

		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: err.Error()}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
