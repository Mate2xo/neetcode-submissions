type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	return strings.Join(strs, "%#%")
}

func (s *Solution) Decode(encoded string) []string {
	return strings.Split(encoded, "%#%")
}
