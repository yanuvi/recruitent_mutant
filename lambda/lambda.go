package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yanuvi/recruitent_mutant/src/database"
	"github.com/yanuvi/recruitent_mutant/src/handlers"
	"github.com/yanuvi/recruitent_mutant/src/models"
)

func HandleRequest(ctx context.Context, request handlers.MutantRequest) (models.Mutant, error) {
	postgresConnector := database.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.CloseConnectionBD()
	if err != nil {
		return models.Mutant{}, err
	}
	db2.AutoMigrate(&models.Mutant{})

	var mutant = &models.Mutant{
		Dna: strings.Join(request.Dna, ","),
	}

	db2.Create(mutant)
	return *mutant, nil
}

func main() {
	lambda.Start(HandleRequest)
}
