package auth

import (
	"context"

	"github.com/orca-infrastructures/orca/internal/database"
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
	refreshToken := as.createRefreshToken(dbUser)
	return refreshToken, nil
}
