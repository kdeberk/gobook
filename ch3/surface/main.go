// Surface computes an SVG redering of a 3D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/surface", handleRenderSVG)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type plotFn func(x, y float64) float64

func handleRenderSVG(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var fn plotFn = ripple
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
	renderSVG(w, fn)
}

func renderSVG(w io.Writer, f plotFn) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, f)
			bx, by, bz := corner(i, j, f)
			cx, cy, cz := corner(i, j+1, f)
			dx, dy, dz := corner(i+1, j+1, f)
			switch {
			case math.IsNaN(ax) || math.IsNaN(ay):
			case math.IsNaN(bx) || math.IsNaN(by):
			case math.IsNaN(cx) || math.IsNaN(cy):
			case math.IsNaN(dx) || math.IsNaN(dy):
			default:
				switch {
				case 0.1 <= math.Max(math.Max(math.Max(az, bz), cz), dz):
					fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: red'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
				case math.Min(math.Min(math.Min(az, bz), cz), dz) <= -0.1:
					fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: blue'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
				default:
					fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, f plotFn) (sx, sy, z float64) {
	// Find point (x,y) at corner of cells (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z = f(x, y)

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
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
