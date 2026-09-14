package auth

import "time"

type AccessToken struct {
	TokenString string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type RefreshToken struct {
	ExpiresAt   time.Time `json:"expires_at"`
	TokenString string    `json:"token"`
}
