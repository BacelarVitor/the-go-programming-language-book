package main

import (
	"image/color"
	"os"
)

var palette = []color.Color{color.RGBA{0x05, 0x99, 0x12, 0xff}, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}
