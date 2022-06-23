package usecases_test

import (
	"testing"

	"github.com/yanuvi/recruitent_mutant/src/entities"
	"github.com/yanuvi/recruitent_mutant/src/usecases"
)

var mutantUsecases *usecases.MutantImpl = usecases.NewMutantImpl() //declaracion explicita, pq va a ser global

func TestAdnCheckSizeVal(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases) //inyecte la implementacion

	dna := []string{"ATGC",
		"GGCT",
		"AAGT",
		"TGAC"}

	request := &entities.MutantRequest{
		Dna: dna,
	} //declarar es &

	result := usecases.IsMutant(request)

	if result {
		t.Errorf("Se esperaba que el resultado fuera falso y da verdadero")
	}
}

func TestAdnCheckSizeInv(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases) //inyecte la implementacion

	dna := []string{"ATGC",
		"GGCT",
		"AAGT"}

	request := &entities.MutantRequest{
		Dna: dna,
	} //declarar es &

	result := usecases.IsMutant(request)

	if result {
		t.Errorf("Se esperaba que el resultado fuera falso y da verdadero")
	}
}

func TestAdnAllowedLettersVal(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ATGC",
		"GGCT",
		"AAGT",
		"TGAC"}
	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if result {
		t.Errorf("Letras incorrectas de las permitidas")
	}
}

func TestAdnAllowedLettersInv(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ATGC",
		"GGCT",
		"ZSAG",
		"TGAC"}
	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if result {
		t.Errorf("Letras incorrectas de las permitidas")
	}
}

/*
func testisMutant(t *testing.T) {
	adn := []string{"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	xmen := isMutant(adn)

	if !xmen {
		t.Errorf("La secuencia de adn no es mutante")
	}
}*/
