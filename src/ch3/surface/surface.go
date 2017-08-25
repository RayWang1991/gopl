package main

import (
	"math"
	"fmt"
	"net/http"
	"log"
	"io"
)

const (
	Avogadro = 6.02214129e23
	Planck   = 6.62606957e-34
)

const (
	width, height = 600, 300            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-range .. +range)
	xyscale       = width / 2 / xyrange // pixels per x,y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // 30 degree
)

var sinA, cosA = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

// original corner func
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cosA*xyscale
	sy := height/2 + (x+y)*sinA*xyscale - z*zscale
	return sx, sy
}

func corner1(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + xyrange*(x-z*sinA*zscale)
	sy := height/2 + xyrange*(y-z*cosA*zscale)
	return sx, sy
}

func pic(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			nums := []float64{ax, ay, bx, by, cx, cy, dx, dy}
			if isValid(nums) {
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g' /> \n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func isValid(nums []float64) bool {
	for _, num := range nums {
		if math.IsInf(num, 1) || math.IsNaN(num) {
			return false
		}
	}
	return true
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("The Path is %s\n", path)
	w.Header().Set("Content-Type", "image/svg+xml")
	pic(w)
}
