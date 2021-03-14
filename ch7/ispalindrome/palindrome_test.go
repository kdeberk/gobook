package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	palindromes := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 2, 3, 3, 2, 1}}
	for _, palindrome := range palindromes {
		if !IsPalindrome(sort.IntSlice(palindrome)) {
			t.Errorf("Expected %v to be a palindrome", palindrome)
		}
	}

	unpalindromes := [][]int{{1, 2}, {1, 2, 3}, {1, 2, 3, 1}}
	for _, unpalindrome := range unpalindromes {
		if IsPalindrome(sort.IntSlice(unpalindrome)) {
			t.Errorf("Expected %v to NOT be a palindrome", unpalindrome)
		}
	}
}
