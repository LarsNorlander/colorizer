package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
	"math/rand"
)

func main() {
	rgbGradientTest()
	hueGradientTest()
	gradientDump()
}

func rgbGradientTest() {
	grad := color.RGBGradient(3,
		color.MustParseHex("#1a1f27").ToRGB(),
		color.MustParseHex("#36FF00").ToRGB(),
		color.MustParseHex("#f7f8fa").ToRGB(),
	)
	for i := range grad {
		fmt.Print(grad[i].ToHex())
	}
	fmt.Println()
}

func hueGradientTest() {
	grad := color.NearestHSLGradient(1, []color.HSL{
		color.MustParseHex("#FFDE00").ToRGB().ToHSL(),
		color.MustParseHex("#00FFB3").ToRGB().ToHSL(),
	}...)
	for i := range grad {
		fmt.Print(grad[i].ToRGB().ToHex())
	}
	fmt.Println()
}

func gradientDump() {
	rgbOne := randHex().ToRGB()
	rgbTwo := randHex().ToRGB()
	rgbThree := randHex().ToRGB()
	between := 9

	hsl := rgbTwo.ToHSL()
	lightnessGradient := color.LightnessGradient(
		hsl.H,
		hsl.S,
		19,
		0,
		0,
	)
	for i := range lightnessGradient {
		fmt.Print(lightnessGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	rgbGradient := color.RGBGradient(between,
		rgbOne,
		rgbTwo,
		rgbThree,
	)
	for i := range rgbGradient {
		fmt.Print(rgbGradient[i].ToHex())
	}
	fmt.Println()

	hslGradient := color.HSLGradient(between,
		rgbOne.ToHSL(),
		rgbTwo.ToHSL(),
		rgbThree.ToHSL(),
	)
	for i := range hslGradient {
		fmt.Print(hslGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	reverseHslGradient := color.ReverseHSLGradient(between,
		rgbOne.ToHSL(),
		rgbTwo.ToHSL(),
		rgbThree.ToHSL(),
	)
	for i := range reverseHslGradient {
		fmt.Print(reverseHslGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	nearestHslGradient := color.NearestHSLGradient(between,
		rgbOne.ToHSL(),
		rgbTwo.ToHSL(),
		rgbThree.ToHSL(),
	)
	for i := range nearestHslGradient {
		fmt.Print(nearestHslGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	hsvGradient := color.HSVGradient(between,
		rgbOne.ToHSV(),
		rgbTwo.ToHSV(),
		rgbThree.ToHSV(),
	)
	for i := range hsvGradient {
		fmt.Print(hsvGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	reverseHsvGradient := color.ReverseHSVGradient(between,
		rgbOne.ToHSV(),
		rgbTwo.ToHSV(),
		rgbThree.ToHSV(),
	)
	for i := range reverseHsvGradient {
		fmt.Print(reverseHsvGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	nearestHsvGradient := color.NearestHSVGradient(between,
		rgbOne.ToHSV(),
		rgbTwo.ToHSV(),
		rgbThree.ToHSV(),
	)
	for i := range nearestHsvGradient {
		fmt.Print(nearestHsvGradient[i].ToRGB().ToHex())
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
