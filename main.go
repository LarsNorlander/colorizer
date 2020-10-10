package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		run()
		fmt.Println()
	}
}

func run() {
	rgbOne := randHex().AsRGB()
	rgbTwo := randHex().AsRGB()
	rgbThree := randHex().AsRGB()
	between := 9

	rgbGradient := color.GenerateRGBGradient(between,
		rgbOne,
		rgbTwo,
		rgbThree,
	)
	for i := range rgbGradient {
		fmt.Print(rgbGradient[i].AsHex())
	}
	fmt.Println()

	hslGradient := color.GenerateHSLGradient(between,
		color.RGBtoHSL(rgbOne),
		color.RGBtoHSL(rgbTwo),
		color.RGBtoHSL(rgbThree),
	)
	for i := range hslGradient {
		fmt.Print(color.HSLtoRGB(hslGradient[i]).AsHex())
	}
	fmt.Println()

	hsvGradient := color.GenerateHSVGradient(between,
		color.RGBtoHSV(rgbOne),
		color.RGBtoHSV(rgbTwo),
		color.RGBtoHSV(rgbThree),
	)
	for i := range hsvGradient {
		fmt.Print(color.HSVtoRGB(hsvGradient[i]).AsHex())
	}
	fmt.Println()
}

func randHex() color.Hex {
	return color.Hex{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
	}
}
