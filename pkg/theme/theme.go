package theme

import "github.com/LarsNorlander/colorizer/pkg/color"

type Theme struct {
	Metadata Metadata
	Black    color.Color
	White    color.Color
	Colors   map[color.Name]color.Color
}

type Metadata struct {
	Name   string
	Author string
}
