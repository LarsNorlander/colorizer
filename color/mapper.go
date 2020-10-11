package color

import "math"

func MapToWold(cw *ColorWheel, blk RGB, wht RGB, source RGB) RGB {
	src := source.ToHSL()

	pureHue := cw.Sample(src.H.Val)
	lum := PartialBlendHSL(blk.ToHSL(), wht.ToHSL(), src.L, HueDistanceCW).ToRGB()
	sat := PartialBlendRGB(lum, pureHue, src.S)

	distance := math.Abs(0.5 - src.L)
	percentage := distance / 0.5

	if src.L == 0.5 {
		return sat
	} else if src.L > 0.5 {
		return PartialBlendRGB(sat, wht, percentage)
	} else {
		return PartialBlendRGB(sat, blk, percentage)
	}
}
