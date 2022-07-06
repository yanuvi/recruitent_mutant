package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConnector struct {
}

func (p *PostgresConnector) GetConnection() (db *PostgresRepository, err error) {
	err = godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	dbName := os.Getenv("DATABASE_URL")

	dbURI := fmt.Sprintf("DATABASE_URL=%s", dbName)
	_, err = sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	return
}
