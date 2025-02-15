package api

import (
	"encoding/json"
	"errors"
	pkgErrors "github.com/YrWaifu/test_go_back/pkg/errors"
	"log"
	"net/http"

	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	authUsecase "github.com/YrWaifu/test_go_back/internal/usecase/auth"
)

type AuthAPI struct {
	d AuthDependency
}

func NewAuthAPI(d AuthDependency) *AuthAPI {
	return &AuthAPI{d: d}
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func (a *AuthAPI) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "invalid request body"}, http.StatusBadRequest)
		return
	} else if req.Username == "" {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "username is required"}, http.StatusBadRequest)
		return
	} else if req.Password == "" {
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "password is required"}, http.StatusBadRequest)
		return
	}

	resp, err := a.d.AuthUsecase.SignIn(r.Context(), authUsecase.SignInRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		if errors.Is(err, userDomain.ErrUserNotFound) {
			pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "invalid credentials"}, http.StatusUnauthorized)
			return
		}
		log.Println(err.Error())
		pkgErrors.JSONError(w, pkgErrors.ErrorResponse{Errors: "internal server error"}, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SignInResponse{Token: resp.AccessToken})
}
