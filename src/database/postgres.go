package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/yanuvi/recruitent_mutant/src/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertMutant(ctx context.Context, mutant *models.Mutant) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO mutant (id, dna, reclutamiento) VALUES ($1, $2, $3)", mutant.Id, mutant.Dna, mutant.Reclutamiento)
	return err
}

func (repo *PostgresRepository) CloseConnectionBD() error {
	return repo.db.Close()
}
