package color

type Name int

const (
	Red Name = iota
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

var (
	NameString = map[Name]string{
		0:  "Red",
		1:  "Orange",
		2:  "Yellow",
		3:  "YellowGreen",
		4:  "Green",
		5:  "GreenCyan",
		6:  "Cyan",
		7:  "CyanBlue",
		8:  "Blue",
		9:  "BlueMagenta",
		10: "Magenta",
		11: "MagentaRed",
	}
	NameValue = map[string]Name{
		"Red":         0,
		"Orange":      1,
		"Yellow":      2,
		"YellowGreen": 3,
		"Green":       4,
		"GreenCyan":   5,
		"Cyan":        6,
		"CyanBlue":    7,
		"Blue":        8,
		"BlueMagenta": 9,
		"Magenta":     10,
		"MagentaRed":  11,
	}
)

type Color interface {
	RGB() RGB
}

func NameToString(name Name) string {
	return NameString[name]
}

func StringToName(name string) Name {
	return NameValue[name]
}
