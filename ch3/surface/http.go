package main

import (
	"math"
	"net/http"

	"dberk.nl/gobook/ch3/render"
)

func handleRenderSVG(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var fn render.PlotFn = ripple
	if f, ok := r.Form["fn"]; ok {
		switch f[0] {
		case "ripple":
			fn = ripple
		case "saddle":
			fn = saddle
		case "eggbox":
			fn = eggbox
		case "moguls":
			fn = moguls
		default:
			http.Error(w, "unknown plotfn: "+f[0], 400)
			return
		}
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	render.SVG(w, fn)
}

// Ripple effect
func ripple(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// Saddle effect
func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	return (y*y/a/a - x*x/b/b)
}

func moguls(x, y float64) float64 {
	return 0.1 * (math.Cos(x) + math.Cos(y))
}

func eggbox(x, y float64) float64 {
	return math.Max(-0.2, math.Min(0.2, 0.2*(math.Cos(x)+math.Cos(y))))
}
