package src

import (
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

type DnaRequestData struct {
	Dna []string `json:"dna" validate:"required"`
}

const maxDna = 4

func Mutant(c *gin.Context) {
	var requestData DnaRequestData
	if err := c.BindJSON(&requestData); err != nil {
		return
	}
	r, s := IsMutant(requestData)
	if r {
		apiResponse(http.StatusOK, s)
	} else if s != "" {
		apiResponse(http.StatusBadRequest, s)
	} else {
		apiResponse(http.StatusForbidden, s)
	}
}

func IsMutant(dna DnaRequestData) (r bool, e string) {
	b, s := IsValid(dna)
	if !b {
		return false, s
	}
	h := ToMatrix(dna)

	x := ValidateX(h[:])
	y := ValidateY(h[:])
	z := ValidateZ(h[:])

	if x+y+z > 1 {
		return true, ""
	}
	return false, ""
}

func ValidateX(m [][]string) int {
	countX := 0
	for i, _ := range m {
		count := 1
		for j := 0; j+1 <= len(m)-1; j++ {
			if m[i][j] == m[i][j+1] {
				count++
				if count == maxDna {
					countX++
					count = 1
				}
			} else {
				count = 1
			}
		}

	}
	return countX
}

func ValidateY(m [][]string) int {
	countY := 0
	for i, _ := range m {
		count := 1
		for j := 0; j < len(m)-2; j++ {
			if m[j][i] == m[j+1][i] {
				count++
				if count == maxDna {
					countY++
					count = 1
				}
			} else {
				count = 1
			}
		}

	}
	return countY
}

func ValidateZ(m [][]string) int {
	countZ := 0
	for i, _ := range m {
		count := 1
		for j := 0; j < len(m)-2; j++ {
			for k := 1; k+i <= len(m)-1 && k < len(m)-j; k++ {
				if m[i][j] == m[i+k][j+k] {
					count++
					if count == maxDna {
						countZ++
						count = 1
					}
				} else {
					count = 1
				}
			}
		}

	}
	return countZ
}

func ToMatrix(dna DnaRequestData) (m [][]string) {
	matrix := make([][]string, len(dna.Dna))
	for i, ele := range dna.Dna {
		t := strings.Split(ele, "")
		matrix[i] = t
	}
	return matrix
}

func IsValid(dna DnaRequestData) (bool, string) {
	if !IsSquareMatrix(dna) {
		return false, "Is not matrix NxN"
	}
	if !IsValidString(dna) {
		return false, "Not valid caracter"
	}
	return true, ""
}

func IsValidString(dna DnaRequestData) bool {
	for _, ele := range dna.Dna {
		t := strings.Split(ele, "")
		for j := 0; j < len(t); j++ {
			if t[j] != "A" && t[j] != "C" && t[j] != "T" && t[j] != "G" {
				return false
			}
		}
	}
	return true
}

func IsSquareMatrix(dna DnaRequestData) bool {
	for _, ele := range dna.Dna {
		if len(ele) != len(dna.Dna) {
			return false
		}
	}
	return true
}
