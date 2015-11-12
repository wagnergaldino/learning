package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{
	Altura, Largura int
}

func (img *Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, img.Altura, img.Largura)
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
	//usando e misturando as funções do exercicio de slices:
  //uint8(x ^ y)
	//uint8((x + y)/2)
	//uint8(x * y)
  return color.RGBA{uint8(x * y), uint8(x ^ y), uint8((x + y)/2), uint8(x + y)}
}

func main() {
	img := Image{100,200}
	pic.ShowImage(&img)
}
