package entities

type MutantRequest struct {
	Dna []string `json:"dna"`
}

type MutantRespone struct {
	IsMutant bool `json:"is_mutant"`
}
