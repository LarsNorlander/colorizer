package cmd

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/pkg/color"
	"github.com/LarsNorlander/colorizer/pkg/theme"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	command := cobra.Command{
		Use:   "map COLOR [flags]",
		Short: "Maps a color from the real world to another color space",
		RunE:  mapFn,
	}
	command.Flags().StringVarP(&themePath, "theme", "t", "", "Path to theme file")

	rootCmd.AddCommand(&command)
}

func mapFn(command *cobra.Command, args []string) error {
	// Setup
	var cw *color.ColorWheel
	var blk, wht color.Color
	var err error

	if themePath == "" {
		cw, err = color.GenerateColorWheel(map[color.Name]color.Color{
			color.Red:   color.MustParseHex("#f00"),
			color.Green: color.MustParseHex("#0f0"),
			color.Blue:  color.MustParseHex("#00f"),
		})
		if err != nil {
			return err
		}
		blk = color.MustParseHex("#000000")
		wht = color.MustParseHex("#ffffff")
	} else {
		thm, err := theme.ParseFile(themePath)
		if err != nil {
			return err
		}

		cw, err = color.GenerateColorWheel(thm.Colors)
		if err != nil {
			return err
		}
		blk = thm.Black
		wht = thm.White
	}

	arg := args[0]

	if !strings.HasPrefix(arg, "#") {
		arg = "#" + arg
	}

	hex, err := color.ParseHex(arg)
	if err != nil {
		return err
	}

	mapped := color.NewMapper(cw, blk, wht)(hex)

	fmt.Println(mapped.RGB().Hex())

	return nil
}
