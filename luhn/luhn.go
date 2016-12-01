package luhn

import (
	"strings"
)

const ASCIIoffset = 48

// Valid checks whether a given number has a proper luhn checksum
func Valid(num string) bool {

	sum := luhnsum(num)
	return sum != 0 && sum % 10 == 0
}

// AddCheck adds the luhn check number to a given number
func AddCheck(num string) string {

	for i := ASCIIoffset; i < ASCIIoffset+10; i++ {
		luhnnum := num +  string(rune(i))
		if Valid(luhnnum) {
			return luhnnum
		}
	}
	return num
}

// luhnsum sums up the digits according to the luhn algo.
func luhnsum(num string) int {
	num = strings.Replace(num, " ", "", -1)

	sum := 0
	len := len(num)
	parity := len % 2

	for i := 0; i < len; i++ {
		digit := int(num[i] - ASCIIoffset)
		if i % 2 == parity {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}
		sum += digit
	}
	return sum
}