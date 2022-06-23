package main_test

import (
	"testing"

	"github.com/yanuvi/recruitent_mutant/src/usecases"
)

mutantUsecases := usecases.NewMutantImpl() 

usecases.SetMutantUseCases(mutantUsecases)

func testadnCheckSize(t *testing.T) {
	adn := []string{"ATGC",
		"GGCT",
		"AAGT",
		"TGAC"}
	size := adnCheckSize(adn)

	if !size {
		t.Errorf("El tamano del adn no es correcto")
	}
}

func testadnAllowedLetters(t *testing.T) {
	adn := []string{"ATGC",
		"GGCT",
		"AAGT",
		"TGAC"}
	letters := adnAllowedLetters(adn)

	if !letters {
		t.Errorf("Letras incorrectas de las permitidas")
	}
}

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
}
