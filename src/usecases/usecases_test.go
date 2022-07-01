package usecases_test

import (
	"testing"

	"github.com/yanuvi/recruitent_mutant/src/entities"
	"github.com/yanuvi/recruitent_mutant/src/usecases"
)

var mutantUsecases *usecases.MutantImpl = usecases.NewMutantImpl() //declaracion explicita, pq va a ser global

func TestIsMutantSeqHoriz(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ATGCA",
		"CAAAA",
		"TTGCT",
		"GGTAC",
		"GGGGT"}

	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if !result {
		t.Errorf("La secuencia de dna no es mutante")
	}
}

func TestIsMutantSeqDiagDown(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ACAGT",
		"TATGT",
		"CTAGA",
		"TGTAC",
		"AGATG"}

	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if !result {
		t.Errorf("La secuencia de dna no es mutante")
	}
}

func TestIsMutantSeqDiagUp(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ATCGT",
		"GTGAC",
		"CGCCG",
		"GTCTA",
		"ACTAC"}

	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if !result {
		t.Errorf("La secuencia de dna no es mutante")
	}
}

func TestIsMutantSeqVert(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ACTTG",
		"CCGTC",
		"TCATG",
		"GCCTA",
		"GTAAC"}

	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if !result {
		t.Errorf("La secuencia de dna no es mutante")
	}
}

func TestIsNotMutant(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ATGCA",
		"CAGTC",
		"TACCT",
		"GTTAC",
		"AGCTT"}

	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if result {
		t.Errorf("La secuencia de dna es mutante")
	}
}

func TestAdnAllowedLetters(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases)
	dna := []string{"ATGCA",
		"GGCTA",
		"GAAAA",
		"TGACA",
		"TGZGG"}

	request := &entities.MutantRequest{
		Dna: dna,
	}
	result := usecases.IsMutant(request)

	if result {
		t.Errorf("Secuencia de dna es correcto")
	}
}

func TestAdnCheckSize(t *testing.T) {
	usecases.SetMutantUseCases(mutantUsecases) //inyecte la implementacion

	dna := []string{"ATGCA",
		"GGCTA",
		"GAAAT",
		"TGACA"}

	request := &entities.MutantRequest{
		Dna: dna,
	}

	result := usecases.IsMutant(request)

	if result {
		t.Errorf("Secuencia de dna es correcto")
	}
}
