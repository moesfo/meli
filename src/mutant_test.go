package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func test_IsMutantTrue_When_DnaIsValid(t *testing.T) {
	var arr = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	var dna = DnaRequestData{arr}
	res, err := isMutant(dna)
	assert.True(t, res)
	assert.Empty(t, err)
}

func test_IsMutantFalse_When_DnaIsNotValid(t *testing.T) {
	var arr = []string{"ATGCGA", "CTGTGC", "TTATGT", "AGAAGG", "CCCGTA", "TCACTG"}
	var dna = DnaRequestData{arr}
	res, err := isMutant(dna)
	assert.False(t, res)
	assert.Empty(t, err)
}

func test_IsMutantFalse_When_DnaInvalidCaracter(t *testing.T) {
	var arr = []string{"XTGCGA", "CTGTGC", "TTATGT", "AGAAGG", "CCCGTA", "TCACTG"}
	var dna = DnaRequestData{arr}
	res, err := isMutant(dna)
	assert.False(t, res)
	assert.NotEmpty(t, err)
}

func test_IsMutantTrue_When_DnaIsNotMatrix(t *testing.T) {
	var arr = []string{"CTGTGC", "TTATGT", "AGAAGG", "CCCGTA", "TCACTG"}
	var dna = DnaRequestData{arr}
	res, err := isMutant(dna)
	assert.False(t, res)
	assert.NotEmpty(t, err)
}
