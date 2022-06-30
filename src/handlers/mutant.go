package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"github.com/yanuvi/recruitent_mutant/src/models"
	"github.com/yanuvi/recruitent_mutant/src/repository"
	"github.com/yanuvi/recruitent_mutant/src/server"
)

type SingUpRequest struct {
	//	Id  string `json:"id"`
	Dna string `json:"dna"`
}

type SingUpResponse struct {
	Id  string `json:"id"`
	Dna string `json:"dna"`
}

func SingUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SingUpRequest{}
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
		err = repository.InsertMutant(r.Context(), &mutant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SingUpResponse{
			Id:  mutant.Id,
			Dna: mutant.Dna,
		})
	}
}
