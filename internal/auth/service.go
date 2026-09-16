package auth

import (
	"context"
	"errors"

	"github.com/minwhal/minwhal/internal/apperr"
	"github.com/minwhal/minwhal/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	q *database.Queries
}

func NewAuthService(q *database.Queries) *AuthService {
	return &AuthService{q: q}
}

func (as *AuthService) Login(ctx context.Context, loginRequest *LoginRequest) (LoginResponse, error) {
	dbUser, err := as.q.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(loginRequest.Password))
	if err != nil {
		return LoginResponse{}, apperr.ErrUnauthorized
	}
	refreshToken, err := as.createRefreshToken(dbUser)
	if err != nil {
		return LoginResponse{}, errors.New("jwt error")
	}
	accessToken, err := as.createAccessToken(dbUser)
	if err != nil {
		return LoginResponse{}, errors.New("jwt error")
	}
	return LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
