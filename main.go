package main

import (
	"fmt"
)

func main() {
	adnSequence := []string{"ATGCTTTTAC",
		"TATTCTGTTT",
		"GAAGCCTTAT",
		"CGGATTGCCG",
		"ATGCTGGAAG",
		"AGTTAAGGGT",
		"GTGGAATGCA",
		"GTGCGTGCTG",
		"GGTGTCTGCG",
		"GTTTTCTTGG",
	}

	result := adnAllowedLetters(adnSequence)
	if result {
		return
	}

	result = adnCheckSize(adnSequence)
	if result {
		return
	}

	result = isMutant(adnSequence)
	if result {
		fmt.Println(result)
		return
	}
}

/*
func adnCheckSize(adnSequence []string) (result bool) {
	if len(adnSequence) != len(adnSequence[0]) {
		fmt.Printf("El tamano esperado debe ser NxN y es %d x %d", len(adnSequence), len(adnSequence[0]))
		result = true
		return
	} else {
		result = false
	}
	return
}*/
