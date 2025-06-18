package main

import (
	"backend/internal/application/services"
	"backend/internal/infra/db/postgres"
	"backend/internal/interface/api/rest"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	userRepo := postgres.NewPGUserRepo(conn)
	sessionRepo := postgres.NewPGSessionRepo(conn)

	authService := services.NewAuthService(userRepo, sessionRepo, conn)

	r := chi.NewRouter()
	rest.NewAuthController(authService).RegisterRoute(r)

	http.ListenAndServe(":3210", r)
}
