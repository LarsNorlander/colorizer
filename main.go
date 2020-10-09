package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
)

func main() {
	hex := "#1a1f27"
	fmt.Printf("Testing: %s\n", hex)

	rgb := color.MustParseHex(hex)
	hsl := color.RGBtoHSL(rgb)
	fmt.Println(rgb)
	fmt.Println(hsl)

	fmt.Println("Test converting HSL to RGB and then HEX")
	fmt.Println(color.HSLtoRGB(hsl).AsHex())

	fmt.Println("Test generating RGB gradient")
	grad1 := color.GenerateRGBGradient2(rgb, color.MustParseHex("#FF4D65"), 11)
	//grad1 := color.GenerateRGBGradient2(color.MustParseHex("#FF4D65"), color.MustParseHex("#f7f8fa"), 13)
	for i := range grad1 {
		fmt.Print(grad1[i].AsHex())
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Test generating gradient from black to white")
	grad2 := color.GenerateBlackToWhiteGradient(hsl, 28)
	for i := range grad2 {
		fmt.Print(color.HSLtoRGB(grad2[i]).AsHex())
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Test generating HSL gradient")
	grad3 := color.GenerateHSLGradient(hsl, color.RGBtoHSL(color.MustParseHex("#FF4D65")), 11)
	for i := range grad3 {
		fmt.Print(color.HSLtoRGB(grad3[i]).AsHex())
	}
	fmt.Println()
	fmt.Println()

	// GenerateLightnessGradient
	fmt.Println("Test generating lightness gradient")
	grad4 := color.GenerateLightnessGradient(hsl.H, hsl.S, hsl.L, 0.025, 10)
	for i := range grad4 {
		fmt.Print(color.HSLtoRGB(grad4[i]).AsHex())
	}
	fmt.Println()
	fmt.Println()

	//GenerateRGBGradient3
	fmt.Println("Test generating 3 point rgb gradient")
	grad5 := color.GenerateRGBGradient3(color.MustParseHex("#1a1f27"), color.MustParseHex("#FF4D65"), color.MustParseHex("#f7f8fa"), 8)
	for i := range grad5 {
		fmt.Print(grad5[i].AsHex())
	}
	fmt.Println()
	fmt.Println()
}
