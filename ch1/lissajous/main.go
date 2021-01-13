package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
}

const (
	bgIndex = 0
	fgIndex = 1
)

const (
	defaultCycles  = 5     // number of complete x oscillator revolutions
	defaultRes     = 0.001 // angular resolution
	defaultSize    = 100   // image canvas covers [-size..size]
	defaultNFrames = 64    // number of animation frames
	defaultDelay   = 8     // delay between frames in 10ms units
)

func main() {
	http.HandleFunc("/lissajous", handleLissajous)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleLissajous(w http.ResponseWriter, r *http.Request) {
	cycles, res, size, nFrames, delay := defaultCycles, defaultRes, defaultSize, defaultNFrames, defaultDelay

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if v, ok := r.Form["cycles"]; ok {
		if cycles, err = strconv.Atoi(v[0]); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if v, ok := r.Form["res"]; ok {
		if res, err = strconv.ParseFloat(v[0], 64); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if v, ok := r.Form["size"]; ok {
		if size, err = strconv.Atoi(v[0]); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if v, ok := r.Form["nFrames"]; ok {
		if nFrames, err = strconv.Atoi(v[0]); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if v, ok := r.Form["delay"]; ok {
		if delay, err = strconv.Atoi(v[0]); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
	renderLissajous(w, cycles, res, size, nFrames, delay)
}

func renderLissajous(out io.Writer, cycles int, res float64, size, nframes, delay int) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(1+int(t)%(len(palette)-1)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
