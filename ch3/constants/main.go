package main

import (
	"fmt"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1024^2
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB, float64(ZiB), float64(YiB))
	fmt.Println(KB, MB, GB, TB, EB, float64(ZB), float64(YB))
}
