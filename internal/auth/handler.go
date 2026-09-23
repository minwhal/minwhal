package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/minwhal/minwhal/internal/httpx"
)

type AuthHandler struct {
	authService *AuthService
}

func (a *AuthHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", a.HandleLogin)
	return r
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

func (a *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := httpx.Decode[LoginRequest](r)
	if err != nil {
		httpx.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}
	loginResponse, err := a.authService.Login(r.Context(), loginRequest)
	if err != nil {
		httpx.ResponseWithError(w, http.StatusUnauthorized, err)
		return
	}

	httpx.WriteJson(w, http.StatusOK, loginResponse)
}
