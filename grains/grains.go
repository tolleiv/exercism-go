package grains

import (
	"github.com/pkg/errors"
	"math"
)

const testVersion = 1

// Square calculates the amount of grains on a chess field
func Square(in int) (uint64, error) {
	if in <= 0 || in > 64 {
		return 0, errors.New("Overflow")
	}

	// Using recursion was less efficient
	// so using 2^in was more straight forward
	return uint64(math.Pow(2, float64(in-1))), nil
}

// Total calculates the amounts of grains on a chess board
func Total() uint64 {
	var sum uint64
	for i:=1;i<=64;i++ {
		if s,err := Square(i); err == nil {
			sum += s
		}

	}
	return sum
}