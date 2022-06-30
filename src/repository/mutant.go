package repository

import (
	"context"

	"github.com/yanuvi/recruitent_mutant/src/models"
)

type MutantRepository interface {
	InsertMutant(ctx context.Context, mutant *models.Mutant) error
	GetMutantById(ctx context.Context, id string) (*models.Mutant, error)
	CloseConnectionBD() error
}

var implementation MutantRepository

func SetRepository(repository MutantRepository) {
	implementation = repository
}

func InsertMutant(ctx context.Context, mutant *models.Mutant) error {
	return implementation.InsertMutant(ctx, mutant)
}

func GetMutantById(ctx context.Context, id string) (*models.Mutant, error) {
	return implementation.GetMutantById(ctx, id)
}

func CloseConnectionBD() error {
	return implementation.CloseConnectionBD()
}
