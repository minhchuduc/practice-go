package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Height, Width int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	tmp := (x ^ y) * (x - y)
	v := uint8(tmp)
	return color.RGBA{v, v, 255, 255}
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
