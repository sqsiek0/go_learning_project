package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sqsiek0/go_learning_project/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No port")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No database")
	}
	connection, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("No database connection ", err)
	}

	apiCnf := apiConfig{
		DB: database.New(connection),
	}

	router := chi.NewRouter()
	routerHealth := chi.NewRouter()
	routerUsers := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300}))
	routerHealth.Get("/health", handleHealth)
	routerHealth.Get("/error", errorHandler)

	routerUsers.Post("/user", apiCnf.handleCreateUser)
	routerUsers.Get("/users", apiCnf.handleGetUsers)

	router.Mount("/check", routerHealth)
	router.Mount("/api", routerUsers)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Println("Serwer wystartował na porcie:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", port)
}
