package main

import (
	"fmt"
	mclr "github.com/LarsNorlander/colorizer/color"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {

	cw := mclr.GenerateColorWheel(
		mclr.MustParseHex("#FF0000").ToRGB(),
		mclr.MustParseHex("#00FF00").ToRGB(),
		mclr.MustParseHex("#0000ff").ToRGB(),
	)
	blk := mclr.MustParseHex("#000000").ToRGB()
	wht := mclr.MustParseHex("#ffffff").ToRGB()

	inFileName := "image"

	imgfile, err := os.Open(fmt.Sprintf("%s.png", inFileName))
	if err != nil {
		panic(err)
	}
	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)
	if err != nil {
		panic(err)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	imgfile.Seek(0, 0)

	// get the image
	img, _, err := image.Decode(imgfile)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	imgFinal := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			tmp := mclr.Hex{
				R: uint8(r / 257),
				G: uint8(g / 257),
				B: uint8(b / 257),
			}
			val := mclr.MapToWold(cw, blk, wht, tmp.ToRGB())
			//val := tmp.ToRGB()
			imgFinal.Set(x, y, color.RGBA{val.ToHex().R, val.ToHex().G, val.ToHex().B, 0xff})
		}
	}

	// Encode as PNG.
	f, _ := os.Create(fmt.Sprintf("%s.out.png", inFileName))
	png.Encode(f, imgFinal)
}

func randHex() mclr.Hex {
	return mclr.Hex{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
	}
}
