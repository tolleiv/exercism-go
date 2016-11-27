package acronym

import (
	"strings"
	"regexp"
)

const testVersion = 1

var regexCamelCase, _ = regexp.Compile(`[a-z]([A-Z])`)
var regexSeparators, _ = regexp.Compile(`[,-;:]`)

func abbreviate(long string) string {

	cleaned := regexCamelCase.ReplaceAllString(long, " ${1}")
	cleaned = regexSeparators.ReplaceAllString(cleaned, " ")

	var apprev string
	for _, part := range strings.Split(cleaned, " ") {
		if part != "" {
			apprev += strings.ToUpper(string(part[0]))
		}
	}
	return apprev
}
