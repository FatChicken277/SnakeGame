package main

import (
	"SnakeGame/backend/handlers"
	"SnakeGame/backend/storage"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const (
	// DatabaseSource reference to the database path
	DatabaseSource = "postgresql://admin@localhost:26257/snake?sslmode=disable"
)

func main() {
	port := ":3000"
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	dbConn, err := storage.DBConection(DatabaseSource)
	if err != nil {
		handlers.LogError(err)
	}
	defer dbConn.Close(context.Background())

	router.Route("/v1/players", func(r chi.Router) {
		r.Post("/register", handlers.PlayerRegister(dbConn))
		r.Get("/leaderboard", handlers.PlayerLeaderboard(dbConn))
	})

	fmt.Println("Server running in localhost" + port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		handlers.LogError(err)
	}
}
