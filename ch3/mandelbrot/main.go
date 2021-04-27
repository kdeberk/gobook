package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type fractalFunc func(complex128) color.RGBA

type config struct {
	x, y, zoom    float64
	f             fractalFunc
	supersampling bool
	width, height int
	nworkers      int
}

func (cfg config) xmin() float64 {
	return cfg.x - 2.0/cfg.zoom
}

func (cfg config) xmax() float64 {
	return cfg.x + 2.0/cfg.zoom
}

func (cfg config) ymin() float64 {
	return cfg.y - 2.0/cfg.zoom
}

func (cfg config) ymax() float64 {
	return cfg.y + 2.0/cfg.zoom
}

func (cfg config) xres() float64 {
	return (cfg.xmax() - cfg.xmin()) / float64(cfg.width)
}

func (cfg config) yres() float64 {
	return (cfg.ymax() - cfg.ymin()) / float64(cfg.height)
}

var fractalFns map[string]fractalFunc

func init() {
	fractalFns = map[string]fractalFunc{
		"mandelbrot": mandelbrot,
		"z^3-1": createNewtonFunc(newtonFract{
			func(z complex128) complex128 { return z*z*z - 1 },
			func(z complex128) complex128 { return 3 * z * z },
			[]complex128{
				1,
				complex(-.5, math.Sqrt(3)/2),
				complex(-.5, -math.Sqrt(3)/2),
			},
			1,
			1,
		}),
		"z^4-1": createNewtonFunc(newtonFract{
			func(z complex128) complex128 { return z*z*z*z - 1 },
			func(z complex128) complex128 { return 4 * z * z * z },
			[]complex128{1, -1, 1i, -1i},
			1,
			1,
		}),
		"z^3-2z+2": createNewtonFunc(newtonFract{
			func(z complex128) complex128 { return z*z*z - 2*z + 2 },
			func(z complex128) complex128 { return 3*z*z - 2 },
			[]complex128{
				-1.7693,
				0.88465 + -0.58974i,
				0.88465 + 0.58974i,
			},
			1,
			1,
		}),
	}
}

func main() {
	http.HandleFunc("/fractal", handleRenderFractal)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleRenderFractal(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		fmt.Printf("took %v\n", time.Since(start))
	}()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	cfg, err := readConfig(r.Form)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	generateImage(r.Context(), cfg, w)
}

func readConfig(vs url.Values) (config, error) {
	c := config{f: mandelbrot, x: 0.0, y: 0.0, zoom: 1, width: 1024, height: 1024, nworkers: 16}

	if name := vs.Get("fractal"); name != "" {
		f, ok := fractalFns[name]
		if !ok {
			var buffer bytes.Buffer
			fmt.Fprintf(&buffer, "Unknown fractal %v, known:\n", name)
			for name := range fractalFns {
				fmt.Fprintf(&buffer, "- %v\n", name)
			}
			return config{}, fmt.Errorf(buffer.String())
		}
		c.f = f
	}

	if x := vs.Get("x"); x != "" {
		x, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return config{}, err
		}
		c.x = x
	}

	if y := vs.Get("y"); y != "" {
		y, err := strconv.ParseFloat(y, 64)
		if err != nil {
			return config{}, err
		}
		c.y = y
	}

	if factor := vs.Get("factor"); factor != "" {
		factor, err := strconv.ParseInt(factor, 10, 64)
		if err != nil {
			return config{}, err
		}
		if factor == 0 {
			return config{}, fmt.Errorf("factor needs to be non-zero")
		}
		c.width *= int(factor)
		c.height *= int(factor)
	}

	if zoom := vs.Get("zoom"); zoom != "" {
		zoom, err := strconv.ParseFloat(zoom, 64)
		if err != nil {
			return config{}, err
		}
		if zoom == 0.0 {
			return config{}, fmt.Errorf("zoom needs to be non-zero")
		}

		c.zoom = zoom
	}

	if supersampling := vs.Get("supersampling"); supersampling != "" {
		supersampling, err := strconv.ParseBool(supersampling)
		if err != nil {
			return config{}, err
		}
		c.supersampling = supersampling
	}
	return c, nil
}

func generateImage(ctx context.Context, cfg config, w io.Writer) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	in := make(chan result, cfg.nworkers)
	out := make(chan result, cfg.nworkers)
	for i := 0; i < cfg.nworkers; i++ {
		go worker(ctx, cfg, in, out)
	}

	img := image.NewRGBA(image.Rect(0, 0, cfg.width, cfg.height))
	go func() {
		for py := 0; py < cfg.height; py++ {
			for px := 0; px < cfg.width; px++ {
				in <- result{px: px, py: py}
			}
		}
	}()

	for needed := cfg.height * cfg.width; 0 < needed; needed-- {
		r := <-out
		img.Set(r.px, r.py, r.c)
	}
	png.Encode(w, img)
}

type result struct {
	px, py int
	x, y   float64
	c      color.Color
}

func worker(ctx context.Context, cfg config, in <-chan result, out chan<- result) {
	xmin := cfg.ymin()
	ymin := cfg.ymin()
	xres := cfg.xres()
	yres := cfg.yres()

	for {
		select {
		case <-ctx.Done():
			return
		case r := <-in:
			x := float64(r.px)*xres + xmin
			y := float64(r.py)*yres + ymin

			switch {
			case cfg.supersampling:
				xx := xres / 4
				yy := yres / 4

				r.c = averageColor(
					cfg.f(complex(x-xx, y-yy)),
					cfg.f(complex(x-xx, y+yy)),
					cfg.f(complex(x+xx, y-yy)),
					cfg.f(complex(x+xx, y+yy)),
				)
			default:
				r.c = cfg.f(complex(x, y))
			}
			out <- r
		}
	}
}

func averageColor(colors ...color.RGBA) color.Color {
	var r, g, b uint16
	for _, c := range colors {
		r += uint16(c.R)
		g += uint16(c.G)
		b += uint16(c.B)
	}
	return color.RGBA{uint8(r / 4), uint8(r / 4), uint8(r / 4), 0xFF}
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if 2 < cmplx.Abs(v) {
			q := 255 * (float64(n) / float64(iterations))
			c := math.Max(0, math.Min(255, q))
			if 128 <= c {
				return color.RGBA{uint8(c), 0xFF, uint8(c), 0xFF}
			}
			return color.RGBA{0x00, uint8(c), 0x00, 0xFF}
		}
	}
	return color.RGBA{0x00, 0x00, 0x00, 0xFF}
}

type newtonFract struct {
	p     func(complex128) complex128
	deriv func(complex128) complex128
	roots []complex128
	a, z  complex128
}

func createNewtonFunc(fract newtonFract) fractalFunc {
	const (
		iterations = 200
		tolerance  = 0.00001
	)
	colors := [...]color.RGBA{
		{0xFF, 0x00, 0x00, 0xFF}, // Red
		{0x00, 0xFF, 0x00, 0xFF}, // Green
		{0x00, 0x00, 0xFF, 0xFF}, // Blue
		{0xFF, 0x00, 0xFF, 0xFF}, // Magenta
		{0xFF, 0xFF, 0x00, 0xFF}, // Yellow
	}

	if len(colors) < len(fract.roots) {
		fmt.Fprintf(os.Stderr, "Not enough colors for the roots: %v < %v", len(colors), len(fract.roots))
		os.Exit(1)
	}

	return func(z complex128) color.RGBA {
		for n := 0; n < iterations; n++ {
			z -= fract.a * fract.p(z) / fract.deriv(z)

			for i, root := range fract.roots {
				if cmplx.Abs(z-root) < tolerance {
					return colors[i]
				}
			}
		}
		return color.RGBA{0x00, 0x00, 0x00, 0xFF}
	}
}
