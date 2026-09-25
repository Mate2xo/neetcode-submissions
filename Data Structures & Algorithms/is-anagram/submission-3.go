import "maps"

func isAnagram(s string, t string) bool {
	s_frequencies := map[rune]int{}
	t_frequencies := map[rune]int{}

	for _, char := range s { s_frequencies[char]++ }
	for _, char := range t { t_frequencies[char]++ }

	return maps.Equal(s_frequencies, t_frequencies)
}
