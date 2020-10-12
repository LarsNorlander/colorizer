package main

import (
	"fmt"
	"math/rand"

	"github.com/LarsNorlander/colorizer/color"
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

	black := color.MustParseHex(blk).RGB()
	white := color.MustParseHex(wht).RGB()

	red := color.MustParseHex(r).RGB()
	green := color.MustParseHex(g).RGB()
	blue := color.MustParseHex(b).RGB()

	cw := color.GenerateColorWheel(
		red,
		green,
		blue,
	)

	fmt.Println("The color wheel:")
	fmt.Println(cw)
	fmt.Println()

	fmt.Println("Shades (Black to White):")
	lightnessGrad := color.HSLGradient(11, color.HueDistanceNearest, black.HSL(), white.HSL())
	for i := range lightnessGrad {
		fmt.Print(lightnessGrad[i].RGB().Hex())
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
			fmt.Print(grad[i])
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
			fmt.Print(grad[i])
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
			fmt.Print(grad[i])
		}
		fmt.Println()
		cw.Next()
	}
	fmt.Println()

	fmt.Println("Tone table (Pure Hue to Gray):")
	gray := color.HSLGradient(1, color.HueDistanceNearest, black.HSL(), white.HSL())[1].RGB()
	cw.Jump(0)
	for i := 0; i < 12; i++ {
		grad := color.RGBGradient(11, cw.Get(), gray)
		for i := range grad {
			fmt.Print(grad[i])
		}
		fmt.Println()
		cw.Next()
	}
	fmt.Println()

	cw.Jump(0)
	lightnessGrad = color.HSLGradient(11, color.HueDistanceNearest, white.HSL(), black.HSL())
	for i := 0; i < 12; i++ {
		val := cw.Get()
		fmt.Printf("HSL Table for %s:\n", val)
		for i := range lightnessGrad {
			grad := color.RGBGradient(11, lightnessGrad[i].RGB(), val)
			for i := range grad {
				fmt.Print(grad[i])
			}
			fmt.Println()
		}
		fmt.Println()
		cw.Next()
	}

}

func rgbGradientTest() {
	grad := color.RGBGradient(3,
		color.MustParseHex("#1a1f27").RGB(),
		color.MustParseHex("#36FF00").RGB(),
		color.MustParseHex("#f7f8fa").RGB(),
	)
	for i := range grad {
		fmt.Print(grad[i])
	}
	fmt.Println()
}

func hueGradientTest() {
	grad := color.HSLGradient(1,
		color.HueDistanceNearest,
		[]color.Color{
			color.MustParseHex("#FFDE00"),
			color.MustParseHex("#00FFB3"),
		}...)
	for i := range grad {
		fmt.Print(grad[i].RGB().Hex())
	}
	fmt.Println()
}

func gradientDump() {
	rgbOne := randHex().RGB()
	rgbTwo := randHex().RGB()
	rgbThree := randHex().RGB()
	between := 9

	hsl := rgbTwo.HSL()
	lightnessGradient := color.HSLLumGradient(hsl.H, hsl.S, 0, 0, 19)
	for i := range lightnessGradient {
		fmt.Print(lightnessGradient[i].RGB().Hex())
	}
	fmt.Println()

	rgbGradient := color.RGBGradient(between,
		rgbOne,
		rgbTwo,
		rgbThree,
	)
	for i := range rgbGradient {
		fmt.Print(rgbGradient[i])
	}
	fmt.Println()

	hslGradient := color.HSLGradient(between,
		color.HueDistanceCW,
		rgbOne.HSL(),
		rgbTwo.HSL(),
		rgbThree.HSL(),
	)
	for i := range hslGradient {
		fmt.Print(hslGradient[i].RGB().Hex())
	}
	fmt.Println()

	reverseHslGradient := color.HSLGradient(between,
		color.HueDistanceCCW,
		rgbOne.HSL(),
		rgbTwo.HSL(),
		rgbThree.HSL(),
	)
	for i := range reverseHslGradient {
		fmt.Print(reverseHslGradient[i].RGB().Hex())
	}
	fmt.Println()

	nearestHslGradient := color.HSLGradient(between,
		color.HueDistanceNearest,
		rgbOne.HSL(),
		rgbTwo.HSL(),
		rgbThree.HSL(),
	)
	for i := range nearestHslGradient {
		fmt.Print(nearestHslGradient[i].RGB().Hex())
	}
	fmt.Println()

	hsvGradient := color.HSVGradient(between,
		color.HueDistanceCW,
		rgbOne.HSV(),
		rgbTwo.HSV(),
		rgbThree.HSV(),
	)
	for i := range hsvGradient {
		fmt.Print(hsvGradient[i].RGB().Hex())
	}
	fmt.Println()

	reverseHsvGradient := color.HSVGradient(between,
		color.HueDistanceCCW,
		rgbOne.HSV(),
		rgbTwo.HSV(),
		rgbThree.HSV(),
	)
	for i := range reverseHsvGradient {
		fmt.Print(reverseHsvGradient[i].RGB().Hex())
	}
	fmt.Println()

	nearestHsvGradient := color.HSVGradient(between,
		color.HueDistanceNearest,
		rgbOne.HSV(),
		rgbTwo.HSV(),
		rgbThree.HSV(),
	)
	for i := range nearestHsvGradient {
		fmt.Print(nearestHsvGradient[i].RGB().Hex())
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
