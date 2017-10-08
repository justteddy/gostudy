// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	// width, height = 600, 320            // canvas size in pixels
	cells   = 100  // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	// xyscale = width / 2 / xyrange // pixels per x or y unit
	// zscale  = height * 0.4        // pixels per z unit
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var colors = []string{"white", "green", "blue"}

func main() {
	http.HandleFunc("/surface", func(w http.ResponseWriter, r *http.Request) {
		height := 320
		width := 600
		color := "white"

		heightParam := r.URL.Query().Get("height")
		widthParam := r.URL.Query().Get("width")
		colorParam := r.URL.Query().Get("color")

		if heightParam != "" {
			height, _ = strconv.Atoi(heightParam)
		}

		if widthParam != "" {
			width, _ = strconv.Atoi(widthParam)
		}

		if colorParam != "" {
			colorNum, _ := strconv.Atoi(colorParam)
			if colorNum > len(colors)-1 {
				color = "white"
			} else {
				color = colors[colorNum]
			}
		}

		xyscale := float64(width / 2 / xyrange) // pixels per x or y unit
		zscale := float64(height) * 0.4         // pixels per z unit

		makesvg(w, height, width, color, xyscale, zscale)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func makesvg(w http.ResponseWriter, height int, width int, color string, xyscale float64, zscale float64) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: red; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyscale, zscale)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j, width, height int, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
