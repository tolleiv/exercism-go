package pangram

import "strings"

const testVersion = 1

const ordA = int('a')
const ordZ = int('z')

func IsPangram(sentence string) bool {
	lower := strings.ToLower(sentence)
	isPanagram := true
	for c := ordA; c < ordZ; c++ {
		isPanagram = isPanagram && strings.Contains(lower, string(rune(c)))
	}
	return isPanagram
}
