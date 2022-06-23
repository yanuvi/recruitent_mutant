package usecases

import (
	"fmt"
	"regexp"
	"strings"
)

type MutantImpl struct {
}

type IMutantUsecases interface {
	AdnCheckSize(adnSequence []string) (result bool)
}

func NewMutantImpl() (mutantImpl *MutantImpl) {
	mutantImpl = &MutantImpl{
		//Mutant: mutant,
	}
	return
}

func adnCheckSize(adnSequence []string) (result bool) {
	if len(adnSequence) != len(adnSequence[0]) {
		fmt.Printf("El tamano esperado debe ser NxN y es %d x %d", len(adnSequence), len(adnSequence[0]))
		result = true
		return
	} else {
		result = false
	}
	return
}

func adnAllowedLetters(adnSequence []string) (result bool) {
	for _, v := range adnSequence {

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

//reciver function
func (implementation *MutantImpl) IsMutant(adn []string) (result bool) {

	result = adnAllowedLetters(adn)
	if result {
		return // retornar q es falso el adn mutante
	}

	result = adnCheckSize(adn)
	if result {
		return // retornar q es falso el adn mutante
	}

	const (
		maxSequence = 20
	) //debe estar en la funcion q lo usa, evitar condicion de carrera

	contador_global := 0
	var matrix [][]byte
	var matrix_horizontal, matrix_vertical, matrix_diagonal_down, matrix_diagonal_up [10][10]int //preguntar como dejarlo sin tamano??

	for _, v := range adn {
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
			if (i <= (altura - 4)) && (j <= (anchura - 4)) && matrix_diagonal_down[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j+1] {
					matrix_diagonal_down[i][j] = 1
					matrix_diagonal_down[i+1][j+1] = 1
					if matrix[i][j] == matrix[i+2][j+2] {
						matrix_diagonal_down[i+2][j+2] = 1
						if matrix[i][j] == matrix[i+3][j+3] {
							matrix_diagonal_down[i+3][j+3] = 1
							contador_global++
							//contador_da++
						}
					}
				}
			}

			//validacion diagonal arriba
			if (i <= (altura - 4)) && (j > 2) && matrix_diagonal_up[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j-1] {
					matrix_diagonal_up[i][j] = 1
					matrix_diagonal_up[i+1][j-1] = 1
					if matrix[i][j] == matrix[i+2][j-2] {
						matrix_diagonal_up[i+2][j-2] = 1
						if matrix[i][j] == matrix[i+3][j-3] {
							matrix_diagonal_up[i+3][j-3] = 1
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
