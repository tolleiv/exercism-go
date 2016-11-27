package isogram

import (
	"unicode"
	"strings"
)

const testVersion = 1

// IsIsogram determines if the input string s is an isogram
func IsIsogram(word string) bool {
	runes := map[rune]int{}
	for _, c := range strings.ToLower(word) {
		if (!unicode.IsLetter(c)) {
			continue
		}

		if runes[c] == 1 {
			return false
		}
		runes[c] ++
	}
	return true
}
