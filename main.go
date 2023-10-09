package main

import (
	"image/color"
	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x60, 0x11, 0x88, 0xff}, color.RGBA{0x55, 0x79, 0x22, 0xff}, color.RGBA{0x05, 0x99, 0x12, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}
