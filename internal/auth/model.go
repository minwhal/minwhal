package auth

type AccessToken struct {
	TokenString string
	ExpiresAt   string
}

type RefreshToken struct {
	ExpirestAt  string
	TokenString string
}
