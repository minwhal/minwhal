package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/minwhal/minwhal/internal/database"
)

func (a *AuthService) createRefreshToken(dbUser database.User) (RefreshToken, error) {

	return RefreshToken{}, nil
}

func (a *AuthService) createAccessToken(dbUser database.User) (AccessToken, error) {
	expiresAt := time.Now().Add(a.jwtAccessTokenDuration)
	claims := jwt.RegisteredClaims{
		Subject:   dbUser.ID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return AccessToken{}, errors.New("something wrong with secret")
	}
	return AccessToken{TokenString: secret, ExpiresAt: expiresAt}, nil
}
