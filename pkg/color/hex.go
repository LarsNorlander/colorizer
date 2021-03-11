package color

import (
	"errors"
	"fmt"
	. "math"
	"strings"
)

var (
	ErrInvalidHex = errors.New("invalid hex")
)

type Hex struct {
	R, G, B uint8
}

func (rgb RGB) Hex() Hex {
	return Hex{
		R: uint8(Round(255 * rgb.R)),
		G: uint8(Round(255 * rgb.G)),
		B: uint8(Round(255 * rgb.B)),
	}
}

func (hex Hex) RGB() RGB {
	return RGB{
		R: float64(hex.R) / 255.0,
		G: float64(hex.G) / 255.0,
		B: float64(hex.B) / 255.0,
	}
}

func (hex Hex) String() string {
	return fmt.Sprintf("#%02X%02X%02X", hex.R, hex.G, hex.B)
}

func ParseHex(hex string) (Hex, error) {
	hex = strings.ToUpper(hex) // Ensure it's in the format the scanners will expect
	out := Hex{}

	var err error
	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02X%02X%02X", &out.R, &out.G, &out.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1X%1X%1X", &out.R, &out.G, &out.B)
		out.R *= 17
		out.G *= 17
		out.B *= 17
	default:
		err = ErrInvalidHex
	}

	if err != nil {
		return Hex{}, err
	}

	return out, nil
}

func MustParseHex(hex string) Hex {
	out, err := ParseHex(hex)
	if err != nil {
		panic(err)
	}
	return out
}
