package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"github.com/yanuvi/recruitent_mutant/src/entities"
	"github.com/yanuvi/recruitent_mutant/src/models"
	"github.com/yanuvi/recruitent_mutant/src/repository"
	"github.com/yanuvi/recruitent_mutant/src/server"
	"github.com/yanuvi/recruitent_mutant/src/usecases"
)

type MutantRequest struct {
	Dna     string   `json:"dna"`
	DnaOtro []string `json:"dnaotro"`
}

type MutantResponse struct {
	Id  string `json:"id"`
	Dna string `json:"dna"`
}

func MutantHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = MutantRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //error del lado del cliente
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //error del servidor
			return
		}
		// incluir la validacion si se mutante o no para insertar
		var mutant = models.Mutant{
			Dna: request.Dna,
			Id:  id.String(),
		}
		var mutantUsecases *usecases.MutantImpl = usecases.NewMutantImpl()
		usecases.SetMutantUseCases(mutantUsecases)
		requestUsecases := &entities.MutantRequest{
			Dna: request.DnaOtro,
		}

		result := usecases.IsMutant(requestUsecases)
		if result {
			//fmt.Println("Es mutante")
			err = repository.InsertMutant(r.Context(), &mutant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			//fmt.Println("No es mutante")
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "", http.StatusForbidden)
		}
		/*
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(MutantResponse{
				Id:  mutant.Id,
				Dna: mutant.Dna,
			})*/
	}
}
