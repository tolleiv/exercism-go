package cryptosquare

import (
	"strings"
	"unicode"
	"math"
)

const testVersion = 2

func Encode(input string) string {
	input = normalize(input)
	output := ""
	charCount := 0
	rowCount := 0
	length := len(input)
	cols, rows, spaces := dimensions(length)

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			p := row * cols + col
			if p >= length {
				continue
			}
			output += string(input[p])
			charCount ++
			if charCount % rows == 0 && charCount < length {
				rowCount ++
				if rowCount >= cols - spaces && charCount + 1 < length {
					charCount++
				}
				output += " "
			}
		}
	}

	return output
}

func dimensions(length int) (cols, rows, spaces int) {
	sqrt := math.Sqrt(float64(length))
	cols = int(math.Ceil(sqrt))
	rows = int(math.Floor(sqrt))

	if cols * rows < length {
		rows += 1
	}
	spaces = (rows * cols) - length
	return
}

func normalize(input string) string {
	input_small := strings.ToLower(input)
	output := ""
	for _, c := range input_small {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			output += string(c)
		}
	}
	return output
}