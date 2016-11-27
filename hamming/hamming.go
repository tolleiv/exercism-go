package hamming

import "errors"

const testVersion = 5

// Distance counts  the number of mutations between two DNS strands (Hamming distance)
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Sequence lenght doesn't match")
	}

	dist := 0
	for p := 0; p < len(a); p ++ {
		if b[p] != a[p] {
			dist += 1
		}
	}
	return dist, nil
}
