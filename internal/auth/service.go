package auth

import (
	"context"

	"github.com/orca-infrastructures/orca/internal/database"
)

type AuthService struct {
	q *database.Queries
}

func (as *AuthService) Login(ctx context.Context, loginRequest LoginRequest) (RefreshToken, error) {
	dbUser, err := as.q.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return RefreshToken{}, database.WrapNotFound(err)
	}
	refreshToken := createRefreshToken(dbUser)
	return refreshToken, nil
}
