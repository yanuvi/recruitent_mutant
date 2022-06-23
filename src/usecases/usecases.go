package usecases

import "github.com/yanuvi/recruitent_mutant/src/entities"

var mutantUsecases IMutantUsecases //variable global privada

func SetMutantUseCases(usecases IMutantUsecases) { //es el punto exacto de la inyeccion de dependencia
	mutantUsecases = usecases
}

func IsMutant(request *entities.MutantRequest) (result bool) {
	return mutantUsecases.IsMutant(request.Dna)
}

//contrato de la logica de negocio, intermediario entre la invocacion y la implementacion
