package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/segmentio/ksuid"
	"github.com/yanuvi/recruitent_mutant/src/entities"
	"github.com/yanuvi/recruitent_mutant/src/models"
	"github.com/yanuvi/recruitent_mutant/src/repository"
	"github.com/yanuvi/recruitent_mutant/src/server"
	"github.com/yanuvi/recruitent_mutant/src/usecases"
)

type MutantRequest struct {
	Dna []string `json:"dna"`
}

type MutantResponse struct {
	IsMutant bool `json:"is_mutant"`
}

func MutantHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = MutantRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		id, err := ksuid.NewRandom()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var mutantUsecases *usecases.MutantImpl = usecases.NewMutantImpl()
		usecases.SetMutantUseCases(mutantUsecases)

		requestUsecases := &entities.MutantRequest{
			Dna: request.Dna,
		}

		result := usecases.IsMutant(requestUsecases)
		resultMutant := ""
		if result {
			resultMutant = "S"
		} else {
			resultMutant = "N"
		}

		var mutant = models.Mutant{
			Dna:           strings.Join(request.Dna, ","),
			Reclutamiento: resultMutant,
			Id:            id.String(),
		}
		err = repository.InsertMutant(r.Context(), &mutant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if resultMutant == "N" {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MutantResponse{
			IsMutant: result,
		})
	}
}
