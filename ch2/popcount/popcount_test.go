package main

import "testing"

func TestPopCount(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		i := uint64(i)
		count := PopCount(i)
		if count != PopCount2(i) {
			t.Errorf("PopCount2 failed, expected: %v\tgot: %v", count, PopCount2(i))
		}
		if count != PopCount3(i) {
			t.Errorf("PopCount3 failed, expected: %v\tgot: %v", count, PopCount3(i))
		}
		if count != PopCount4(i) {
			t.Errorf("PopCount4 failed, expected: %v\tgot: %v", count, PopCount4(i))
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount(uint64(n))
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount2(uint64(n))
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount3(uint64(n))
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount4(uint64(n))
	}
}
