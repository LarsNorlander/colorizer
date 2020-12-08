package color

import "errors"

func GenerateColorWheel(colors map[Name]Color) (*ColorWheel, error) {
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
		index := counter
		color := cw.GetAt(index)

		if color != nil {
			gaps = append(gaps, pair{x: last, y: counter})
			last = index
		}

		if index%12 == start {
			break
		}
		counter++
	}

	for _, gap := range gaps {
		gapSize := gap.y - gap.x - 1
		colors := HSLGradient(gapSize, HueDistanceCW, cw.GetAt(gap.x), cw.GetAt(gap.y))
		for i, color := range colors {
			cw.Jump(gap.x + i)
			cw.Set(color)
		}
	}

	return cw, nil
}

func GenerateColorWheelFromRGB(red Color, green Color, blue Color) *ColorWheel {
	cw, _ := GenerateColorWheel(map[Name]Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	})
	return cw
}
