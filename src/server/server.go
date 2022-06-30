package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yanuvi/recruitent_mutant/src/database"
	"github.com/yanuvi/recruitent_mutant/src/repository"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}
