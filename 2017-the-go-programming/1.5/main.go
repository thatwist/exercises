package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	cycles  = 5
	nframes = 64
	res     = 0.001
	delay   = 1
	size    = 100
)

func main() {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: 64}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
	}
	gif.EncodeAll(os.Stdout, &anim)
}
