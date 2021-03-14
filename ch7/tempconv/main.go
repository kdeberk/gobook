package main

import (
	"flag"
	"fmt"

	"dberk.nl/gobook/ch2/tempconv"
)

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) (err error) {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
	case "K", "°K":
		f.Celsius = tempconv.KToC(tempconv.Kelvin(value))
	default:
		err = fmt.Errorf("invalid temperature %q", s)
	}
	return
}

func CelsiusFlag(name string, def tempconv.Celsius, usage string) *tempconv.Celsius {
	f := &celsiusFlag{def}
	flag.CommandLine.Var(f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

// Help text contains °C even when 20.0 (the number) does not, because golang converts
// 20.0 to a value of Celsius which includes °C in its String() output, which the help
// text invokes.

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
