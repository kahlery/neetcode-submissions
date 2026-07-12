func isAnagram(s string, t string) bool {
    freq := make(map[rune]int, len(s)+len(t))
    
    for _, v := range s {
        freq[v]++
    }

    for _, v := range t {
        freq[v]--
    }

    for _, v := range freq {
        if v != 0 {
            return false
        }
    }

    return true
}
