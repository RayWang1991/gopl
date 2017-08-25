package main

import (
	"math/cmplx"
	"image/color"
	"image"
	"image/png"
	"os"
)

const (
	xmin, xmax, ymin, ymax = -2, 2, -2, 2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// orignal one
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	// average
	for py := 0; py < height; py += 2 {
		for px := 0; px < width; px += 2 {
			superSampling(*img, px, py)
		}
	}
	png.Encode(os.Stdout, img)
}

func superSampling(img image.RGBA, px, py int) {
	if px < 0 || px+1 > width || py < 0 || py+1 > height {
		return
	}
	c0 := img.RGBAAt(px, py)
	c1 := img.RGBAAt(px+1, py)
	c2 := img.RGBAAt(px, py+1)
	c3 := img.RGBAAt(px+1, py+1)
	ac := averageForRGBAs([]color.RGBA{c0, c1, c2, c3})
	img.Set(px, py, ac)
	img.Set(px+1, py, ac)
	img.Set(px, py+1, ac)
	img.Set(px+1, py+1, ac)
}

func averageForRGBAs(rgbas []color.RGBA) color.RGBA {
	n := len(rgbas)
	tr, tg, tb, ta := 0, 0, 0, 0
	for _, rgba := range rgbas {
		tr += int(rgba.R)
		tg += int(rgba.G)
		tb += int(rgba.B)
		ta += int(rgba.A)
	}
	return color.RGBA{uint8(tr / n), uint8(tg / n), uint8(tb / n), uint8(ta / n)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			con := 255 - contrast*n
			return color.RGBA{con, con, 200, 200}
		}
	}
	return color.Black
}
