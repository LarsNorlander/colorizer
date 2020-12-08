package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
)

func main() {
	cw := color.GenerateColorWheelFromRGB(
		color.MustParseHex("#FF0000").RGB(),
		color.MustParseHex("#00ff00").RGB(),
		color.MustParseHex("#0000ff").RGB(),
	)
	blk := color.MustParseHex("#000000").RGB()
	wht := color.MustParseHex("#ffffff").RGB()

	rgb := []color.RGB{
		//color.MustParseHex("#ff0000").RGB(),
		//color.MustParseHex("#000000").RGB(),
		//color.MustParseHex("#ffffff").RGB(),
		color.MustParseHex("#503030").RGB(),
		//color.MustParseHex("#602020").RGB(),
		//color.MustParseHex("#bf4040").RGB(),
		//color.MustParseHex("#bf4040").RGB(),
		//color.MustParseHex("#bfff40").RGB(),
	}

	for i := range rgb {
		original := rgb[i]
		mapped := color.NewMapper(cw, blk, wht)(original)
		fmt.Printf("original: %s\nmapped  : %s\n", original, mapped)
		fmt.Println()
	}
}
