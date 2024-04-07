package challenge3

func LengthOfLongestSubstring(s string) int {
	r := []rune(s)
	m := make(map[rune]int)
	i := 0
	maximum := 0
	for i < len(r) {
		index, ok := m[r[i]]
		if !ok {
			m[r[i]] = i
		} else {
			m = make(map[rune]int)
			i = index
		}
		if len(m) > maximum {
			maximum = len(m)
		}
		i++
	}
	return maximum
}

// Most efficient solution
func lengthOfLongestSubstring(s string) int {
	m, l, r := 0, 0, 0
	used := map[byte]int{}
	for r < len(s) {
		if p, ok := used[s[r]]; ok {
			if p >= l {
				l = p + 1
			}
		}
		used[s[r]] = r
		r++
		m = max(m, r-l)
	}
	return m
}
