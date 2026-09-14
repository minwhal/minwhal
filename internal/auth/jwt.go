package auth

import (
	"github.com/orca-infrastructures/orca/internal/database"
)

func (a *AuthService) createRefreshToken(dbUser database.User) RefreshToken {

	return RefreshToken{}
}
