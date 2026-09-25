import "maps"

func groupAnagrams(strs []string) [][]string {
	groups := [][]string{}
	for _, str := range strs {
		if len(groups) == 0 {
			groups = append(groups, []string{str})
			continue
		}

		grouped := false
		for i, group := range groups {
			if isAnagram(str, group[0]) {
				groups[i] = append(group, str)
				grouped = true
				break
			}
		}

		if !grouped {
			groups = append(groups, []string{str})
		}
	}

	return groups
}

func isAnagram(a, b string) bool {
	aFrequencies := map[rune]int{}
	bFrequencies := map[rune]int{}

	for _, char := range a {
		aFrequencies[char]++
	}
	for _, char := range b {
		bFrequencies[char]++
	}

	return maps.Equal(aFrequencies, bFrequencies)
}
