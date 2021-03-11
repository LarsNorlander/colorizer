package cmd

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/pkg/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "map",
		Short: "Maps a color from the real world to another color space",
		RunE:  mapFn,
	})
}

func mapFn(command *cobra.Command, args []string) error {
	cw, err := color.GenerateColorWheel(map[color.Name]color.Color{
		color.Red:   color.MustParseHex("#FF0000").RGB(),
		color.Green: color.MustParseHex("#00ff00").RGB(),
		color.Blue:  color.MustParseHex("#0000ff").RGB(),
	})
	if err != nil {
		return err
	}

	blk := color.MustParseHex("#1A1F27").RGB()
	wht := color.MustParseHex("#f7f8fa").RGB()

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

	return nil
}
