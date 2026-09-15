package auth

import (
	"net/http"

	"github.com/minwhal/minwhal/internal/httpx"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type LoginRequest struct {
	Email    string
	Password string
}

func (a *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := httpx.Decode[LoginRequest](r)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	token, err := a.authService.Login(r.Context(), loginRequest)
	if err != nil {
		httpx.ResponseWithError(w, http.StatusNotFound, err)
		return
	}
	httpx.WriteJson(w, http.StatusOK, token)
}
