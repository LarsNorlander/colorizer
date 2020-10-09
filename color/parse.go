package color

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidHex = errors.New("invalid hex")
)

func ParseHex(hex string) (RGB, error) {
	hex = strings.ToUpper(hex) // Ensure it's in the format the scanners will expect
	values := struct {
		R uint8
		G uint8
		B uint8
	}{}

	var err error
	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02X%02X%02X", &values.R, &values.G, &values.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1X%1X%1X", &values.R, &values.G, &values.B)
		values.R *= 17
		values.G *= 17
		values.B *= 17
	default:
		err = ErrInvalidHex
	}

	if err != nil {
		return RGB{}, err
	}

	return RGB{
		R: float64(values.R) / 255.0,
		G: float64(values.G) / 255.0,
		B: float64(values.B) / 255.0,
	}, nil
}

func MustParseHex(hex string) RGB {
	color, err := ParseHex(hex)
	if err != nil {
		panic(err)
	}
	return color
}
