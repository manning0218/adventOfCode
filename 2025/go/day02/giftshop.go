package day02

import (
	"fmt"
	"strconv"
)

type ProductID string

func (p ProductID) IsInvalid() bool {
	fmt.Println("Checking ProductID:", p)
	if len(p)%2 != 0 {
		return false
	}

	if p[0:len(p)/2] == p[len(p)/2:] {
		return true
	}

	return false
}

func (p ProductID) IsInvalid2(id int) bool {
	n := len(p)

	// Try all possible pattern lengths from 1 to n/2
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue
		}

		pattern := p[:patternLen]
		repeats := n / patternLen
		repeated := ""
		for i := 0; i < repeats; i++ {
			repeated += string(pattern)
		}

		if repeated == string(p) {
			return true
		}
	}

	return false
}

func (p ProductID) Value() int {
	value, _ := strconv.Atoi(string(p))
	return value
}
