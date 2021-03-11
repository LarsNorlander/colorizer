package color

var (
	// Define luminosity points
	whitePoint = point{0, 1}
	blackPoint = point{0, 0}
	// Calculate pure hue point
	huePointX = triangleHeight(1, areaEqTriangle(1))
	huePoint  = point{huePointX, 0.5}
	// Calculate lines
	whiteHueLine = lineFromPoints(whitePoint, huePoint)
	blackHueLine = lineFromPoints(blackPoint, huePoint)
)

type Mapper func(input Color) Color

func NewMapper(hues *ColorWheel, black Color, white Color) Mapper {
	return func(input Color) Color {
		src := input.RGB().HSL()

		pureHue := hues.Sample(src.H.Val)
		lumSample := PartialHSLBlend(black, white, src.L, HueDistanceCW).RGB()
		lumLine := line{slope: 0, yIntercept: src.L}

		if src.L == 0.5 {
			return PartialRGBBlend(lumSample, pureHue, src.S)
		} else if src.L > 0.5 {
			inter := intersect(lumLine, whiteHueLine)
			dist := distanceBetweenPoints(inter, whitePoint)
			mixer := PartialRGBBlend(white, pureHue, dist)
			return PartialRGBBlend(lumSample, mixer, src.S)
		} else {
			inter := intersect(lumLine, blackHueLine)
			dist := distanceBetweenPoints(inter, blackPoint)
			mixer := PartialRGBBlend(black, pureHue, dist)
			return PartialRGBBlend(lumSample, mixer, src.S)
		}
	}
}
