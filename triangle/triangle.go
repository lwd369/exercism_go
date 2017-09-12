package triangle

import (
	"math"
)

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	s := []float64{a, b, c}
	for _, slide := range s {
		if math.IsNaN(slide) || math.IsInf(slide, 0) {
			return NaT
		}
	}

	if a*b*c <= 0 {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else {
		return Sca
	}
}

type Kind int

const (
	NaT Kind = iota // not a traingle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)
