package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
)

func main() {
	cw := color.GenerateColorWheel(
		color.MustParseHex("#FF4D65").ToRGB(),
		color.MustParseHex("#75ff3a").ToRGB(),
		color.MustParseHex("#278bff").ToRGB(),
	)
	blk := color.MustParseHex("#1a1f27").ToRGB()
	wht := color.MustParseHex("#f7f8fa").ToRGB()

	fmt.Println(color.MapToWold(cw, blk, wht, color.MustParseHex("#00ccff").ToRGB()))
}
