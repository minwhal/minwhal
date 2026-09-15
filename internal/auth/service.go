package auth

import (
	"context"
	"errors"

	"github.com/minwhal/minwhal/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	q *database.Queries
}

func NewAuthService(q *database.Queries) *AuthService {
	return &AuthService{q: q}
}

func (as *AuthService) Login(ctx context.Context, loginRequest *LoginRequest) (RefreshToken, error) {
	dbUser, err := as.q.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return RefreshToken{}, database.WrapNotFound(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(loginRequest.Password))
	if err != nil {
		return RefreshToken{}, errors.New("wrong password")
	}
	refreshToken := as.createRefreshToken(dbUser)
	return refreshToken, nil
}
