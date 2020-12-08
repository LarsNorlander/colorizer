package color

import "errors"

//goland:noinspection GoNameStartsWithPackageName
type ColorName int

//goland:noinspection GoUnusedConst
const (
	Red ColorName = iota
	Orange
	Yellow
	YellowGreen
	Green
	GreenCyan
	Cyan
	CyanBlue
	Blue
	BlueMagenta
	Magenta
	MagentaRed
)

func GenerateColorWheel(colors map[ColorName]Color) (*ColorWheel, error) {
	if len(colors) < 2 {
		return nil, errors.New("at least two colors should be provided")
	}

	cw := NewColorWheel()
	for i, color := range colors {
		cw.Jump(int(i))
		cw.Set(color)
	}

	var start int
	for i, color := range colors {
		if color == nil {
			continue
		}
		start = int(i)
		break
	}

	type pair struct {
		x, y int
	}
	var gaps []pair

	counter := start + 1
	last := start
	for {
		index := counter % 12
		color := cw.GetAt(index)

		if color != nil {
			gaps = append(gaps, pair{x: last, y: counter})
			last = index
		}

		if index == start {
			break
		}
		counter++
	}

	for _, gap := range gaps {
		gapSize := gap.y - gap.x - 1
		colors := HSLGradient(gapSize, HueDistanceCW, cw.GetAt(gap.x%12), cw.GetAt(gap.y%12))
		for i, color := range colors {
			cw.Jump((gap.x + i) % 12)
			cw.Set(color)
		}
	}

	return cw, nil
}

func GenerateColorWheelFromRGB(red Color, green Color, blue Color) *ColorWheel {
	cw, _ := GenerateColorWheel(map[ColorName]Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	})
	return cw
}
