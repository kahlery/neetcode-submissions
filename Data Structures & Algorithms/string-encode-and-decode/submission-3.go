type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string

	i := 0
	for i < len(encoded) {
		// Find the '#'
		j := i
		for encoded[j] != '#' {
			j++
		}

		// Parse the length
		length, _ := strconv.Atoi(encoded[i:j])

		// Read the string
		start := j + 1
		end := start + length

		result = append(result, encoded[start:end])

		i = end
	}

	return result
}