package color

var (
	// Define luminosity points
	whitePoint = Point{0, 1}
	blackPoint = Point{0, 0}
	// Calculate pure hue point
	huePointX = triangleHeight(1, areaEqTriangle(1))
	huePoint  = Point{huePointX, 0.5}
	// Calculate lines
	whiteHueLine = lineFromPoints(whitePoint, huePoint)
	blackHueLine = lineFromPoints(blackPoint, huePoint)
)

func MapToWold(cw *ColorWheel, blk Color, wht Color, source RGB) Color {
	src := source.HSL()

	pureHue := cw.Sample(src.H.Val)
	lumSample := PartialHSLBlend(blk, wht, src.L, HueDistanceCW).RGB()
	lumLine := Line{Slope: 0, YIntercept: src.L}

	if src.L == 0.5 {
		return PartialRGBBlend(lumSample, pureHue, src.S)
	} else if src.L > 0.5 {
		inter := intersect(lumLine, whiteHueLine)
		dist := distanceBetweenPoints(inter, whitePoint)
		mixer := PartialRGBBlend(wht, pureHue, dist)
		return PartialRGBBlend(lumSample, mixer, src.S)
	} else {
		inter := intersect(lumLine, blackHueLine)
		dist := distanceBetweenPoints(inter, blackPoint)
		mixer := PartialRGBBlend(blk, pureHue, dist)
		return PartialRGBBlend(lumSample, mixer, src.S)
	}
}
