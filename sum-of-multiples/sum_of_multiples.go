package summultiples

func MultipleSummer(a... int) (func(int) int) {

	return func(limit int) int {
		sumsCache := make(map[int]int)
		var s int
		for _, n := range a {
			for nn := n; nn < limit; nn += n {
				if sumsCache[nn] == 0 {
					s += nn
					sumsCache[nn] = 1
				}
			}
		}
		return s
	}
}
