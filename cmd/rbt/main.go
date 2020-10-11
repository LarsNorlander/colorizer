package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
	"math/rand"
)

func main() {
	testGenerateColors(
		"#000000",
		"#FFFFFF",
		"#FF0000",
		"#00FF00",
		"#0000FF")
	fmt.Println("---")
	fmt.Println()
	testGenerateColors(
		"#1a1f27",
		"#f7f8fa",
		"#FF4D65",
		"#75ff3a",
		"#278bff")
	fmt.Println("---")
	fmt.Println()
	testGenerateColors(
		"#1a1f27",
		"#f7f8fa",
		"#FF4D65",
		"#35FF55",
		"#278bff")
	fmt.Println("---")
	fmt.Println()

	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 1; i++ {
	//	testGenerateColors(randHex().String(), randHex().String(), randHex().String(), randHex().String(), randHex().String())
	//	fmt.Println("---")
	//	fmt.Println()
	//}
}

func testGenerateColors(blk string, wht string, r string, g string, b string) {
	fmt.Printf("blk: %s\n", blk)
	fmt.Printf("wht: %s\n", wht)
	fmt.Printf("red: %s\n", r)
	fmt.Printf("grn: %s\n", g)
	fmt.Printf("blu: %s\n", b)
	fmt.Println()

	black := color.MustParseHex(blk).ToRGB()
	white := color.MustParseHex(wht).ToRGB()

	red := color.MustParseHex(r).ToRGB()
	green := color.MustParseHex(g).ToRGB()
	blue := color.MustParseHex(b).ToRGB()

	cw := color.GenerateColorWheel(
		red,
		green,
		blue,
	)

	fmt.Println("The color wheel:")
	fmt.Println(cw)
	fmt.Println()

	fmt.Println("Shades (Black to White):")
	lightnessGrad := color.HSLGradient(11, color.HueDistanceNearest, black.ToHSL(), white.ToHSL())
	for i := range lightnessGrad {
		fmt.Print(lightnessGrad[i].ToRGB().ToHex())
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("The color table:")
	for i := 0; i < 12; i++ {
		grad := color.RGBGradient(5,
			black,
			cw.Get(),
			white,
		)
		for i := range grad {
			fmt.Print(grad[i].ToHex())
		}
		fmt.Println()
		cw.Next()
	}
	fmt.Println()

	fmt.Println("Shade table (Pure Hue to Black):")
	cw.Jump(0)
	for i := 0; i < 12; i++ {
		grad := color.RGBGradient(11, cw.Get(), black)
		for i := range grad {
			fmt.Print(grad[i].ToHex())
		}
		fmt.Println()
		cw.Next()
	}
	fmt.Println()

	fmt.Println("Tint table (Pure Hue to White):")
	cw.Jump(0)
	for i := 0; i < 12; i++ {
		grad := color.RGBGradient(11, cw.Get(), white)
		for i := range grad {
			fmt.Print(grad[i].ToHex())
		}
		fmt.Println()
		cw.Next()
	}
	fmt.Println()

	fmt.Println("Tone table (Pure Hue to Gray):")
	gray := color.HSLGradient(1, color.HueDistanceNearest, black.ToHSL(), white.ToHSL())[1].ToRGB()
	cw.Jump(0)
	for i := 0; i < 12; i++ {
		grad := color.RGBGradient(11, cw.Get(), gray)
		for i := range grad {
			fmt.Print(grad[i].ToHex())
		}
		fmt.Println()
		cw.Next()
	}
	fmt.Println()

	cw.Jump(0)
	lightnessGrad = color.HSLGradient(11, color.HueDistanceNearest, white.ToHSL(), black.ToHSL())
	for i := 0; i < 12; i++ {
		val := cw.Get()
		fmt.Printf("HSL Table for %s:\n", val.ToHex())
		for i := range lightnessGrad {
			grad := color.RGBGradient(11, lightnessGrad[i].ToRGB(), val)
			for i := range grad {
				fmt.Print(grad[i].ToHex())
			}
			fmt.Println()
		}
		fmt.Println()
		cw.Next()
	}

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
	grad := color.HSLGradient(1,
		color.HueDistanceNearest,
		[]color.HSL{
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
		color.HueDistanceCW,
		rgbOne.ToHSL(),
		rgbTwo.ToHSL(),
		rgbThree.ToHSL(),
	)
	for i := range hslGradient {
		fmt.Print(hslGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	reverseHslGradient := color.HSLGradient(between,
		color.HueDistanceCCW,
		rgbOne.ToHSL(),
		rgbTwo.ToHSL(),
		rgbThree.ToHSL(),
	)
	for i := range reverseHslGradient {
		fmt.Print(reverseHslGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	nearestHslGradient := color.HSLGradient(between,
		color.HueDistanceNearest,
		rgbOne.ToHSL(),
		rgbTwo.ToHSL(),
		rgbThree.ToHSL(),
	)
	for i := range nearestHslGradient {
		fmt.Print(nearestHslGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	hsvGradient := color.HSVGradient(between,
		color.HueDistanceCW,
		rgbOne.ToHSV(),
		rgbTwo.ToHSV(),
		rgbThree.ToHSV(),
	)
	for i := range hsvGradient {
		fmt.Print(hsvGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	reverseHsvGradient := color.HSVGradient(between,
		color.HueDistanceCCW,
		rgbOne.ToHSV(),
		rgbTwo.ToHSV(),
		rgbThree.ToHSV(),
	)
	for i := range reverseHsvGradient {
		fmt.Print(reverseHsvGradient[i].ToRGB().ToHex())
	}
	fmt.Println()

	nearestHsvGradient := color.HSVGradient(between,
		color.HueDistanceNearest,
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
