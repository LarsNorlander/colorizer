package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	. "github.com/LarsNorlander/colorizer/color"
)

func main() {
	// TODO Input Arguments
	// Primary Colors
	red := MustParseHex("#FF0000")
	green := MustParseHex("#00FF00")
	blue := MustParseHex("#0000FF")
	// Shades
	black := MustParseHex("#000000")
	white := MustParseHex("#FFFFFF")
	// Calculate the color wheel
	cw := GenerateColorWheel(red, green, blue)

	// TODO Input Argument
	inputFileName := "image"
	inputImg, err := os.Open(fmt.Sprintf("%s.png", inputFileName))
	if err != nil {
		panic(err)
	}
	defer inputImg.Close()

	inputConfig, _, err := image.DecodeConfig(inputImg)
	if err != nil {
		panic(err)
	}
	width := inputConfig.Width
	height := inputConfig.Height

	_, _ = inputImg.Seek(0, 0)
	img, _, err := image.Decode(inputImg)
	if err != nil {
		panic(err)
	}

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}
	imgFinal := image.NewRGBA(image.Rectangle{
		Min: upLeft,
		Max: lowRight,
	})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			val := MapToWold(cw, black, white, Hex{
				R: uint8(r / 257),
				G: uint8(g / 257),
				B: uint8(b / 257),
			}.RGB())
			imgFinal.Set(x, y, color.RGBA{
				R: val.RGB().Hex().R,
				G: val.RGB().Hex().G,
				B: val.RGB().Hex().B,
				A: 0xff,
			})
		}
	}

	// TODO Input Argument
	f, _ := os.Create(fmt.Sprintf("%s.out.png", inputFileName))
	_ = png.Encode(f, imgFinal)
}
