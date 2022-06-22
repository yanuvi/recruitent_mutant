package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	maxSequence = 20
)

func main() {
	input := []string{"ATGCTTTTAC",
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

	result := adnAllowedLetters(input)
	if result {
		return
	}

	result = adnCheckSize(input)
	if result {
		return
	}

	result = isMutant(input)
	if result {
		fmt.Println(result)
		return
	}
}

func adnCheckSize(input []string) (result bool) {
	if len(input) != len(input[0]) {
		fmt.Printf("El tamano esperado debe ser NxN y es %d x %d", len(input), len(input[0]))
		result = true
		return
	} else {
		result = false
	}
	return
}

func adnAllowedLetters(input []string) (result bool) {
	for _, v := range input {

		valueInByte := []byte(strings.ToUpper(v))
		valueInString := string(valueInByte)

		found, _ := regexp.MatchString("[^ATGC]", valueInString)
		if found {
			fmt.Printf("El ADN solo permite las letras ATCG. Se ingreso %s \n", valueInString)
			result = true
			return
		} else {
			result = false
		}
	}
	return
}

func isMutant(input []string) (result bool) {

	contador_global := 0
	var matrix [][]byte
	var matrix_horizontal, matrix_vertical, matrix_diagonal_abajo, matrix_diagonal_arriba [10][10]int //preguntar como dejarlo sin tamano??

	for _, v := range input {
		valueInByte := []byte(strings.ToUpper(v))
		matrix = append(matrix, valueInByte)
	}

	altura := len(matrix)
	anchura := len(matrix[0])

	for i := 0; i < altura && contador_global < maxSequence; i++ {
		for j := 0; j < anchura && contador_global < maxSequence; j++ {
			//Validacion horizontal
			if (j <= (anchura - 4)) && matrix_horizontal[i][j] != 1 {
				if matrix[i][j] == matrix[i][j+1] {
					matrix_horizontal[i][j] = 1
					matrix_horizontal[i][j+1] = 1
					if matrix[i][j] == matrix[i][j+2] {
						matrix_horizontal[i][j+2] = 1
						if matrix[i][j] == matrix[i][j+3] {
							matrix_horizontal[i][j+3] = 1
							contador_global++
							//contador_h++
						}
					}
				}
			}

			//validacion vertical
			if (i <= (altura - 4)) && matrix_vertical[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j] {
					matrix_vertical[i][j] = 1
					matrix_vertical[i+1][j] = 1
					if matrix[i][j] == matrix[i+2][j] {
						matrix_vertical[i+2][j] = 1
						if matrix[i][j] == matrix[i+3][j] {
							matrix_vertical[i+3][j] = 1
							contador_global++
							//contador_v++
						}
					}
				}
			}

			//validacion diagonal abajo
			if (i <= (altura - 4)) && (j <= (anchura - 4)) && matrix_diagonal_abajo[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j+1] {
					matrix_diagonal_abajo[i][j] = 1
					matrix_diagonal_abajo[i+1][j+1] = 1
					if matrix[i][j] == matrix[i+2][j+2] {
						matrix_diagonal_abajo[i+2][j+2] = 1
						if matrix[i][j] == matrix[i+3][j+3] {
							matrix_diagonal_abajo[i+3][j+3] = 1
							contador_global++
							//contador_da++
						}
					}
				}
			}

			//validacion diagonal arriba
			if (i <= (altura - 4)) && (j > 2) && matrix_diagonal_arriba[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j-1] {
					matrix_diagonal_arriba[i][j] = 1
					matrix_diagonal_arriba[i+1][j-1] = 1
					if matrix[i][j] == matrix[i+2][j-2] {
						matrix_diagonal_arriba[i+2][j-2] = 1
						if matrix[i][j] == matrix[i+3][j-3] {
							matrix_diagonal_arriba[i+3][j-3] = 1
							contador_global++
						}
					}
				}
			}
		}
	}

	if contador_global != 0 {
		result = true
		return
	} else {
		result = false
		return
	}
}
