package repository

import (
	"context"

	"github.com/yanuvi/recruitent_mutant/src/models"
)

type Repository interface {
	InsertMutant(ctx context.Context, mutant *models.Mutant) error
	CloseConnectionBD() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertMutant(ctx context.Context, mutant *models.Mutant) error {
	return implementation.InsertMutant(ctx, mutant)
}

func CloseConnectionBD() error {
	return implementation.CloseConnectionBD()
}
