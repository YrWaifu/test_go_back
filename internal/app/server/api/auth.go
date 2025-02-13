package api

import (
	"encoding/json"
	"errors"
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
		JSONError(w, ErrorResponse{Errors: "invalid request body"}, http.StatusBadRequest)
		return
	} else if req.Username == "" {
		JSONError(w, ErrorResponse{Errors: "username is required"}, http.StatusBadRequest)
		return
	} else if req.Password == "" {
		JSONError(w, ErrorResponse{Errors: "password is required"}, http.StatusBadRequest)
		return
	}

	resp, err := a.d.AuthUsecase.SignIn(r.Context(), authUsecase.SignInRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		if errors.Is(err, userDomain.ErrUserNotFound) {
			JSONError(w, ErrorResponse{Errors: "invalid credentials"}, http.StatusUnauthorized)
			return
		}
		log.Println(err.Error())
		JSONError(w, ErrorResponse{Errors: "internal server error"}, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SignInResponse{Token: resp.AccessToken})
}
