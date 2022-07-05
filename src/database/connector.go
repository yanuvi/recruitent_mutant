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
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	_, err = sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	return
}
