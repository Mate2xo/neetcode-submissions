import (
	"maps"
	"slices"
)

func minWindow(s, t string) string {
	left, shortest, charFreq := math.MinInt, "", make(map[byte]int, len(t))

	for i := 0; i < len(t); i++ {
		charFreq[t[i]]++
	}

	for right := 0; right < len(s); right++ {
		letter := s[right]
		if freq, found := charFreq[letter]; found {
			if left == math.MinInt {
				left = right
			}

			if freq > 0 {
				charFreq[letter]--
			} else {
				left++
				for !slices.Contains(slices.Collect(maps.Keys(charFreq)), s[left]) {
					left++
				}
			}

			if zeroFrequencyMap(charFreq) {
				current := s[left : right+1]
				if shortest == "" {
					shortest = current
				} else {
					shortest = min(shortest, current)
				}
			}
		}
	}
	return shortest
}

func zeroFrequencyMap(freqMap map[byte]int) bool {
	for _, count := range freqMap {
		if count == 0 {
			continue
		} else {
			return false
		}
	}

	return true
}
