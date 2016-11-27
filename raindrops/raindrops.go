package raindrops

import (
	"strconv"
)

const testVersion = 2

func Convert(drops int) string {

	var res string
	dropped := false
	if drops%3 == 0 {
		res = "Pling"
		dropped = true
	}
	if drops%5 == 0 {
		res += "Plang"
		dropped = true
	}
	if drops%7 == 0 {
		res += "Plong"
		dropped = true
	}
	if !dropped {
		res = strconv.Itoa(drops)
	}
	return res
}