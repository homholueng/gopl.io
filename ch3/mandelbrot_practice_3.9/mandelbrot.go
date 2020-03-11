// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	x, _ := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	y, _ := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	zoom, _ := strconv.ParseFloat(r.URL.Query().Get("zoom"), 64)
	draw(w, x, y, zoom)
}

func draw(w io.Writer, x float64, y float64, zoom float64) {
	const (
		width, height = 1024, 1024
	)
	xmin, ymin, xmax, ymax := -x, -y, x, y

	zoomWidth := int(width * zoom)
	zoomHeight := int(height * zoom)

	img := image.NewRGBA(image.Rect(0, 0, zoomWidth, zoomHeight))
	for py := 0; py < zoomHeight; py++ {
		y := float64(py)/(width*zoom)*(ymax-ymin) + ymin
		for px := 0; px < zoomWidth; px++ {
			x := float64(px)/(height*zoom)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
