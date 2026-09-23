package auth

import (
	"context"
	"errors"
	"time"

	"github.com/minwhal/minwhal/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	q                       *database.Queries
	jwtSecret               string
	jwtAccessTokenDuration  time.Duration
	jwtRefreshTokenDuration time.Duration
}

func NewAuthService(q *database.Queries, jwtSecret string, accessTokenDuration time.Duration, refreshTokenDuration time.Duration) *AuthService {
	return &AuthService{q: q, jwtSecret: jwtSecret, jwtAccessTokenDuration: accessTokenDuration, jwtRefreshTokenDuration: refreshTokenDuration}
}

func (as *AuthService) Login(ctx context.Context, loginRequest *LoginRequest) (LoginResponse, error) {
	dbUser, err := as.q.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(loginRequest.Password))
	if err != nil {
		return LoginResponse{}, errors.New("unauthorized")
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
