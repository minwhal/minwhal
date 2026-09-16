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
type LoginResponse struct {
	RefreshToken RefreshToken
	AccessToken  AccessToken
}

func (a *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	loginRequest, err := httpx.Decode[LoginRequest](r)
	if err != nil {
		return err
	}
	loginResponse, err := a.authService.Login(r.Context(), loginRequest)
	if err != nil {
		return err
	}

	httpx.WriteJson(w, http.StatusOK, loginResponse)
	return nil
}
