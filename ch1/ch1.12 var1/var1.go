// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
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
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()

		cycles, err := strconv.ParseInt(values.Get("cycles"), 10, 8)
		size, err := strconv.ParseInt(values.Get("size"), 10, 8)
		delay, err := strconv.ParseInt(values.Get("delay"), 2, 0)
		nframes, err := strconv.ParseInt(values.Get("nframes"), 10, 8)

		lissajous(w, int(cycles), int(nframes), int(size), int(delay))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, nframes int, size int, delay int) {
	var palette = []color.Color{
		color.RGBA{63, 141, 63, 1},
		color.RGBA{63, 63, 191, 1},
		color.RGBA{191, 63, 191, 1},
		color.RGBA{63, 191, 191, 1},
	}
	const (
		res = 0.001 // angular resolution
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			randInt := rand.Intn(3)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randInt))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
