package auth

import (
	"net/http"

	"github.com/orca-infrastructures/orca/internal/httpx"
)

type AuthHandler struct {
	authService *AuthService
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
		httpx.WriteError(w)
	}
	httpx.WriteJson(w, http.StatusOK, token)
}
