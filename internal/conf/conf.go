package conf

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Cfg struct {
	Port                    string
	Domain                  string
	JwtSecret               string
	JwtAccessTokenDuration  time.Duration
	JwtRefreshTokenDuration time.Duration
}

func LoadConfig() (*Cfg, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portEnv := os.Getenv("PORT")
	domainEnv := os.Getenv("DOMAIN")

	secretEnv := os.Getenv("JWT_SECRET")
	if secretEnv == "" {
		return &Cfg{}, errors.New("required fields of config cannot be empty")
	}

	jwtAccessTokenDurationEnv := os.Getenv("JWT_ACCESS_TOKEN_DURATION")
	jwtRefreshTokenDurationEnv := os.Getenv("JWT_REFRESH_TOKEN_DURATION")

	var jwtAccessTokenDuration time.Duration
	jwtAccessTokenDuration, err = time.ParseDuration(jwtAccessTokenDurationEnv)
	if err != nil {
		jwtAccessTokenDuration = time.Duration(time.Minute * 5)
	}

	var jwtRefreshTokenDuration time.Duration
	jwtRefreshTokenDuration, err = time.ParseDuration(jwtRefreshTokenDurationEnv)
	if err != nil {
		jwtRefreshTokenDuration = time.Duration(time.Minute * 5)
	}

	return &Cfg{
		Port:                    portEnv,
		Domain:                  domainEnv,
		JwtSecret:               secretEnv,
		JwtAccessTokenDuration:  jwtAccessTokenDuration,
		JwtRefreshTokenDuration: jwtRefreshTokenDuration,
	}, nil
}
