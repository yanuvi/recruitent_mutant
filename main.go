package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yanuvi/recruitent_mutant/src/handlers"
	"github.com/yanuvi/recruitent_mutant/src/server"
)

type CreateMutantRequest struct {
	Dna []string `json:"dna"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}
	s.Start(BindRouters)
	lambda.Start(BindRouters)
}

func BindRouters(s server.Server, r *mux.Router) {
	r.HandleFunc("/mutant", handlers.MutantHandler(s)).Methods(http.MethodPost)
}
