package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene

)

// KindFromSides determines the type of triangle ccording to side dimensions
func KindFromSides(a, b, c float64) Kind {

	if unusableNum(a) || unusableNum(b) || unusableNum(c) {
		return NaT
	}

	if a + b < c || b + c < a || c + a < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func unusableNum(num float64) bool {
	return num <= 0 || math.IsNaN(num) || math.IsInf(num, 0)
}