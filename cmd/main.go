package cmd

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/orca-infrastructures/orca/internal/auth"
	"github.com/orca-infrastructures/orca/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portEnv := os.Getenv("PORT")
	domainEnv := os.Getenv("DOMAIN")

	conn, err := sql.Open("sqlite", "app.db")
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer conn.Close()
	queries := database.New(conn)

	AuthHandler := auth.AuthHandler{&auth.AuthService{queries}}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", AuthHandler.HandleLogin)

	log.Fatal(http.ListenAndServe(domainEnv+portEnv, mux))
}
