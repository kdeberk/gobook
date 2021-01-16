package anagram

func are_anagrams(a, b string) bool {
	counts := map[rune]int{}

	for _, r := range a {
		counts[r]++
	}
	for _, r := range b {
		counts[r]--
	}
	for _, c := range counts {
		if 0 != c {
			return false
		}
	}
	return true
}
