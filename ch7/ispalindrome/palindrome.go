package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	for i := 0; i < l/2; i++ {
		if s.Less(i, l-i-1) || s.Less(l-i-1, i) {
			return false
		}
	}
	return true
}
