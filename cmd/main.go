package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/minwhal/minwhal/internal/auth"
	"github.com/minwhal/minwhal/internal/conf"
	"github.com/minwhal/minwhal/internal/database"
)

func main() {
	cfg, err := conf.LoadConfig()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	conn, err := sql.Open("sqlite", "app.db")
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}
	defer conn.Close()
	queries := database.New(conn)

	authService := auth.NewAuthService(queries, cfg.JwtSecret, cfg.JwtAccessTokenDuration, cfg.JwtRefreshTokenDuration)
	authHandler := auth.NewAuthHandler(authService)

	router := chi.NewRouter()
	router.Mount("/auth", authHandler.Routes())

	log.Fatal(http.ListenAndServe(cfg.Domain+":"+cfg.Port, router))
}
