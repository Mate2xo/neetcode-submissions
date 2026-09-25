type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var result strings.Builder
	for _, str := range strs {
		fmt.Fprintf(&result, "%d#%s", len(str), str)
	}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	var result []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		result = append(result, encoded[i:i+length])
		i += length
	}
	return result
}
