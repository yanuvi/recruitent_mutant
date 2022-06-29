package database

import (
	"context"
	"database/sql"
	"log"

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
	_, err := repo.db.ExecContext(ctx, "INSERT INTO mutant (id, dna) VALUES (&1, &2)", mutant.Id, mutant.Dna)
	return err
}

func (repo *PostgresRepository) GetMutantById(ctx context.Context, id int64) (*models.Mutant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT dna FROM mutant WHERE id = &1", id)
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var mutant = models.Mutant{}
	for rows.Next() {
		if err = rows.Scan(&mutant.Id, &mutant.Dna); err == nil {
			return &mutant, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &mutant, nil
}
func (repo *PostgresRepository) CloseConnectionBD() error {
	return repo.db.Close()
}
