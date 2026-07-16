func groupAnagrams(strs []string) [][]string {
	hash2Strings := make(map[[26]int][]string, len(strs))

	for _, str := range strs {
		var key [26]int

		for _, c := range str {
			key[c-'a']++
		}

		hash2Strings[key] = append(hash2Strings[key], str)
	}

	res := make([][]string, 0, len(hash2Strings))

	for _, v := range hash2Strings {
		res = append(res, v)
	}

	return res
}