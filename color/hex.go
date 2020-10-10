package color

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	ErrInvalidHex = errors.New("invalid hex")
)

type Hex struct {
	R, G, B uint8
}

func (hex Hex) String() string {
	return fmt.Sprintf("#%02X%02X%02X", hex.R, hex.G, hex.B)
}

func (rgb RGB) ToHex() Hex {
	return Hex{
		R: uint8(math.Round(255 * rgb.R)),
		G: uint8(math.Round(255 * rgb.G)),
		B: uint8(math.Round(255 * rgb.B)),
	}
}

func (hex Hex) ToRGB() RGB {
	return RGB{
		R: float64(hex.R) / 255.0,
		G: float64(hex.G) / 255.0,
		B: float64(hex.B) / 255.0,
	}
}

func ParseHex(hex string) (Hex, error) {
	hex = strings.ToUpper(hex) // Ensure it's in the format the scanners will expect
	hex_ := Hex{}

	var err error
	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02X%02X%02X", &hex_.R, &hex_.G, &hex_.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1X%1X%1X", &hex_.R, &hex_.G, &hex_.B)
		hex_.R *= 17
		hex_.G *= 17
		hex_.B *= 17
	default:
		err = ErrInvalidHex
	}

	if err != nil {
		return Hex{}, err
	}

	return hex_, nil
}

func MustParseHex(hex string) Hex {
	hex_, err := ParseHex(hex)
	if err != nil {
		panic(err)
	}
	return hex_
}
