package bob

import (
	"strings"
	"regexp"
)

const (
	testVersion = 2
	yellingResponse = "Whoa, chill out!"
	nothingResponse = "Fine. Be that way!"
	questionResponse = "Sure."
	defaultResponse = "Whatever."
)

var regexHasCharacters, _ = regexp.Compile(`[A-Za-z]`)

func Hey(greeting string) string {

	greeting = strings.Trim(greeting, " \t\n\r\v\u00a0\u2002")

	if len(greeting) == 0 {
		return nothingResponse
	}

	if strings.ToUpper(greeting) == greeting && regexHasCharacters.MatchString(greeting) {
		return yellingResponse
	}

	if greeting[len(greeting) - 1] == '?' {
		return questionResponse
	}

	return defaultResponse
}
