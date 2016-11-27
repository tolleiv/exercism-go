package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, chars := range input {
		for _,char := range chars {
			output[strings.ToLower(char)] = score
		}
	}
	return output
}
