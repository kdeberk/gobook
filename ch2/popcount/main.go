package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	var count byte = 0
	for i := 0; i < 8; i++ {
		count += pc[byte(x)]
		x >>= 8
	}
	return int(count)
}

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// PopCount4 returns the population count (number of set bits) of x.
func PopCount4(x uint64) int {
	count := 0
	for 0 < x {
		x &= x - 1
		count++
	}
	return count
}
