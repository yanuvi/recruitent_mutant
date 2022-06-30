package models

type Mutant struct {
	Id  string `json:"id"`
	Dna string `json:"dna"` //preguntar como se guarda ese arreglo en la bd
}
