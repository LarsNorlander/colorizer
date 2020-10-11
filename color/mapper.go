package color

import "fmt"

var (
	whitePoint = Point{
		X: 0,
		Y: 1,
	}
	blackPoint = Point{
		X: 0,
		Y: 0,
	}
	whiteHueLine = Line{
		Slope: -0.268,
		Yi:    1,
	}
	blackHueLine = Line{
		Slope: 1.732,
		Yi:    0,
	}
)

func MapToWold(cw *ColorWheel, blk RGB, wht RGB, source RGB) RGB {
	src := source.ToHSL()

	pureHue := cw.Sample(src.H.Val)

	lum := PartialBlendHSL(blk.ToHSL(), wht.ToHSL(), src.L, HueDistanceCW).ToRGB()
	sat := PartialBlendRGB(lum, pureHue, src.S)

	fmt.Println(src.FormalString())
	fmt.Print("source  : ")
	fmt.Println(source)
	fmt.Print("pure hue: ")
	fmt.Println(pureHue)
	fmt.Print("sat     : ")
	fmt.Println(sat)
	fmt.Print("lum     : ")
	fmt.Println(lum)

	var result RGB

	lumLine := Line{
		Slope: 0,
		Yi:    src.L,
	}

	if src.L == 0.5 {
		result = sat
	} else if src.L > 0.5 {
		inter := intersect(lumLine, whiteHueLine)
		dist := calcDistance(inter, whitePoint)
		mixer := PartialBlendRGB(wht, pureHue, dist)
		fmt.Print("mixer   : ")
		fmt.Println(mixer)
		result = PartialBlendRGB(lum, mixer, src.S)

	} else {
		inter := intersect(lumLine, blackHueLine)
		dist := calcDistance(inter, blackPoint)
		mixer := PartialBlendRGB(blk, pureHue, dist)
		fmt.Print("mixer   : ")
		fmt.Println(mixer)
		result = PartialBlendRGB(lum, mixer, src.S)
	}

	return result
}
