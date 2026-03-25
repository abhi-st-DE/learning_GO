package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/abhi-st-DE/learning_GO/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// this takes the .env file and loads it into the environment variables.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	fmt.Printf("Server will start on port: %s\n", portString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:  		[]string{"https://*", "http://*"},
		AllowedMethods: 	[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: 	[]string{"*"},
		ExposedHeaders: 	[]string{"Link"},
		AllowCredentials: 	false,
		MaxAge:           	300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: 	router,
		Addr: 		":" + portString,
	}

	log.Printf("Starting server on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}