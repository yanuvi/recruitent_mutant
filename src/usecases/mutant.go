package usecases

import (
	"regexp"
	"strings"
)

type MutantImpl struct {
}

type IMutantUsecases interface {
	IsMutant(adnSequence []string) (result bool)
}

func NewMutantImpl() (mutantImpl *MutantImpl) {
	mutantImpl = &MutantImpl{
		//Mutant: mutant,
	}
	return
}

func adnPrerequisite(adnSequence []string) (result bool) {
	result = true
	//fmt.Println("entre aca adnpre   ")
	for _, v := range adnSequence {

		valueInByte := []byte(strings.ToUpper(v))
		valueInString := string(valueInByte)

		found, _ := regexp.MatchString("[^ATGC]", valueInString)
		if found || (len(valueInString) != len(adnSequence)) {
			result = false
			return
		}
	}
	return
}

func (implementation *MutantImpl) IsMutant(dna []string) (result bool) {
	result = adnPrerequisite(dna)
	if !result {
		return
	}

	const (
		maxSequence = 2
	)

	contador_global := 0
	var matrix [][]byte

	for _, v := range dna {
		valueInByte := []byte(strings.ToUpper(v))
		matrix = append(matrix, valueInByte)
	}

	size := len(matrix)

	matrix_horizontal := make([][]int, size)
	matrix_vertical := make([][]int, size)
	matrix_diagonal_down := make([][]int, size)
	matrix_diagonal_up := make([][]int, size)
	for idx := range matrix {
		matrix_horizontal[idx] = make([]int, size)
		matrix_vertical[idx] = make([]int, size)
		matrix_diagonal_down[idx] = make([]int, size)
		matrix_diagonal_up[idx] = make([]int, size)
	}

	for i := 0; i < size && contador_global < maxSequence; i++ {
		for j := 0; j < size && contador_global < maxSequence; j++ {
			//Validacion horizontal
			if (j <= (size - 4)) && matrix_horizontal[i][j] != 1 {
				if matrix[i][j] == matrix[i][j+1] {
					matrix_horizontal[i][j] = 1
					matrix_horizontal[i][j+1] = 1
					if matrix[i][j] == matrix[i][j+2] {
						matrix_horizontal[i][j+2] = 1
						if matrix[i][j] == matrix[i][j+3] {
							matrix_horizontal[i][j+3] = 1
							contador_global++
						}
					}
				}
			}

			//validacion vertical
			if (i <= (size - 4)) && matrix_vertical[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j] {
					matrix_vertical[i][j] = 1
					matrix_vertical[i+1][j] = 1
					if matrix[i][j] == matrix[i+2][j] {
						matrix_vertical[i+2][j] = 1
						if matrix[i][j] == matrix[i+3][j] {
							matrix_vertical[i+3][j] = 1
							contador_global++
						}
					}
				}
			}

			//validacion diagonal abajo
			if (i <= (size - 4)) && (j <= (size - 4)) && matrix_diagonal_down[i][j] != 1 {
				if matrix[i][j] == matrix[i+1][j+1] {
					matrix_diagonal_down[i][j] = 1
					matrix_diagonal_down[i+1][j+1] = 1
					if matrix[i][j] == matrix[i+2][j+2] {
						matrix_diagonal_down[i+2][j+2] = 1
						if matrix[i][j] == matrix[i+3][j+3] {
							matrix_diagonal_down[i+3][j+3] = 1
							contador_global++
						}
					}
				}
			}

			//validacion diagonal arriba
			if (i <= (size - 4)) && (j > 2) && matrix_diagonal_up[i][j] != 1 {
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
