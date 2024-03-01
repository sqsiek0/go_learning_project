package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No port")
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

	routerUsers.Get("/users", getUsers)
	routerUsers.Get("/health", handleHealth)

	router.Mount("/check", routerHealth)
	router.Mount("/api", routerUsers)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Println("Serwer wystartował na porcie:", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", port)
}
