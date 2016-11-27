package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency performs Frequency in parallel for all texts
func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap)
	m := FreqMap{}
	for _, text := range texts {
		go func(text string) {
			ch <- Frequency(text)
		}(text)
	}
	// Merge results
	for range texts {
		for r, c := range <-ch {
			m[r] += c
		}
	}
	return m
}