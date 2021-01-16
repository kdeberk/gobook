package anagram

import (
	"fmt"
	"testing"
)

func TestAnagrams(t *testing.T) {
	ts := []struct {
		a, b    string
		anagram bool
	}{
		{a: "foo", b: "bar", anagram: false},
		{a: "smaismrmilmepoetaleumibunenugttauiras", b: "altissimumplanetamtergeminumobservavi", anagram: true},
		{a: "smaismrmilmepoetaleumibunenugttauiras", b: "salueumbistineumgeminatummartiaproles", anagram: true},
	}

	for _, tt := range ts {
		t.Run(fmt.Sprintf("%v - %v", tt.a, tt.b), func(t *testing.T) {
			t.Parallel()
			if tt.anagram != are_anagrams(tt.a, tt.b) {
				t.Fatalf("a: %v, b: %v, got: %v, expected: %v", tt.a, tt.b, are_anagrams(tt.a, tt.b), tt.anagram)
			}
		})
	}
}
