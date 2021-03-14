package main

import (
	"math"
	"net/http"

	"dberk.nl/gobook/ch3/render"
	"dberk.nl/gobook/ch7/eval"
)

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	render.SVG(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0, 0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}
