package color

import (
	"fmt"
	"math"
)

type RGB struct {
	R, G, B float64
}

func (c RGB) String() string {
	return fmt.Sprintf("rgb(%f,%f,%f)", c.R, c.G, c.B)
}

func (c RGB) AsHex() string {
	r := uint8(math.Round(255 * c.R))
	g := uint8(math.Round(255 * c.G))
	b := uint8(math.Round(255 * c.B))
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
