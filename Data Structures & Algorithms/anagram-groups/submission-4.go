func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		anagrams[count] = append(anagrams[count], str)
	}

	result := [][]string{}
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result

}