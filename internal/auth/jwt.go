package auth

import (
	"github.com/minwhal/minwhal/internal/database"
)

func (a *AuthService) createRefreshToken(dbUser database.User) (RefreshToken, error) {

	return RefreshToken{}, nil
}

func (a *AuthService) createAccessToken(dbUser database.User) (AccessToken, error) {
	return AccessToken{}, nil
}
