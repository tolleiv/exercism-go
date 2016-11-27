package secret

const testVersion = 1

var moves = []struct {
	n    uint
	move string
}{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

func Handshake(num uint) []string {
	shake := []string{}

	reverse := num & 16 > 0
	for _, move := range moves {
		if num & move.n == 0 {
			continue
		}
		if reverse {
			shake = append([]string{move.move}, shake...)
		} else {
			shake = append(shake, move.move)
		}
	}
	return shake
}
