func groupAnagrams(strs []string) [][]string {
	hash2Strings := make(map[int][]string, len(strs))

	for _, str := range strs {
		key := hash(str)
		hash2Strings[key] = append(hash2Strings[key], str)
	}

	s := len(hash2Strings)
	res := make([][]string, s)

	i := 0
	for _, v := range hash2Strings {
		res[i] = v
		i++
	}

	return res
}

func hash(str string) int {
	res := 1

	if str == "" {
		return res
	}

	for _, v := range str {
		res *= int(v)
	}

	return res
}