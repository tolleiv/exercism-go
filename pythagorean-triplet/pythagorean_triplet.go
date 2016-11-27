package pythagorean

import "math"

type Triplet [3]int

// Sum returns all triplets with the sides adding up to the given sum
func Sum(p int) []Triplet {

	res := make([]Triplet, 0)
	triplets := Range(1, p / 2)
	for _, t := range triplets {
		if t[0] + t[1] + t[2] == p {
			res = append(res, t)
		}
	}
	return res
}

// Range returns all pythagorean triplets with the sides kept to the given bounds
func Range(min, max int) []Triplet {
	res := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			c := int(math.Sqrt(float64(a * a + b * b)))
			if a * a + b * b == c * c && c <= max {
				res = append(res, Triplet{a, b, c})

			}
		}
	}
	return res
}
