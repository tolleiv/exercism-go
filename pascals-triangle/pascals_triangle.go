package pascal

// Triangle generates a pascal triangle for given n
func Triangle(levels int) [][]int {
	if levels == 1 {
		return [][]int{{1}}
	}
	top := Triangle(levels - 1)
	line := make([]int, levels)
	for i := 0; i < levels; i++ {
		if i == 0 || i == levels - 1 {
			line[i] = 1
		} else {
			line[i] = top[levels - 2][i - 1] + top[levels - 2][i]
		}
	}
	return append(top, line)
}