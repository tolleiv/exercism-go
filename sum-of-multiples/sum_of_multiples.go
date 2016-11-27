package summultiples

func MultipleSummer(multiples... int) (func(int) int) {
	return func(limit int) (sum int) {
		for i := 0; i < limit; i++ {
			for _, m := range multiples {
				if i % m == 0 {
					sum += i
					break
				}
			}
		}
		return sum
	}
}
