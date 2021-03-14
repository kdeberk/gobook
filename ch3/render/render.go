package render

import (
	"fmt"
	"io"
	"math"
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

type PlotFn func(x, y float64) float64

func SVG(w io.Writer, f PlotFn) {
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

func corner(i, j int, f PlotFn) (sx, sy, z float64) {
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
