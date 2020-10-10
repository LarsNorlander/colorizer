package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
	"math/rand"
)

func main() {
	grad := color.GenerateRGBGradient(3,
		color.MustParseHex("#1a1f27").AsRGB(),
		color.MustParseHex("#36FF00").AsRGB(),
		color.MustParseHex("#f7f8fa").AsRGB(),
	)
	for i := range grad {
		fmt.Print(grad[i].AsHex())
	}
	fmt.Println()
}

func hueGradientTest() {
	grad := color.GenerateNearestHSLGradient(1, []color.HSL{
		color.RGBtoHSL(color.MustParseHex("#FFDE00").AsRGB()),
		color.RGBtoHSL(color.MustParseHex("#00FFB3").AsRGB()),
	}...)
	for i := range grad {
		fmt.Print(color.HSLtoRGB(grad[i]).AsHex())
	}
	fmt.Println()
}

func gradientDump() {
	rgbOne := randHex().AsRGB()
	rgbTwo := randHex().AsRGB()
	rgbThree := randHex().AsRGB()
	between := 9

	hsl := color.RGBtoHSL(rgbTwo)
	lightnessGradient := color.GenerateLightnessGradient(
		hsl.H,
		hsl.S,
		19,
		0,
		0,
	)
	for i := range lightnessGradient {
		fmt.Print(color.HSLtoRGB(lightnessGradient[i]).AsHex())
	}
	fmt.Println()

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

	reverseHslGradient := color.GenerateReverseHSLGradient(between,
		color.RGBtoHSL(rgbOne),
		color.RGBtoHSL(rgbTwo),
		color.RGBtoHSL(rgbThree),
	)
	for i := range reverseHslGradient {
		fmt.Print(color.HSLtoRGB(reverseHslGradient[i]).AsHex())
	}
	fmt.Println()

	nearestHslGradient := color.GenerateNearestHSLGradient(between,
		color.RGBtoHSL(rgbOne),
		color.RGBtoHSL(rgbTwo),
		color.RGBtoHSL(rgbThree),
	)
	for i := range nearestHslGradient {
		fmt.Print(color.HSLtoRGB(nearestHslGradient[i]).AsHex())
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

	reverseHsvGradient := color.GenerateReverseHSVGradient(between,
		color.RGBtoHSV(rgbOne),
		color.RGBtoHSV(rgbTwo),
		color.RGBtoHSV(rgbThree),
	)
	for i := range reverseHsvGradient {
		fmt.Print(color.HSVtoRGB(reverseHsvGradient[i]).AsHex())
	}
	fmt.Println()

	nearestHsvGradient := color.GenerateNearestHSVGradient(between,
		color.RGBtoHSV(rgbOne),
		color.RGBtoHSV(rgbTwo),
		color.RGBtoHSV(rgbThree),
	)
	for i := range nearestHsvGradient {
		fmt.Print(color.HSVtoRGB(nearestHsvGradient[i]).AsHex())
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
