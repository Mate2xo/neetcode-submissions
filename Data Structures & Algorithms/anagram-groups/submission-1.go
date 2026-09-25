import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	grouped := [][]string{}

	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		groups[string(runes)] = append(groups[string(runes)], str)
	}

	for _, group := range groups {
		grouped = append(grouped, group)
	}

	return grouped

}
